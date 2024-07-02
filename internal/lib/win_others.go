//go:build windows && !v108 && !v4975

package lib

const (
	FILE_NAME = "non-embed"
	VERSION   = "108"
	ARCH      = "x32"
)

func ReadFile() ([]byte, error) {
	panic("请选择嵌入的链接库版本，如：go build -tags 'v108' . ")
}
