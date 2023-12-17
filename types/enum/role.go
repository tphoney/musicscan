// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package enum

import "encoding/json"

// Role defines the member role.
type Role int

// Role enumeration.
const (
	RoleDeveloper Role = iota
	RoleAdmin
)

const DEVELOPER = "developer"
const ADMIN = "admin"

// String returns the Role as a string.
func (e Role) String() string {
	switch e {
	case RoleDeveloper:
		return DEVELOPER
	case RoleAdmin:
		return ADMIN
	default:
		return DEVELOPER
	}
}

// MarshalJSON marshals the Type as a JSON string.
func (e Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value.
func (e *Role) UnmarshalJSON(b []byte) error {
	var v string
	_ = json.Unmarshal(b, &v)
	switch v {
	case ADMIN:
		*e = RoleAdmin
	case DEVELOPER:
		*e = RoleDeveloper
	default:
		*e = RoleDeveloper
	}
	return nil
}
