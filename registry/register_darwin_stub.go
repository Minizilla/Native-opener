//go:build !darwin

package registry

func RegisterOnMac(protocol, progPath, args string) {
	// empty declaration, for compilation / typing
}
