package util

import (
	"crypto/md5"
	"encoding/hex"
	"fsc/global"
)

func GetMd5(b []byte) string  {
	b = []byte(global.FSC_CONFIG.SportCampusConfig.EncryptionKey + "data" + string(b))
    h := md5.New()
    h.Write(b)
    return hex.EncodeToString(h.Sum(nil))
}