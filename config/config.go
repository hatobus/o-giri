package config

import "github.com/kelseyhightower/envconfig"

type MySQLConfig struct {
	Host string `default:"localhost" required:"true"`
	Port string `default:"3306" required:"true"`
	User string `default:"root" required:"true"`
	Password string `default:"" required:"true"`
	DataBase string `default:"ogiri" required:"true"`
}

type Config struct {
	MySQL MySQLConfig `envconfig:"MYSQL" required:"true"`
}

func Init() (*Config, error) {
	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
