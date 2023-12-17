// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package swagger

import (
	"net/http"

	"github.com/tphoney/musicscan/types"
)

// swagger:route GET /users/{user} user getUser
//
// Get the user with the matching email address.
//
//     Responses:
//       200: user
//
func userFind(w http.ResponseWriter, r *http.Request) {}

// swagger:route GET /user user getCurrentUser
//
// Get the authenticated user.
//
//     Responses:
//       200: user
//
func userCurrent(w http.ResponseWriter, r *http.Request) {}

// swagger:route GET /users user getUserList
//
// Get the list of all registered users.
//
//     Responses:
//       200: userList
//
func userList(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /users user createUser
//
// Create a new user.
//
//     Responses:
//       200: user
//
func userCreate(w http.ResponseWriter, r *http.Request) {}

// swagger:route PATCH /users/{user} user updateUser
//
// Update the user with the matching email address.
//
//     Responses:
//       200: user
//
func userUpdate(w http.ResponseWriter, r *http.Request) {}

// swagger:route DELETE /users/{user} user deleteUser
//
// Delete the user with the matching email address.
//
//     Responses:
//       204:
//
func userDelete(w http.ResponseWriter, r *http.Request) {}

// swagger:route GET /users/projects user getProjectList
//
// Get the currently authenticated user's project list.
//
//     Responses:
//       200: projectList
//
func projectList(w http.ResponseWriter, r *http.Request) {}

// swagger:parameters getUser deleteUser
type userReq struct {
	// in: path
	Email string `json:"user"`
}

// swagger:parameters createUser
type userCreateReq struct {
	// in: body
	Body types.UserInput
}

// swagger:parameters updateUser
type userUpdateReq struct {
	// in: path
	Email string `json:"user"`

	// in: body
	Body types.UserInput
}

// swagger:response user
type userResp struct {
	// in: body
	Body types.User
}

// swagger:response userList
type userListResp struct {
	// in: body
	Body []types.User
}
