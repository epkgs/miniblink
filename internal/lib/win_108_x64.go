//go:build windows && !386 && v108

package lib

import (
	"embed"
)

const (
	FILE_NAME = "miniblink_108_x64.dll"
	VERSION   = "108"
	ARCH      = "x64"
)

//go:embed static/miniblink_108_x64.dll
var fs embed.FS

func ReadFile() ([]byte, error) {
	return fs.ReadFile("static/" + FILE_NAME)
}
