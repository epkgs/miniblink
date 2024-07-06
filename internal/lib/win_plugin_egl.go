//go:build windows && plugin_egl

package lib

import "embed"

//go:embed static/plugins/egl/libEGL.dll
var libEGL embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

//go:embed static/plugins/egl/libGLESv2.dll
var libGLESv2 embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

func init() {
	plugins["egl/libEGL.dll"] = libEGL
	plugins["egl/libGLESv2.dll"] = libGLESv2
}
