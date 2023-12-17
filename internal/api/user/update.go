// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"encoding/json"
	"net/http"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/api/request"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword returns the bcrypt hash of the
// password at the given cost.
var hashPassword = bcrypt.GenerateFromPassword

// HandleUpdate returns an http.HandlerFunc that processes an http.Request
// to update the current user account.
func HandleUpdate(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())

		in := new(types.UserInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", viewer.Email).
				Errorf("cannot unmarshal request")
			return
		}

		if !in.Password.IsZero() {
			hash, hashErr := hashPassword([]byte(in.Password.String), bcrypt.DefaultCost)
			if hashErr != nil {
				render.InternalError(w, hashErr)
				logger.FromRequest(r).
					WithError(hashErr).
					Debugln("cannot hash password")
				return
			}
			viewer.Password = string(hash)
		}

		if !in.Name.IsZero() {
			viewer.Name = in.Name.String
		}

		if !in.Company.IsZero() {
			viewer.Company = in.Company.String
		}

		err = users.Update(r.Context(), viewer)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", viewer.Email).
				Errorf("cannot update user")
		} else {
			render.JSON(w, viewer, 200)
		}
	}
}
