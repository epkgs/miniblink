//go:build windows
// +build windows

package lib

import "golang.org/x/sys/windows"

func loadLibrary(libfile string) (handle uintptr, err error) {
	h, e := windows.LoadLibrary(libfile)
	return uintptr(h), e
}

func release(handle uintptr) (err error) {
	return windows.FreeLibrary(windows.Handle(handle))
}
