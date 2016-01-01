package config

type Config struct {
	DbConnStr string
}

var conf *Config

func Instance() *Config {
	if conf == nil {
		conf = &Config{
			DbConnStr: "test:test@/livermore?parseTime=true&loc=Local",
		}
	}
	return conf
}
