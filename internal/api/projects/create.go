// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package projects

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/api/request"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"
	"github.com/tphoney/musicscan/types/enum"

	"github.com/dchest/uniuri"
)

// HandleCreate returns an http.HandlerFunc that creates
// a new project.
func HandleCreate(projects store.ProjectStore, members store.MemberStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		viewer, _ := request.UserFrom(ctx)

		in := new(types.ProjectInput)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot unmarshal json request")
			return
		}

		project := &types.Project{
			Name:    in.Name.String,
			Desc:    in.Desc.String,
			Token:   uniuri.NewLen(uniuri.UUIDLen),
			Created: time.Now().Unix(),
			Updated: time.Now().Unix(),
		}

		err = projects.Create(ctx, project)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", in.Name).
				Errorf("cannot create project")
			return
		}

		membership := &types.Membership{
			Project: project.ID,
			User:    viewer.ID,
			Role:    enum.RoleAdmin,
		}
		if err := members.Create(ctx, membership); err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("user.email", viewer.Email).
				WithField("project.id", project.ID).
				WithField("project.name", project.Name).
				Errorln("cannot create default membership")
			return
		}

		render.JSON(w, project, 200)
	}
}
