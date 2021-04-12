// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package mutex provides a global mutex.
package mutex

import "sync"

var m sync.RWMutex

// RLock locks the global mutex for reads.
func RLock() { m.RLock() }

// RUnlock unlocks the global mutex.
func RUnlock() { m.RUnlock() }

// Lock locks the global mutex for writes.
func Lock() { m.Lock() }

// Unlock unlocks the global mutex.
func Unlock() { m.Unlock() }
