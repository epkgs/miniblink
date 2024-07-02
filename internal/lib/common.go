package lib

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/ebitengine/purego"
)

var HMODULE uintptr

// LoadLibrary 尝试加载指定路径下的 DLL 文件或嵌入的 DLL。
//
// 如果成功加载 DLL，则返回 nil；否则返回错误信息。
//
// 参数 libfile 表示 DLL 文件的路径，可以为空字符串。
// 如果 libfile 不为空，则优先尝试直接加载该 DLL 文件。
// 如果 libfile 为空或加载 DLL 文件失败，则尝试加载嵌入的 DLL。
func LoadLibrary(libfile string) error {

	if libfile != "" {

		// 尝试直接加载 DLL
		if h, err := loadLibrary(libfile); err == nil {
			HMODULE = h
			return nil
		}
	}

	return LoadEmbedLibrary()
}

type LazyFunc[F any] struct {
	Name     string
	Call     F
	register sync.Once
}

func (lf *LazyFunc[F]) LoadOnce() {

	lf.register.Do(func() {
		if HMODULE == 0 {
			panic("HMODULE is not initialized")
		}
		purego.RegisterLibFunc(&(lf.Call), HMODULE, lf.Name)
	})
}

// LoadEmbedLibrary 函数用于加载嵌入的静态链接库文件
// 如果加载成功，则返回nil；否则返回错误信息
func LoadEmbedLibrary() error {
	targetFile, err := generateLibFile()
	if err != nil {
		return err
	}

	if h, err := loadLibrary(targetFile); err == nil {
		HMODULE = h
		return nil
	} else {
		return err
	}
}

func Release() (err error) {
	if HMODULE != 0 {
		return release(HMODULE)
	}

	return nil
}

func generateLibFile() (string, error) {

	tmpDir := filepath.Join(os.TempDir(), "miniblink")
	targetFile := filepath.Join(tmpDir, FILE_NAME)

	// 检查是否已有释放的 DLL
	if _, err := os.Stat(targetFile); err == nil {
		return targetFile, nil
	}

	// 创建临时文件夹
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return "", errors.New("无法创建临时文件夹，err: " + err.Error())
	}

	// 读取内嵌文件
	data, err := ReadFile()
	if err != nil {
		return "", err
	}

	// 创建dll文件
	newFile, err := os.Create(targetFile)
	if err != nil {
		return "", errors.New("无法创建静态链接库文件，err: " + err.Error())
	}
	defer newFile.Close()

	n, err := newFile.Write(data)
	if err != nil {
		return "", errors.New("写入dll文件失败，err: " + err.Error())
	}
	if n != len(data) {
		return targetFile, errors.New("写入校验失败")
	}

	return targetFile, nil
}
