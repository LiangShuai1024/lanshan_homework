package config

type Configuration struct {
	Ec       Ec       `mapstructure:"ec" json:"ec" yaml:"ec"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"` //用标签指定序列化的映射规则

}
