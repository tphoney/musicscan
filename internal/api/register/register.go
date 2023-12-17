// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package register

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"
	"golang.org/x/crypto/bcrypt"

	"github.com/dchest/uniuri"
)

// HandleRegister returns an http.HandlerFunc that processes an http.Request
// to register the named user account with the system.
func HandleRegister(users store.UserStore, system store.SystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := r.FormValue("username")
		password := r.FormValue("password")

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot hash password")
			return
		}

		user := &types.User{
			Email:    username,
			Password: string(hash),
			Token:    uniuri.NewLen(uniuri.UUIDLen),
			Created:  time.Now().Unix(),
			Updated:  time.Now().Unix(),
		}

		if err = users.Create(ctx, user); err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", user.Email).
				Errorf("cannot create user")
			return
		}

		// if the registered user is the first user of the system,
		// assume they are the system administrator and grant the
		// user system admin access.
		if user.ID == 1 {
			user.Admin = true
			if err = users.Update(ctx, user); err != nil {
				logger.FromRequest(r).
					WithError(err).
					WithField("id", user.ID).
					WithField("email", user.Email).
					Errorf("cannot enable admin user")
			}
		}

		expires := time.Now().Add(system.Config(ctx).Token.Expire)
		token, err := generate(user.ID, expires.Unix(), user.Token)
		if err != nil {
			render.InternalErrorf(w, "Failed to create session")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", user.Email).
				Errorln("failed to generate token")
			return
		}

		render.JSON(w, &types.UserToken{
			User: user,
			Token: &types.Token{
				Value:   token,
				Expires: expires.UTC(),
			},
		}, 200)
	}
}

// helper function generate a JWT token.
func generate(sub, exp int64, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp,
		"sub": sub,
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(secret))
}
