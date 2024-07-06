//go:build windows && plugin_swiftshader

package lib

import "embed"

//go:embed static/plugins/swiftshader/libEGL.dll
var swiftshader_libEGL embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

//go:embed static/plugins/swiftshader/libGLESv2.dll
var swiftshader_libGLESv2 embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

func init() {
	plugins["swiftshader/libEGL.dll"] = swiftshader_libEGL
	plugins["swiftshader/libGLESv2.dll"] = swiftshader_libGLESv2
}
