package config

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string   `mapstructure:"host" json:"host"`
	Port int      `mapstructure:"port" json:"port"`
	Tags []string `mapstructure:"tags" json:"tags"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	Host       string       `mapstructure:"host" json:"host"`
	Port       int          `mapstructure:"port" json:"port"`
	JWTInfo    JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	OssInfo    OssConfig    `mapstructure:"oss" json:"oss"`
}

type OssConfig struct {
	ApiKey      string `mapstructure:"key" json:"key"`
	ApiSecret   string `mapstructure:"secret" json:"secret"`
	Host        string `mapstructure:"host" json:"host"`
	CallBackUrl string `mapstructure:"callback_url" json:"callback_url"`
	UploadDir   string `mapstructure:"upload_dir" json:"upload_dir"`
}
