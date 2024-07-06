//go:build windows && !386 && plugin_pdfium

package lib

import "embed"

//go:embed static/plugins/pdfium_x64.dll
var pdfium embed.FS // 不能直接嵌入为二进制，会被杀毒拦截

func init() {
	plugins["pdfium_x64.dll"] = pdfium
}
