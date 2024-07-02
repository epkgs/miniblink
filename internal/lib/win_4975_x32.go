//go:build windows && 386 && v4975

package lib

import (
	"embed"
)

const (
	FILE_NAME = "miniblink_4975_x32.dll"
	VERSION   = "4975"
	ARCH      = "x32"
)

//go:embed static/miniblink_4975_x32.dll
var fs embed.FS

func ReadFile() ([]byte, error) {
	return fs.ReadFile("static/" + FILE_NAME)
}
