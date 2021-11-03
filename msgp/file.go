//go:build (linux || darwin || dragonfly || freebsd || netbsd || openbsd) && !appengine
// +build linux darwin dragonfly freebsd netbsd openbsd
// +build !appengine

package msgp

// MarshalSizer is the combination
// of the Marshaler and Sizer
// interfaces.
type MarshalSizer interface {
	Marshaler
	Sizer
}
