package model

// 登录成功之后的用户信息
type UserStruct struct {
	UserId   string `json:"userid"`
	School   string `json:"school"`
	SchoolId string `json:"schoolId"`
	Username string `json:"username"`
	Goal     string `json:"goal"`
	Surplus  string `json:"surplus"`
	Teacher  string `json:"teacher"`
	Course   string `json:"course"`
	UToken   string `json:"utoken"`
}
