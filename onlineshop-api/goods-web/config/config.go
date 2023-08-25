package config

type GoodsSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string   `mapstructure:"host" json:"host"`
	Port int      `mapstructure:"port" json:"port"`
	Tags []string `mapstructure:"tags" json:"tags"`
}

type ServerConfig struct {
	Name        string         `mapstructure:"name" json:"name"`
	Host        string         `mapstructure:"host" json:"host"`
	Port        int            `mapstructure:"port" json:"port"`
	UserSrvInfo GoodsSrvConfig `mapstructure:"goods_srv" json:"goods_srv"`
	JWTInfo     JWTConfig      `mapstructure:"jwt" json:"jwt"`
	ConsulInfo  ConsulConfig   `mapstructure:"consul" json:"consul"`
}
