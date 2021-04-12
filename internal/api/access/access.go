// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package access

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/api/request"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types/enum"

	"github.com/go-chi/chi"
)

// ProjectAccess returns an http.HandlerFunc middleware that authorizes
// the user read access to the project.
func ProjectAccess(members store.MemberStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user, ok := request.UserFrom(ctx)
			if !ok {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}

			// if the user is an administrator they are automatically
			// granted access to the endpoint.
			if user.Admin {
				logger.FromRequest(r).
					Debugln("admin user granted read access")
				next.ServeHTTP(w, r)
				return
			}

			id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
			if err != nil {
				render.BadRequest(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot parse project id")
				return
			}

			member, err := members.Find(ctx, id, user.ID)
			if err != nil {
				render.NotFound(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot find project membership")
				return
			}

			logger.FromRequest(r).
				WithField("role", member.Role).
				Debugln("user granted read access")

			next.ServeHTTP(w, r)
		})
	}
}

// ProjectAdmin returns an http.HandlerFunc middleware that authorizes
// the user admin access to the project.
func ProjectAdmin(members store.MemberStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user, ok := request.UserFrom(ctx)
			if !ok {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}

			// if the user is an administrator they are automatically
			// granted access to the endpoint.
			if user.Admin {
				next.ServeHTTP(w, r)
				return
			}

			id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
			if err != nil {
				render.BadRequest(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot parse project id")
				return
			}

			member, err := members.Find(ctx, id, user.ID)
			if err != nil {
				render.NotFound(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot find project membership")
				return
			}

			if member.Role != enum.RoleAdmin {
				render.ErrorCode(w, errors.New("Forbidden"), 403)
				logger.FromRequest(r).
					WithError(err).
					Debugln("insufficient privileges")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// SystemAdmin returns an http.HandlerFunc middleware that authorizes
// the user access to system administration capabilities.
func SystemAdmin() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			user, ok := request.UserFrom(ctx)
			if !ok {
				render.ErrorCode(w, errors.New("Requires authentication"), 401)
				return
			}
			if !user.Admin {
				render.ErrorCode(w, errors.New("Forbidden"), 403)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
