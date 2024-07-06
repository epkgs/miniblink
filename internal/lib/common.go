package lib

import (
	"embed"
	"errors"
	"os"
	"path/filepath"
)

type DLL struct {
	Name   string
	Handle uintptr
}

func (d *DLL) Release() error {
	if Dll == nil || Dll.Handle == 0 {
		return nil
	}

	return release(Dll.Handle)
}

var Dll *DLL
var tmpDir = filepath.Join(os.TempDir(), "miniblink")

var plugins map[string]embed.FS = make(map[string]embed.FS)

// LoadLibrary 尝试加载指定路径下的 DLL 文件或嵌入的 DLL。
//
// 如果成功加载 DLL，则返回 nil；否则返回错误信息。
//
// 参数 libfile 表示 DLL 文件的路径，可以为空字符串。
// 如果 libfile 不为空，则优先尝试直接加载该 DLL 文件。
// 如果 libfile 为空或加载 DLL 文件失败，则尝试加载嵌入的 DLL。
func LoadLibrary(libfile string) (*DLL, error) {

	if Dll != nil && Dll.Name == libfile {
		return Dll, nil
	}

	if libfile != "" {
		// 尝试直接加载 DLL
		if h, err := loadLibrary(libfile); err == nil {
			Dll = &DLL{Name: libfile, Handle: h}
			return Dll, nil
		}
	}

	return loadEmbedLibrary()
}

func loadEmbedLibrary() (*DLL, error) {
	targetFile, err := extractMbDLL()
	if err != nil {
		Dll = nil
		return nil, err
	}

	if Dll != nil && Dll.Name == targetFile {
		return Dll, nil
	}

	h, err := loadLibrary(targetFile)
	if err != nil {
		Dll.Release()
		Dll = nil
		return nil, err
	}

	go extractPluginsDLL()

	Dll = &DLL{Name: targetFile, Handle: h}
	return Dll, nil
}

func Release() (err error) {
	return Dll.Release()
}

func extractMbDLL() (string, error) {

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

func extractPluginsDLL() error {

	pluginDir := filepath.Join(tmpDir, "plugins")

	for p, fs := range plugins {

		// 使用闭包，使 defer 每个循环都会执行
		err := func() error {

			targetFile := filepath.Join(pluginDir, p)
			// 检查是否已有释放的 DLL
			if _, err := os.Stat(targetFile); err == nil {
				return nil
			}

			dir, _ := filepath.Split(targetFile)

			// 创建临时文件夹
			if err := os.MkdirAll(dir, 0755); err != nil {
				return errors.New("无法创建临时文件夹，err: " + err.Error())
			}

			// 读取内嵌文件
			data, err := fs.ReadFile("static/plugins/" + p)
			if err != nil {
				return err
			}

			// 创建dll文件
			newFile, err := os.Create(targetFile)
			if err != nil {
				return errors.New("无法创建静态链接库文件，err: " + err.Error())
			}
			defer newFile.Close()

			n, err := newFile.Write(data)
			if err != nil {
				return errors.New("写入dll文件失败，err: " + err.Error())
			}
			if n != len(data) {
				return errors.New("写入校验失败")
			}
			return nil
		}()

		if err != nil {
			return err
		}

	}

	return nil
}
