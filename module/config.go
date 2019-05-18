package module

type redisConfig struct {
	Addr string `yaml:"addr"`
	DB int32 `yaml:"db"`
}

type Config struct {
	Redis struct {
		Message redisConfig
		Common redisConfig
		Session redisConfig
	}
	Mongodb struct{
		DBhost string `yaml:"host"`
		Authdb string `yaml:"auth"`
		Authuser string `yaml:"user"`
		Authpass string `yaml:"pass"`
		Poollimit int32 `yaml:"poollimit"`
	}
}
