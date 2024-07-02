package main

import (
	"github.com/epkgs/miniblink/wke"
)

func main() {

	if err := wke.LoadLibrary(""); err != nil {
		panic(err)
	}

	wke.Init()

	view := wke.CreateWebWindow(wke.WINDOW_TYPE_POPUP, 0, 0, 0, 800, 600)

	wke.LoadURL(view, "https://www.baidu.com")

	wke.ShowWindow(view, true)

	wke.RunMessageLoop()
}
