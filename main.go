package main

import (
	"flag"
	"fsc/controller"
	"fsc/core"
	"fsc/global/response"
	"os"
)

func main() {
	Init()
	if isArgEmpty() {
		response.FailWithMessage("缺失参数，通过 -h 以获取帮助", nil)
		return
	}
}

func Init()  {
	core.Init()
}

func Login()  {
	var(
		mobile string
		password string
		phoneType string
	)
	flag.StringVar(&mobile, "m", "", "高校体育登录手机号")
	flag.StringVar(&password, "p", "", "高校体育登录密码")
	flag.StringVar(&phoneType, "pt", "", "登录设备类型：例：iPhone X")
	flag.Parse()
	controller.Login(mobile, password, phoneType)
}

func isArgEmpty()bool  {
	return len(os.Args) <= 1
}