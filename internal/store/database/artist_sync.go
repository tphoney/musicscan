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

var _ store.artistStore = (*artistStoreSync)(nil)

// NewartistStoreSync returns a new artistStoreSync.
func NewartistStoreSync(store *artistStore) *artistStoreSync {
	return &artistStoreSync{store}
}

// artistStoreSync synronizes read and write access to the
// artist store. This prevents race conditions when the database
// type is sqlite3.
type artistStoreSync struct{ *artistStore }

// Find finds the artist by id.
func (s *artistStoreSync) Find(ctx context.Context, id int64) (*types.artist, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.artistStore.Find(ctx, id)
}

// List returns a list of artists.
func (s *artistStoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.artist, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.artistStore.List(ctx, id, opts)
}

// Create saves the artist details.
func (s *artistStoreSync) Create(ctx context.Context, artist *types.artist) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.artistStore.Create(ctx, artist)
}

// Update updates the artist details.
func (s *artistStoreSync) Update(ctx context.Context, artist *types.artist) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.artistStore.Update(ctx, artist)
}

// Delete deletes the artist.
func (s *artistStoreSync) Delete(ctx context.Context, artist *types.artist) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.artistStore.Delete(ctx, artist)
}
