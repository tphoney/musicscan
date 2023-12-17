// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"
	"golang.org/x/crypto/bcrypt"

	"github.com/dchest/uniuri"
)

type userCreateInput struct {
	Username string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

// HandleCreate returns an http.HandlerFunc that processes an http.Request
// to create the named user account in the system.
func HandleCreate(users store.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := new(userCreateInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot unmarshal json request")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot hash password")
			return
		}

		user := &types.User{
			Email:    in.Username,
			Admin:    in.Admin,
			Password: string(hash),
			Token:    uniuri.NewLen(uniuri.UUIDLen),
			Created:  time.Now().Unix(),
			Updated:  time.Now().Unix(),
		}

		err = users.Create(r.Context(), user)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", user.Email).
				Errorf("cannot create user")
		} else {
			render.JSON(w, user, 200)
		}
	}
}
