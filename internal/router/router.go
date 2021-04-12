// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package router provides http handlers for serving the
// web applicationa and API endpoints.
package router

import (
	"context"
	"net/http"
	"path/filepath"

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
func New(
	albumStore store.albumStore,
	artistStore store.artistStore,
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

		cors := cors.New(
			cors.Options{
				AllowedOrigins:   config.Cors.AllowedOrigins,
				AllowedMethods:   config.Cors.AllowedMethods,
				AllowedHeaders:   config.Cors.AllowedHeaders,
				ExposedHeaders:   config.Cors.ExposedHeaders,
				AllowCredentials: config.Cors.AllowCredentials,
				MaxAge:           config.Cors.MaxAge,
			},
		)
		r.Use(cors.Handler)

		// project endpoints
		r.Route("/projects", func(r chi.Router) {
			r.Use(auth)
			r.Post("/", projects.HandleCreate(projectStore, memberStore))

			// project endpoints
			r.Route("/{project}", func(r chi.Router) {
				r.Use(access.ProjectAccess(memberStore))

				r.Get("/", projects.HandleFind(projectStore))
				r.Patch("/", projects.HandleUpdate(projectStore))
				r.Delete("/", projects.HandleDelete(projectStore))

				// artist endpoints
				r.Route("/artists", func(r chi.Router) {
					r.Get("/", artists.HandleList(artistStore))
					r.Post("/", artists.HandleCreate(artistStore))
					r.Get("/{artist}", artists.HandleFind(artistStore))
					r.Patch("/{artist}", artists.HandleUpdate(artistStore))
					r.With(
						access.ProjectAdmin(memberStore),
					).Delete("/{artist}", artists.HandleDelete(artistStore))

					// album endpoints
					r.Route("/{artist}/albums", func(r chi.Router) {
						r.Get("/", albums.HandleList(artistStore, albumStore))
						r.Post("/", albums.HandleCreate(artistStore, albumStore))
						r.Get("/{album}", albums.HandleFind(artistStore, albumStore))
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
			r.Get("/projects", user.HandleProjects(projectStore))
			r.Patch("/user", user.HandleUpdate(userStore))
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

	// server all other routes from the filesystem.
	fs := http.FileServer(dist.FileSystem())
	r.With(sec.Handler).NotFound(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// because this is a single page application,
			// we need to always load the index.html file
			// in the root of the project, unless the path
			// points to a file with an extension (css, js, etc)
			if filepath.Ext(r.URL.Path) == "" {
				// HACK: alter the path to point to the
				// root of the project.
				r.URL.Path = "/"
			}
			// and finally server the file.
			fs.ServeHTTP(w, r)
		}),
	)

	return r
}
