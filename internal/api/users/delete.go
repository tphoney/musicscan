// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"net/http"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"

	"github.com/go-chi/chi"
)

// HandleDelete returns an http.HandlerFunc that processes an http.Request
// to delete the named user account from the system.
func HandleDelete(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "user")
		user, err := users.FindKey(r.Context(), key)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user", key).
				Errorf("cannot find user")
			return
		}
		err = users.Delete(r.Context(), user)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", user.ID).
				WithField("email", user.Email).
				Errorf("cannot delete user")
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
