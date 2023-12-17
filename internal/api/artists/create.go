// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artists

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"
)

// HandleCreate returns an http.HandlerFunc that creates
// the object and persists to the datastore.
func HandleCreate(artists store.ArtistStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		project, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		in := new(types.ArtistInput)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("project", project).
				Debugln("cannot unmarshal json request")
			return
		}

		artist := &types.Artist{
			Project: project,
			Name:    in.Name.String,
			Desc:    in.Desc.String,
			Wanted:  in.Wanted.Bool,
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		err = artists.Create(r.Context(), artist)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", artist.Name).
				WithField("project", project).
				Errorln("cannot create artist")
		} else {
			render.JSON(w, artist, http.StatusCreated)
		}
	}
}
