package service

import (
	"fmt"
	"fsc/model"
	"fsc/util"
)


func Login(loginModel *model.LoginStruct)(err error)  {
	code ,s, err := util.Get("https://www.baidu.com").Exec().String()
	if err != nil {
			fmt.Printf("%d %s", code, s)
	}
	return nil
}