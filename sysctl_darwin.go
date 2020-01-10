package sysctl

import "syscall"

// DarwinGet returns a sysctl from a given key.
func DarwinGet(key string) (string, error) {
	return syscall.Sysctl(key)
}
