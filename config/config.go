package config

// FSC 系统配置
type FSCConfig struct {
	FSC FSC `mapstructure:"fsc" json:"fsc" yaml:"fsc"`
	SportCampusConfig SportCampusConfig `mapstructure:"sportCampus" json:"sportCampus" yaml:"sportCampus"`
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