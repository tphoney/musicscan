// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package router provides http handlers for serving the
// web applicationa and API endpoints.
package router

import (
	"context"
	"net/http"

	"github.com/tphoney/musicscan/internal/api/access"
	"github.com/tphoney/musicscan/internal/api/albums"
	"github.com/tphoney/musicscan/internal/api/artists"
	"github.com/tphoney/musicscan/internal/api/login"
	"github.com/tphoney/musicscan/internal/api/members"
	"github.com/tphoney/musicscan/internal/api/projects"
	"github.com/tphoney/musicscan/internal/api/register"
	"github.com/tphoney/musicscan/internal/api/system"
	"github.com/tphoney/musicscan/internal/api/token"
	"github.com/tphoney/musicscan/internal/api/user"
	"github.com/tphoney/musicscan/internal/api/users"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/internal/swagger"
	"github.com/tphoney/musicscan/web/dist"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
)

// empty context
var nocontext = context.Background()

// New returns a new http.Handler that routes traffic
// to the appropriate http.Handlers.
//nolint:funlen
func New(
	albumStore store.AlbumStore,
	artistStore store.ArtistStore,
	memberStore store.MemberStore,
	projectStore store.ProjectStore,
	userStore store.UserStore,
	systemStore store.SystemStore,
) http.Handler {
	// create the router with caching disabled
	// for API endpoints
	r := chi.NewRouter()

	// create the auth middleware.
	auth := token.Must(userStore)

	// retrieve system configuration in order to
	// retrieve security and cors configuration options.
	config := systemStore.Config(nocontext)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.NoCache)
		r.Use(middleware.Recoverer)
		r.Use(logger.Middleware)

		c := cors.New(
			cors.Options{
				AllowedOrigins:   config.Cors.AllowedOrigins,
				AllowedMethods:   config.Cors.AllowedMethods,
				AllowedHeaders:   config.Cors.AllowedHeaders,
				ExposedHeaders:   config.Cors.ExposedHeaders,
				AllowCredentials: config.Cors.AllowCredentials,
				MaxAge:           config.Cors.MaxAge,
			},
		)
		r.Use(c.Handler)

		// project endpoints
		r.Route("/projects", func(r chi.Router) {
			r.Use(auth)
			r.Post("/", projects.HandleCreate(projectStore, memberStore))

			// project endpoints
			r.Route("/{project}", func(r chi.Router) {
				r.Use(access.ProjectAccess(memberStore))

				r.Get("/", projects.HandleFind(projectStore, memberStore))
				r.Patch("/", projects.HandleUpdate(projectStore))
				r.Delete("/", projects.HandleDelete(projectStore))
				r.Get("/scan", projects.HandleScan(artistStore, albumStore))
				r.Get("/bad_albums", projects.HandleFindBadAlbums(projectStore))
				r.Get("/recommended_artists", projects.HandleFindRecommendedArtists(projectStore))
				r.Get("/wanted_albums/{year}", projects.HandleFindWantedAlbums(projectStore))
				// artist endpoints
				r.Route("/artists", func(r chi.Router) {
					r.Get("/", artists.HandleList(artistStore))
					r.Post("/", artists.HandleCreate(artistStore))
					r.Get("/{artist}", artists.HandleFind(artistStore))
					r.Get("/lookup", artists.LookupAllArtists(artistStore, albumStore))
					r.Get("/lookup/{artist}", artists.LookupSingleArtist(artistStore, albumStore))
					r.Get("/search/{artist}", artists.HandleFindByName(artistStore))
					r.Patch("/{artist}", artists.HandleUpdate(artistStore))
					r.With(
						access.ProjectAdmin(memberStore),
					).Delete("/{artist}", artists.HandleDelete(artistStore))

					// album endpoints
					r.Route("/{artist}/albums", func(r chi.Router) {
						r.Get("/", albums.HandleList(artistStore, albumStore))
						r.Post("/", albums.HandleCreate(artistStore, albumStore))
						r.Get("/{album}", albums.HandleFind(artistStore, albumStore))
						r.Get("/search/{album}", albums.HandleFindByName(artistStore, albumStore))
						r.Patch("/{album}", albums.HandleUpdate(artistStore, albumStore))
						r.With(
							access.ProjectAdmin(memberStore),
						).Delete("/{album}", albums.HandleDelete(artistStore, albumStore))
					})
				})

				// project member endpoints
				r.Route("/members", func(r chi.Router) {
					r.Use(access.ProjectAdmin(memberStore))

					r.Get("/", members.HandleList(projectStore, memberStore))
					r.Get("/{user}", members.HandleFind(userStore, projectStore, memberStore))
					r.Post("/{user}", members.HandleCreate(userStore, projectStore, memberStore))
					r.Patch("/{user}", members.HandleUpdate(userStore, projectStore, memberStore))
					r.Delete("/{user}", members.HandleDelete(userStore, projectStore, memberStore))
				})
			})
		})

		// authenticated user endpoints
		r.Route("/user", func(r chi.Router) {
			r.Use(auth)

			r.Get("/", user.HandleFind())
			r.Patch("/", user.HandleUpdate(userStore))
			r.Get("/projects", user.HandleProjects(projectStore))
			r.Post("/token", user.HandleToken(userStore))
		})

		// user management endpoints
		r.Route("/users", func(r chi.Router) {
			r.Use(auth)
			r.Use(access.SystemAdmin())

			r.Get("/", users.HandleList(userStore))
			r.Post("/", users.HandleCreate(userStore))
			r.Get("/{user}", users.HandleFind(userStore))
			r.Patch("/{user}", users.HandleUpdate(userStore))
			r.Delete("/{user}", users.HandleDelete(userStore))
		})

		// system management endpoints
		r.Route("/system", func(r chi.Router) {
			r.Get("/health", system.HandleHealth)
			r.Get("/version", system.HandleVersion)
		})

		// user login endpoint
		r.Post("/login", login.HandleLogin(userStore, systemStore))

		// user registration endpoint
		r.Post("/register", register.HandleRegister(userStore, systemStore))
	})

	// serve swagger for embedded filesystem.
	r.Handle("/swagger", http.RedirectHandler("/swagger/", http.StatusSeeOther))
	r.Handle("/swagger/*", http.StripPrefix("/swagger/", swagger.Handler()))

	// create middleware to enforce security best practices.
	sec := secure.New(
		secure.Options{
			AllowedHosts:          config.Secure.AllowedHosts,
			HostsProxyHeaders:     config.Secure.HostsProxyHeaders,
			SSLRedirect:           config.Secure.SSLRedirect,
			SSLTemporaryRedirect:  config.Secure.SSLTemporaryRedirect,
			SSLHost:               config.Secure.SSLHost,
			SSLProxyHeaders:       config.Secure.SSLProxyHeaders,
			STSSeconds:            config.Secure.STSSeconds,
			STSIncludeSubdomains:  config.Secure.STSIncludeSubdomains,
			STSPreload:            config.Secure.STSPreload,
			ForceSTSHeader:        config.Secure.ForceSTSHeader,
			FrameDeny:             config.Secure.FrameDeny,
			ContentTypeNosniff:    config.Secure.ContentTypeNosniff,
			BrowserXssFilter:      config.Secure.BrowserXSSFilter,
			ContentSecurityPolicy: config.Secure.ContentSecurityPolicy,
			ReferrerPolicy:        config.Secure.ReferrerPolicy,
		},
	)

	// serve all other routes from the embedded filesystem.
	r.With(sec.Handler).NotFound(
		dist.Handler(),
	)

	return r
}
