// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"

	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/internal/store/database/mutex"
	"github.com/tphoney/musicscan/types"
)

var _ store.UserStore = (*UserStoreSync)(nil)

// NewUserStoreSync returns a new UserStoreSync.
func NewUserStoreSync(str *UserStore) *UserStoreSync {
	return &UserStoreSync{str}
}

// UserStoreSync synronizes read and write access to the
// user store. This prevents race conditions when the database
// type is sqlite3.
type UserStoreSync struct{ *UserStore }

// Find finds the user by id.
func (s *UserStoreSync) Find(ctx context.Context, id int64) (*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.UserStore.Find(ctx, id)
}

// FindEmail finds the user by email.
func (s *UserStoreSync) FindEmail(ctx context.Context, email string) (*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.UserStore.FindEmail(ctx, email)
}

// FindKey finds the user unique key (email or id).
func (s *UserStoreSync) FindKey(ctx context.Context, key string) (*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.UserStore.FindKey(ctx, key)
}

// FindToken finds the user by token.
func (s *UserStoreSync) FindToken(ctx context.Context, token string) (*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.UserStore.FindToken(ctx, token)
}

// List returns a list of users.
func (s *UserStoreSync) List(ctx context.Context, opts types.Params) ([]*types.User, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.UserStore.List(ctx, opts)
}

// Create saves the user details.
func (s *UserStoreSync) Create(ctx context.Context, user *types.User) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.UserStore.Create(ctx, user)
}

// Update updates the user details.
func (s *UserStoreSync) Update(ctx context.Context, user *types.User) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.UserStore.Update(ctx, user)
}

// Delete deletes the user.
func (s *UserStoreSync) Delete(ctx context.Context, user *types.User) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.UserStore.Delete(ctx, user)
}

// Count returns a count of users.
func (s *UserStoreSync) Count(ctx context.Context) (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()
	return s.UserStore.Count(ctx)
}
