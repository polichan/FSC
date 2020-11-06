package main

import (
	"fsc/core"
	"fsc/gui"
)

func main() {
	// 注册核心
	InitCore()
	// 展示面板
	gui.ShowPanel()
}

func InitCore()  {
	core.Init()
}