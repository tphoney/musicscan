// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package swagger

import (
	"net/http"

	"github.com/tphoney/musicscan/types"
)

// swagger:route GET /projects/{project}/members/{user} member getMember
//
// Get the project member with the matching email address.
//
//     Responses:
//       200: member
//
func memberFind(w http.ResponseWriter, r *http.Request) {}

// swagger:route GET  /projects/{project}/members member getMemberList
//
// Get the list of all project members.
//
//     Responses:
//       200: memberList
//
func memberList(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /projects/{project}/members member createMember
//
// Create a new project member.
//
//     Responses:
//       200: member
//
func memberCreate(w http.ResponseWriter, r *http.Request) {}

// swagger:route PATCH /projects/{project}/members/{user} member updateMember
//
// Update the project member.
//
//     Responses:
//       200: member
//
func memberUpdate(w http.ResponseWriter, r *http.Request) {}

// swagger:route DELETE /projects/{project}/members/{user} member deleteMember
//
// Delete the project member.
//
//     Responses:
//       204:
//
func memberDelete(w http.ResponseWriter, r *http.Request) {}

// swagger:parameters getMember deleteMember
type memberReq struct {
	// in: path
	Project int64 `json:"project"`

	// in: path
	Email string `json:"user"`
}

// swagger:parameters getMemberList
type memberListReq struct {
	// in: path
	Project int64 `json:"project"`
}

// swagger:parameters createMember
type memberCreateReq struct {
	// in: path
	Project int64 `json:"project"`

	// in: body
	Body types.MembershipInput
}

// swagger:parameters updateMember
type memberUpdateReq struct {
	// in: path
	Project int64 `json:"project"`

	// in: path
	Email string `json:"user"`

	// in: body
	Body types.MembershipInput
}

// swagger:response member
type memberResp struct {
	// in: body
	Body types.Member
}

// swagger:response memberList
type memberListResp struct {
	// in: body
	Body []types.Member
}
