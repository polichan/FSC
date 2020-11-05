package config

// FSC 系统配置
type FSCConfig struct {
	Debug bool `mapstructure:"debug" json:"debug" yaml:"fsc-debug"`
	SportCampusConfig SportCampusConfig `mapstructure:"sportCampus" json:"sportCampus" yaml:"sportCampus"`
}

// 高校体育配置
type SportCampusConfig struct {
	Api string `mapstructure:"api" json:"api" yaml:"api"`
}