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

var _ store.ProjectStore = (*ProjectStoreSync)(nil)

// NewProjectStoreSync returns a new ProjectStoreSync.
func NewProjectStoreSync(str *ProjectStore) *ProjectStoreSync {
	return &ProjectStoreSync{str}
}

// ProjectStoreSync synronizes read and write access to the
// project store. This prevents race conditions when the database
// type is sqlite3.
type ProjectStoreSync struct{ *ProjectStore }

// Find finds the project by id.
func (s *ProjectStoreSync) Find(ctx context.Context, id int64) (*types.Project, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.ProjectStore.Find(ctx, id)
}

// List returns a list of projects by user.
func (s *ProjectStoreSync) List(ctx context.Context, id int64, opts types.Params) ([]*types.Project, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	return s.ProjectStore.List(ctx, id, opts)
}

// Create saves the project details.
func (s *ProjectStoreSync) Create(ctx context.Context, project *types.Project) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ProjectStore.Create(ctx, project)
}

// Update updates the project details.
func (s *ProjectStoreSync) Update(ctx context.Context, project *types.Project) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ProjectStore.Update(ctx, project)
}

// Delete deletes the project.
func (s *ProjectStoreSync) Delete(ctx context.Context, project *types.Project) error {
	mutex.Lock()
	defer mutex.Unlock()
	return s.ProjectStore.Delete(ctx, project)
}
