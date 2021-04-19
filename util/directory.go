package util

import (
	"fsc/global"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.FSC_LOG.Debug("create directory ", v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.FSC_LOG.Error("create directory", v, " error:", err)
			}
		}
	}
	return err
}
