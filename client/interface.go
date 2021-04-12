// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package client

import "github.com/tphoney/musicscan/types"

// Client to access the remote APIs.
type Client interface {
	// Login authenticates the user and returns a JWT token.
	Login(username, password string) (*types.Token, error)

	// Register registers a new  user and returns a JWT token.
	Register(username, password string) (*types.Token, error)

	// Self returns the currently authenticated user.
	Self() (*types.User, error)

	// Token returns an oauth2 bearer token for the currently
	// authenticated user.
	Token() (*types.Token, error)

	// User returns a user by ID or email.
	User(key string) (*types.User, error)

	// UserList returns a list of all registered users.
	UserList() ([]*types.User, error)

	// UserCreate creates a new user account.
	UserCreate(user *types.User) (*types.User, error)

	// UserUpdate updates a user account by ID or email.
	UserUpdate(key string, input *types.UserInput) (*types.User, error)

	// UserDelete deletes a user account by ID or email.
	UserDelete(key string) error

	// Project returns a project by ID.
	Project(id int64) (*types.Project, error)

	// ProjectList returns a list of all projects.
	ProjectList() ([]*types.Project, error)

	// ProjectCreate creates a new project.
	ProjectCreate(user *types.Project) (*types.Project, error)

	// ProjectUpdate updates a project.
	ProjectUpdate(id int64, input *types.ProjectInput) (*types.Project, error)

	// ProjectDelete deletes a project.
	ProjectDelete(id int64) error

	// Member returns a membrer by ID.
	Member(project int64, user string) (*types.Member, error)

	// MemberList returns a list of all project members.
	MemberList(project int64) ([]*types.Member, error)

	// MemberCreate creates a new project member.
	MemberCreate(member *types.MembershipInput) (*types.Member, error)

	// MemberUpdate updates a project member.
	MemberUpdate(member *types.MembershipInput) (*types.Member, error)

	// MemberDelete deletes a project member.
	MemberDelete(project int64, user string) error

	// artist returns a artist by ID.
	artist(project, id int64) (*types.artist, error)

	// artistList returns a list of all artists by project id.
	artistList(project int64) ([]*types.artist, error)

	// artistCreate creates a new artist.
	artistCreate(project int64, artist *types.artist) (*types.artist, error)

	// artistUpdate updates a artist.
	artistUpdate(project, id int64, input *types.artistInput) (*types.artist, error)

	// artistDelete deletes a artist.
	artistDelete(project, id int64) error

	// album returns a album by ID.
	album(project, artist, album int64) (*types.album, error)

	// albumList returns a list of all albums by project id.
	albumList(project, artist int64) ([]*types.album, error)

	// albumCreate creates a new album.
	albumCreate(project, artist int64, input *types.album) (*types.album, error)

	// albumUpdate updates a album.
	albumUpdate(project, artist, album int64, input *types.albumInput) (*types.album, error)

	// albumDelete deletes a album.
	albumDelete(project, artist, album int64) error
}

// remoteError store the error payload returned
// fro the remote API.
type remoteError struct {
	Message string `json:"message"`
}

// Error returns the error message.
func (e *remoteError) Error() string {
	return e.Message
}
