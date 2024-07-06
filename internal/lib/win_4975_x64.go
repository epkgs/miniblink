//go:build windows && !386 && v4975

package lib

import (
	"embed"
)

const (
	FILE_NAME = "miniblink_4975_x64.dll"
	VERSION   = "4975"
	ARCH      = "x64"
)

//go:embed static/miniblink_4975_x64.dll
var fs embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

func ReadFile() ([]byte, error) {
	return fs.ReadFile("static/" + FILE_NAME)
}
