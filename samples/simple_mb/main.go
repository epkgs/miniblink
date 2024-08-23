package main

import (
	"fmt"

	"github.com/epkgs/miniblink/mb"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("PANIC: %v", err)
		}
	}()

	// 通过添加版本 tag 加载对应的内嵌版本
	if err := mb.LoadLibrary(""); err != nil {
		panic(err)
	}

	mb.Init(nil)

	view := mb.CreateWebWindow(mb.WINDOW_TYPE_POPUP, 0, 0, 0, 1200, 1000)

	// TODO: 下载回调崩溃
	// mb.OnDownloadInBlinkThread(view, func(webView mb.WebView, expectedContentLength uintptr, url, mime, disposition string, job mb.NetJob, dataBind *mb.NetJobDataBind) mb.DownloadOpt {
	// 	fmt.Printf("mb.OnDownloadInBlinkThread: %s\n", url)
	// 	return 1
	// })

	mb.LoadURL(view, "https://tim.qq.com/download.html")
	mb.ShowWindow(view, true)

	mb.RunMessageLoop()
}
