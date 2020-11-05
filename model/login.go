package model

/**
用户登录模型
 */
type LoginStruct struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	PhoneType string `json:"type"`
}