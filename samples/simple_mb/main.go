package main

import (
	"github.com/epkgs/miniblink/mb"
)

func main() {

	if err := mb.LoadLibrary(""); err != nil {
		panic(err)
	}

	mb.Init(nil)

	view := mb.CreateWebWindow(mb.WINDOW_TYPE_POPUP, 0, 0, 0, 800, 600)

	mb.LoadURL(view, "https://www..baidu.com")
	mb.ShowWindow(view, true)

	mb.RunMessageLoop()
}
