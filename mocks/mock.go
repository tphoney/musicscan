// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package mocks provides mock interfaces.
package mocks

//go:generate mockgen -package=mocks -destination=mock_store.go github.com/tphoney/musicscan/internal/store albumStore,artistStore,MemberStore,ProjectStore,SystemStore,UserStore
//go:generate mockgen -package=mocks -destination=mock_client.go github.com/tphoney/musicscan/client Client
