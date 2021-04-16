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

var _ store.AlbumStore = (*AlbumStoreSync)(nil)

// NewAlbumStoreSync returns a new AlbumStoreSync.
func NewAlbumStoreSync(str *AlbumStore) *AlbumStoreSync {
	return &AlbumStoreSync{str}
}

// AlbumStoreSync synronizes read and write access to the
// album store. This prevents race conditions when the database
// type is sqlite3.
type AlbumStoreSync struct{ *AlbumStore }

// Find finds the album by id.
func (s *AlbumStoreSync) Find(ctx context.Context, id int64) (*types.Album, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.AlbumStore.Find(ctx, id)
}

// Find finds the album by string.
func (s *AlbumStoreSync) FindByName(ctx context.Context, name string) (*types.Album, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.AlbumStore.FindByName(ctx, name)
}

// List returns a list of albums.
func (s *AlbumStoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.Album, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.AlbumStore.List(ctx, id, opts)
}

// Create saves the album details.
func (s *AlbumStoreSync) Create(ctx context.Context, album *types.Album) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.AlbumStore.Create(ctx, album)
}

// Update updates the album details.
func (s *AlbumStoreSync) Update(ctx context.Context, album *types.Album) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.AlbumStore.Update(ctx, album)
}

// Delete deletes the album.
func (s *AlbumStoreSync) Delete(ctx context.Context, album *types.Album) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.AlbumStore.Delete(ctx, album)
}
