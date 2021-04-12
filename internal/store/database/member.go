// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"

	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/jmoiron/sqlx"
)

var _ store.MemberStore = (*MemberStore)(nil)

// NewMemberStore returns a new MemberStore.
func NewMemberStore(db *sqlx.DB) *MemberStore {
	return &MemberStore{db}
}

// MemberStore implements a MemberStore backed by a
// relational database.
type MemberStore struct {
	db *sqlx.DB
}

// Find finds the member by project and user id.
func (s *MemberStore) Find(ctx context.Context, project int64, user int64) (*types.Member, error) {
	dst := new(types.Member)
	err := s.db.Get(dst, memberSelect, project, user)
	return dst, err
}

// List returns a list of members.
func (s *MemberStore) List(ctx context.Context, project int64, opts types.Params) ([]*types.Member, error) {
	dst := []*types.Member{}
	err := s.db.Select(&dst, Params, project)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the membership details.
func (s *MemberStore) Create(ctx context.Context, membership *types.Membership) error {
	_, err := s.db.Exec(
		memberInsert,
		membership.Project,
		membership.User,
		membership.Role,
	)
	return err
}

// Update updates the membership details.
func (s *MemberStore) Update(ctx context.Context, membership *types.Membership) error {
	_, err := s.db.Exec(
		memberUpdate,
		membership.Role,
		membership.Project,
		membership.User,
	)
	return err
}

// Delete deletes the membership.
func (s *MemberStore) Delete(ctx context.Context, project, user int64) error {
	_, err := s.db.Exec(memberDelete, project, user)
	return err
}

const memberBase = `
SELECT
 user_email
,member_project_id
,member_user_id
,member_role
FROM members
`

const Params = memberBase + `
INNER JOIN users
ON members.member_user_id = users.user_id
WHERE member_project_id = $1
ORDER BY users.user_email
`

const memberSelect = memberBase + `
INNER JOIN users
ON members.member_user_id = users.user_id
WHERE member_project_id = $1
  AND member_user_id    = $2
`

const memberDelete = `
DELETE FROM members
WHERE member_project_id = $1
  AND member_user_id    = $2
`

const memberInsert = `
INSERT INTO members (
 member_project_id
,member_user_id
,member_role
) values ($1, $2, $3)
`

const memberUpdate = `
UPDATE members
SET member_role = $1
WHERE member_project_id = $2
  AND member_user_id    = $3
`
