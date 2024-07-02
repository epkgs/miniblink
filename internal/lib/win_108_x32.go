//go:build windows && 386 && v108

package lib

import (
	"embed"
)

const (
	FILE_NAME = "miniblink_108_x32.dll"
	VERSION   = "108"
	ARCH      = "x32"
)

//go:embed static/miniblink_108_x32.dll
var fs embed.FS

func ReadFile() ([]byte, error) {
	return fs.ReadFile("static/" + FILE_NAME)
}
