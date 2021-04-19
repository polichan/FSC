package util

import (
	"encoding/json"
	"reflect"
)

func Transfer(beConverted, converted interface{}) error {
	t := reflect.TypeOf(beConverted)
	var bt []byte
	var err error
	if t.Name() == "string" {
		bt = []byte(beConverted.(string))
	} else {
		bt, err = json.Marshal(beConverted)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(bt, converted)
	if err != nil {
		return err
	}
	return nil
}
