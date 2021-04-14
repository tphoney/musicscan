// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"
	"strconv"

	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/jmoiron/sqlx"
)

var _ store.UserStore = (*UserStore)(nil)

// NewUserStore returns a new UserStore.
func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{db}
}

// UserStore implements a UserStore backed by a relational
// database.
type UserStore struct {
	db *sqlx.DB
}

// Find finds the user by id.
func (s *UserStore) Find(ctx context.Context, id int64) (*types.User, error) {
	dst := new(types.User)
	err := s.db.Get(dst, userSelectID, id)
	return dst, err
}

// FindEmail finds the user by email.
func (s *UserStore) FindEmail(ctx context.Context, email string) (*types.User, error) {
	dst := new(types.User)
	err := s.db.Get(dst, userSelectEmail, email)
	return dst, err
}

// FindKey finds the user unique key (email or id).
func (s *UserStore) FindKey(ctx context.Context, key string) (*types.User, error) {
	id, err := strconv.ParseInt(key, 10, 64)
	if err == nil {
		return s.Find(ctx, id)
	}
	return s.FindEmail(ctx, key)
}

// FindToken finds the user by token.
func (s *UserStore) FindToken(ctx context.Context, token string) (*types.User, error) {
	dst := new(types.User)
	err := s.db.Get(dst, userSelectToken, token)
	return dst, err
}

// List returns a list of users.
func (s *UserStore) List(ctx context.Context, opts types.Params) ([]*types.User, error) {
	dst := []*types.User{}
	err := s.db.Select(&dst, userSelect)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the user details.
func (s *UserStore) Create(ctx context.Context, user *types.User) error {
	query := userInsert

	if s.db.DriverName() == POSTGRESSTRING {
		query = userInsertPg
	}

	query, arg, err := s.db.BindNamed(query, user)
	if err != nil {
		return err
	}

	if s.db.DriverName() == POSTGRESSTRING {
		return s.db.QueryRow(query, arg...).Scan(&user.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	user.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the user details.
func (s *UserStore) Update(ctx context.Context, user *types.User) error {
	query, arg, err := s.db.BindNamed(userUpdate, user)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the user.
func (s *UserStore) Delete(ctx context.Context, user *types.User) error {
	_, err := s.db.Exec(userDelete, user.ID)
	return err
}

// Count returns a count of users.
func (s *UserStore) Count(context.Context) (int64, error) {
	var count int64
	err := s.db.QueryRow(userCount).Scan(&count)
	return count, err
}

const userCount = `
SELECT count(*)
FROM users
`

const userBase = `
SELECT
 user_id
,user_email
,user_password
,user_token
,user_admin
,user_blocked
,user_created
,user_updated
,user_authed
FROM users
`

const userSelect = userBase + `
ORDER BY user_email ASC
`

const userSelectID = userBase + `
WHERE user_id = $1
`

const userSelectEmail = userBase + `
WHERE user_email = $1
`

const userSelectToken = userBase + `
WHERE user_token = $1
`

const userDelete = `
DELETE FROM users
WHERE user_id = $1
`

const userInsert = `
INSERT INTO users (
 user_email
,user_password
,user_token
,user_admin
,user_blocked
,user_created
,user_updated
,user_authed
) values (
 :user_email
,:user_password
,:user_token
,:user_admin
,:user_blocked
,:user_created
,:user_updated
,:user_authed
)
`

const userInsertPg = userInsert + `
RETURNING user_id
`

const userUpdate = `
UPDATE users
SET
 user_email     = :user_email
,user_password  = :user_password
,user_token     = :user_token
,user_admin     = :user_admin
,user_blocked   = :user_blocked
,user_created   = :user_created
,user_updated   = :user_updated
,user_authed    = :user_authed
WHERE user_id = :user_id
`
