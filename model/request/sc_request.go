package request

type SCRequestStruct struct {
	Sign string `json:"sign" mapstructure:"sign"`
	Data string `json:"data" mapstructure:"data"`
}
