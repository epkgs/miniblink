//go:build !windows

package lib

const (
	FILE_NAME = "non-embed"
	VERSION   = "108"
	ARCH      = "x32"
)

func ReadFile() ([]byte, error) {
	panic("请选择嵌入的链接库版本，如：go build -tags 'v108' . ")
}
