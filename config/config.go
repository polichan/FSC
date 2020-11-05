package config

type FSCConfig struct {
	Debug bool `yaml:"debug"`
}

/**
高校体育接口配置
 */
type SportCampusInterfaceConfig struct {
	Api string `yaml:"api"`
	Header SportCampusRequestHeader `yaml:"sportCampus"`
}

/**
高校体育请求头
 */
type SportCampusRequestHeader struct {

}