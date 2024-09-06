package config

func Init() (*Config, error) {

	config, err := InitConfig()
	if err != nil {
		return nil, err
	}
	log, err := InitLog()
	// 初始化日志
	if err != nil {
		return nil, err
	}
	// 初始化mysql
	mysql, err := InitMySQL()
	if err != nil {
		return nil, err
	}
	redis, err := InitRedis()
	if err != nil {
		return nil, err
	}

	config.Log = log
	config.MySQL = mysql
	config.Redis = redis

	return config, nil
}

func New() *Config {
	return &Config{}
}

func (c *Config) Get(key string, defaultVal ...interface{}) interface{} {
	return nil
}

func (c *Config) Set(key string, val interface{}) {

}
