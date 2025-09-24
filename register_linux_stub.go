//go:build !linux

package main

func RegisterOnLinux(protocol, progPath, args string) {
	// empty declaration, for compilation / typing
}
