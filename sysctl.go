// Package sysctl provides functions wrapping the sysctl interface.
package sysctl

import "runtime"

// Get returns a sysctl from a given key.
func Get(key string) (string, error) {
	var value string
	var err error
	switch runtime.GOOS {
	case `darwin`:
		value, err = DarwinGet(key)
	case `linux`:
		value, err = LinuxGet(key)
	}
	return value, err
}
