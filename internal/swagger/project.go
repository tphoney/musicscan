// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package swagger

import (
	"net/http"

	"github.com/tphoney/musicscan/types"
)

// swagger:route GET /projects/{project} project findProject
//
// Get the project with the matching project id.
//
//     Responses:
//       200: project
//
func projectFind(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /projects project createProject
//
// Create a new project.
//
//     Responses:
//       200: project
//
func projectCreate(w http.ResponseWriter, r *http.Request) {}

// swagger:route PATCH /projects/{project} project updateProject
//
// Update the project with the matching project id.
//
//     Responses:
//       200: project
//
func projectUpdate(w http.ResponseWriter, r *http.Request) {}

// swagger:route DELETE /projects/{project} project deleteProject
//
// Delete the project with the matching project id.
//
//     Responses:
//       204:
//
func projectDelete(w http.ResponseWriter, r *http.Request) {}

// swagger:parameters findProject deleteProject
type projectReq struct {
	// in: path
	ID int64 `json:"project"`
}

// swagger:parameters createProject
type projectCreateReq struct {
	// in: body
	Body types.ProjectInput
}

// swagger:parameters updateProject
type projectUpdateReq struct {
	// in: path
	ID int64 `json:"project"`

	// in: body
	Body types.ProjectInput
}

// swagger:parameters projectDelete
type projectDeleteInput struct {
	// in: path
	ID int64 `json:"project"`
}

// swagger:response project
type projectResp struct {
	// in: body
	Body types.Project
}

// swagger:response projectList
type projectListResp struct {
	// in: body
	Body []types.Project
}
