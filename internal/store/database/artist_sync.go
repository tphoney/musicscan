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

var _ store.ArtistStore = (*ArtistStoreSync)(nil)

// NewArtistStoreSync returns a new ArtistStoreSync.
func NewArtistStoreSync(store *ArtistStore) *ArtistStoreSync {
	return &ArtistStoreSync{store}
}

// ArtistStoreSync synronizes read and write access to the
// artist store. This prevents race conditions when the database
// type is sqlite3.
type ArtistStoreSync struct{ *ArtistStore }

// Find finds the artist by id.
func (s *ArtistStoreSync) Find(ctx context.Context, id int64) (*types.Artist, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.ArtistStore.Find(ctx, id)
}

// List returns a list of artists.
func (s *ArtistStoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.Artist, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.ArtistStore.List(ctx, id, opts)
}

// Create saves the artist details.
func (s *ArtistStoreSync) Create(ctx context.Context, artist *types.Artist) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ArtistStore.Create(ctx, artist)
}

// Update updates the artist details.
func (s *ArtistStoreSync) Update(ctx context.Context, artist *types.Artist) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ArtistStore.Update(ctx, artist)
}

// Delete deletes the artist.
func (s *ArtistStoreSync) Delete(ctx context.Context, artist *types.Artist) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ArtistStore.Delete(ctx, artist)
}
