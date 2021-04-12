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

var _ store.albumStore = (*albumStoreSync)(nil)

// NewalbumStoreSync returns a new albumStoreSync.
func NewalbumStoreSync(store *albumStore) *albumStoreSync {
	return &albumStoreSync{store}
}

// albumStoreSync synronizes read and write access to the
// album store. This prevents race conditions when the database
// type is sqlite3.
type albumStoreSync struct{ *albumStore }

// Find finds the album by id.
func (s *albumStoreSync) Find(ctx context.Context, id int64) (*types.album, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.albumStore.Find(ctx, id)
}

// List returns a list of albums.
func (s *albumStoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.album, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.albumStore.List(ctx, id, opts)
}

// Create saves the album details.
func (s *albumStoreSync) Create(ctx context.Context, album *types.album) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.albumStore.Create(ctx, album)
}

// Update updates the album details.
func (s *albumStoreSync) Update(ctx context.Context, album *types.album) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.albumStore.Update(ctx, album)
}

// Delete deletes the album.
func (s *albumStoreSync) Delete(ctx context.Context, album *types.album) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.albumStore.Delete(ctx, album)
}
