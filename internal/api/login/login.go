// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package login

import (
	"net/http"
	"time"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HandleLogin returns an http.HandlerFunc that authenticates
// the user and returns an authentication token on success.
func HandleLogin(users store.UserStore, system store.SystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := r.FormValue("username")
		password := r.FormValue("password")
		user, err := users.FindEmail(ctx, username)
		if err != nil {
			render.NotFoundf(w, "Invalid email or password")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", username).
				Debugln("cannot find user")
			return
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(password),
		)
		if err != nil {
			render.NotFoundf(w, "Invalid email or password")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", username).
				Debugln("invalid password")
			return
		}

		expires := time.Now().Add(system.Config(ctx).Token.Expire)
		token, err := generate(user.ID, expires.Unix(), user.Token)
		if err != nil {
			render.InternalErrorf(w, "Failed to create session")
			logger.FromRequest(r).
				WithError(err).
				WithField("user", username).
				Debugln("failed to generate token")
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
