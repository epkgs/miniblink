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
var fs embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

func ReadFile() ([]byte, error) {
	return fs.ReadFile("static/" + FILE_NAME)
}
