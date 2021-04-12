// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"encoding/json"
	"net/http"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/chi"
	"gopkg.in/guregu/null.v4"
)

type userUpdateInput struct {
	Username null.String `json:"email"`
	Password null.String `json:"password"`
	Admin    null.Bool   `json:"admin"`
}

// HandleUpdate returns an http.HandlerFunc that processes an http.Request
// to update a user account.
func HandleUpdate(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(userUpdateInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Errorf("cannot unmarshal request")
			return
		}

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

		if in.Username.IsZero() == false {
			user.Email = in.Username.String
		}

		if in.Admin.Ptr() != nil {
			user.Admin = in.Admin.Bool
		}

		if in.Password.IsZero() == false {
			hash, err := bcrypt.GenerateFromPassword([]byte(in.Password.String), bcrypt.DefaultCost)
			if err != nil {
				render.InternalError(w, err)
				logger.FromRequest(r).
					WithError(err).
					Debugln("cannot hash password")
				return
			}
			user.Password = string(hash)
		}

		err = users.Update(r.Context(), user)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", user.ID).
				WithField("email", user.Email).
				Errorf("cannot update user")
		} else {
			render.JSON(w, user, 200)
		}
	}
}
