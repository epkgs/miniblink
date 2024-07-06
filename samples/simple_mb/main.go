package main

import (
	"github.com/epkgs/miniblink/mb"
)

func main() {

	// 通过添加版本 tag 加载对应的内嵌版本
	if err := mb.LoadLibrary(""); err != nil {
		panic(err)
	}

	mb.Init(nil)

	view := mb.CreateWebWindow(mb.WINDOW_TYPE_POPUP, 0, 0, 0, 1200, 1000)

	mb.LoadURL(view, "https://www.douyin.com")
	mb.ShowWindow(view, true)

	mb.RunMessageLoop()
}
