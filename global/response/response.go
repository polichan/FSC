package response

import (
	"fmt"
	"fsc/global"
)

func OkWithMessage(message string)  {
	fmt.Printf("「FSC」成功：%s", message)
}

func FailWithMessage(message string, err error)  {
	if global.FSC_CONFIG.Debug {
		fmt.Printf("「FSC」失败：%s 。异常：%s", message, fmt.Sprintf("%s", err))
	}
	fmt.Printf("「FSC」失败：%s 。", message)
}