//go:build !windows
// +build !windows

package acl

import (
	"os"

	"github.com/hectane/go-acl/api"
)

// Chmod is os.Chmod.
var Chmod = os.Chmod

func Apply(name string, replace, inherit bool, entries ...api.ExplicitAccess) error {
	panic("Not Implemented under non Windows Systems")
}

func GrantName(accessPermissions uint32, name string) api.ExplicitAccess {
	panic("Not implemented")
}
