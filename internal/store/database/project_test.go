// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"database/sql"
	"testing"

	"github.com/tphoney/musicscan/types"
)

func TestProjectFindID(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewProjectStoreSync(NewProjectStore(db))
	result, err := store.Find(noContext, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := result.Name, "gitlab"; want != got {
		t.Errorf("Want name %q, got %q", want, got)
	}
}

func TestProjectFindToken(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewProjectStoreSync(NewProjectStore(db))
	result, err := store.FindToken(noContext, "a87ff679a2f3e71d9181a67b7542122c")
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := result.Name, "gitlab"; want != got {
		t.Errorf("Want name %q, got %q", want, got)
	}
}

func TestProjectList(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewProjectStoreSync(NewProjectStore(db))
	results, err := store.List(noContext, 1, types.Params{})
	if err != nil {
		t.Error(err)
	}
	if got, want := len(results), 2; got != want {
		t.Errorf("Want %d entities, got %d", want, got)
		return
	}
	if got, want := results[0].Name, "gitlab"; want != got {
		t.Errorf("Want name %q, got %q", want, got)
	}
	if got, want := results[1].Name, "sourcegraph"; want != got {
		t.Errorf("Want name %q, got %q", want, got)
	}
}

func TestProjectCreate(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}
	store := NewProjectStoreSync(NewProjectStore(db))
	create := &types.Project{
		Name:  "vault",
		Token: "74a03674ab3a6da96ca2ae22532d225c",
	}
	if err := store.Create(noContext, create); err != nil {
		t.Error(err)
		return
	}
	if create.ID == 0 {
		t.Errorf("Expect unique ID assigned on insert")
	}

	result, err := store.Find(noContext, create.ID)
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := result.Name, "vault"; want != got {
		t.Errorf("Want name %q, got %q", want, got)
	}
}

func TestProjectUpdate(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewProjectStoreSync(NewProjectStore(db))
	result, err := store.Find(noContext, 1)
	if err != nil {
		t.Error(err)
	}

	result.Active = !result.Active
	err = store.Update(noContext, result)
	if err != nil {
		t.Error(err)
	}

	updated, err := store.Find(noContext, result.ID)
	if err != nil {
		t.Error(err)
	}

	if got, want := updated.Active, result.Active; got != want {
		t.Errorf("Want active %v, got %v", want, got)
	}
}

func TestProjectDelete(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err := seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewProjectStoreSync(NewProjectStore(db))
	result, err := store.Find(noContext, 1)
	if err != nil {
		t.Error(err)
	}

	err = store.Delete(noContext, result)
	if err != nil {
		t.Error(err)
	}

	_, err = store.Find(noContext, 1)
	if err != sql.ErrNoRows {
		t.Errorf("Expected ErrNoRows, got %v", err)
	}
}
