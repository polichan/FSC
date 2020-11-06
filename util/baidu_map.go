package util

import (
	"fmt"
	"fsc/global"
	"github.com/menduo/gobaidumap"
)

func GetSchoolLocation(school string) (err error) {
	mp := make(map[string]string, 0)
	mp["address"] = school
	mp["output"] = "json"

	bc := gobaidumap.NewBaiduMapClient(global.FSC_CONFIG.BaiduMapConfig.AccessKey)
	addressToGEO, err := bc.GetGeoViaAddress(school)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("从地址到坐标 - All", addressToGEO)
		fmt.Println("从地址到坐标 - Lat", addressToGEO.Result.Location.Lat)
		fmt.Println("从地址到坐标 - Lng", addressToGEO.Result.Location.Lng)
	}
	return err
}