//go:build windows && 386 && plugin_ffmpeg

package lib

import "embed"

//go:embed static/plugins/ffmpeg/ffmpeg.dll
var ffmpeg embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

func init() {
	plugins["ffmpeg/ffmpeg.dll"] = ffmpeg
}
