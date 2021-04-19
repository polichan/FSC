package main

import (
	"fsc/core"
	"fsc/gui"
)

func main() {
	// 注册核心
	core.Init()
	// 展示面板
	gui.ShowPanel()
}
