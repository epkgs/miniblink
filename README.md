# [MINIBLINK](https://github.com/weolar/miniblink49) 接口纯 GO 绑定

## 简介

- 纯GO实现的MINIBLINK接口绑定，无须CGO
- 内嵌miniblink，可指定版本，如：`go build -tags 'v108' .` 支持一下版本：
  - v4975
  - v108

## 安装

```bash
go get github.com/epkgs/miniblink
```

## 示例

```go
package main

import (
   "github.com/epkgs/miniblink/mb"
)

func main() {

    // LoadLibrary 加载 dll 文件, 空字符串根据版本tags和当前环境加载内嵌的 dll 文件
    if err := mb.LoadLibrary(""); err != nil {
        panic(err)
    }

    mb.Init(nil)

    view := mb.CreateWebWindow(mb.WINDOW_TYPE_POPUP, 0, 0, 0, 800, 600)

    mb.LoadURL(view, "https://www..baidu.com")

    mb.ShowWindow(view, true)

    mb.RunMessageLoop()
}
```
