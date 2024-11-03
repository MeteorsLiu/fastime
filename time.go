package fastime

import (
	_ "unsafe"
)

var Now = Cputicks

//go:linkname Nanotime runtime.nanotime
func Nanotime() int64

//go:linkname Cputicks runtime.cputicks
func Cputicks() int64

func timeElasped(fn func()) int64 {
	now := Nanotime()
	fn()
	return Nanotime() - now
}

func init() {
	nt := timeElasped(func() {
		Nanotime()
	})
	ct := timeElasped(func() {
		Cputicks()
	})
	// if the VDSO is slow, don't use
	if nt <= ct+100 {
		Now = Nanotime
	}
}
