package config

import "fsc/model"

// FSC 系统配置
type FSCConfig struct {
	FSC FSC `mapstructure:"fsc" json:"fsc" yaml:"fsc"`
	SportCampusConfig SportCampusConfig `mapstructure:"sportCampus" json:"sportCampus" yaml:"sportCampus"`
	Account model.LoginStruct `mapstructure:"account" json:"account" yaml:"account"`
	Panel FSCPanel `mapstructure:"panel" json:"panel" yaml:"panel"`
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
}

// FSC 核心
type FSC struct {
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
}

// 高校体育配置
type SportCampusConfig struct {
	Api string `mapstructure:"api" json:"api" yaml:"api"`
	EncryptionKey string `json:"encryptionKey" yaml:"encryptionKey" mapstructure:"encryptionKey"`
}

// 主菜单界面
type FSCPanel struct {
	RedDot string `json:"red-dot"`
	GreenDot string `json:"green-dot"`
}

type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile" yaml:"log-file"`
	Stdout  string `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
	File    string `mapstructure:"file" json:"file" yaml:"file"`
}