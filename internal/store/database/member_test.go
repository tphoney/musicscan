// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.
//nolint:goconst
package database

import (
	"database/sql"
	"testing"

	"github.com/tphoney/musicscan/types"
	"github.com/tphoney/musicscan/types/enum"
)

func TestMembershipFind(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err = seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(&MemberStore{db})
	result, err := store.Find(noContext, 1, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := result.Project, int64(1); want != got {
		t.Errorf("Want account ID %d, got %d", want, got)
	}
	if got, want := result.User, int64(2); want != got {
		t.Errorf("Want user ID %d, got %d", want, got)
	}
	if got, want := result.Email, "john@example.com"; want != got {
		t.Errorf("Want user email %q, got %q", want, got)
	}
}

func TestMembershipList(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err = seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(NewMemberStore(db))
	result, err := store.List(noContext, 1, types.Params{})
	if err != nil {
		t.Error(err)
		return
	}
	if got, want := len(result), 2; want != got {
		t.Errorf("Want member count %d, got %d", want, got)
	}
	if got, want := result[0].Project, int64(1); want != got {
		t.Errorf("Want account ID %d, got %d", want, got)
	}
	if got, want := result[0].User, int64(1); want != got {
		t.Errorf("Want user ID %d, got %d", want, got)
	}
	if got, want := result[0].Email, "jane@example.com"; want != got {
		t.Errorf("Want user email %q, got %q", want, got)
	}
}

func TestMembershipCreate(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err = seed(db); err != nil {
		t.Error(err)
		return
	}
	store := NewMemberStoreSync(NewMemberStore(db))
	create := &types.Membership{
		Project: 2,
		User:    2,
		Role:    enum.RoleAdmin,
	}
	if err = store.Create(noContext, create); err != nil {
		t.Error(err)
		return
	}

	found, err := store.Find(noContext, 2, 2)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := found.Role, found.Role; got != want {
		t.Errorf("Want role %v, got %v", want, got)
	}
}

func TestMembershipUpdate(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err = seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(NewMemberStore(db))
	result, err := store.Find(noContext, 1, 2)
	if err != nil {
		t.Error(err)
		return
	}
	if result.Role != enum.RoleAdmin {
		t.Errorf("Expected admin role, got role %s", result.Role)
		return
	}

	result.Role = enum.RoleDeveloper
	err = store.Update(noContext, &types.Membership{
		Project: result.Project,
		User:    result.User,
		Role:    enum.RoleDeveloper,
	})
	if err != nil {
		t.Error(err)
		return
	}

	updated, err := store.Find(noContext, result.Project, result.User)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := updated.Role, result.Role; got != want {
		t.Errorf("Want role %v, got %v", want, got)
	}
}

func TestMembershipDelete(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	if err = seed(db); err != nil {
		t.Error(err)
		return
	}

	store := NewMemberStoreSync(NewMemberStore(db))
	_, err = store.Find(noContext, 1, 1)
	if err != nil {
		t.Error(err)
	}

	err = store.Delete(noContext, 1, 1)
	if err != nil {
		t.Error(err)
	}

	_, err = store.Find(noContext, 1, 1)
	if err != sql.ErrNoRows {
		t.Errorf("Expected ErrNoRows, got %v", err)
	}
}

func TestMembershipUniqueIndex(t *testing.T) {
	t.Skip()
}
