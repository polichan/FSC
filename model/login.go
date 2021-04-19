package model

import uuid "github.com/satori/go.uuid"

/**
用户登录模型
*/
type LoginStruct struct {
	Mobile    string    `json:"mobile"  yaml:"mobile"`
	Password  string    `json:"password"  yaml:"password"`
	PhoneType string    `json:"type"  yaml:"phoneType"`
	Info      uuid.UUID `json:"info" yaml:"uuid"`
}
