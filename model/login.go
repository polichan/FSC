package model

import uuid "github.com/satori/go.uuid"

/**
用户登录模型
 */
type LoginStruct struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	PhoneType string `json:"type"`
	Info uuid.UUID `json:"info"`
}