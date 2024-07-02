//go:build !windows
// +build !windows

package lib

import (
	"github.com/ebitengine/purego"
)

func loadLibrary(libfile string) (handle uintptr, err error) {
	return purego.Dlopen(libfile, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
