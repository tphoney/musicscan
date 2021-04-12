// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package albums

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"
	"github.com/go-chi/chi"
)

// HandleCreate returns an http.HandlerFunc that creates
// the object and persists to the datastore.
func HandleCreate(artists store.artistStore, albums store.albumStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		artistID, err := strconv.ParseInt(chi.URLParam(r, "artist"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse artist id")
			return
		}

		in := new(types.albumInput)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("project", projectID).
				WithField("artist", artistID).
				Debugln("cannot unmarshal json request")
			return
		}

		artist, err := artists.Find(r.Context(), artistID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", artistID).
				Debugln("artist not found")
			return
		}

		if artist.Project != projectID {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("artist", artistID).
				WithField("project", projectID).
				Debugln("project id mismatch")
			return
		}

		album := &types.album{
			artist:     artist.ID,
			Name:    in.Name.String,
			Desc:    in.Desc.String,
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		err = albums.Create(r.Context(), album)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("artist.id", artist.ID).
				WithField("artist.name", artist.Name).
				Errorln("cannot create album")
		} else {
			render.JSON(w, album, 200)
		}
	}
}
