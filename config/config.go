package config

import "github.com/kelseyhightower/envconfig"

type MySQLConfig struct {
	Host string `default:"localhost" required:"true"`
	Port string `default:"3306" required:"true"`
	User string `default:"root" required:"true"`
	Password string `default:"" required:"true"`
	DataBase string `default:"ogiri" required:"true"`
}

type OogiriConfig struct {
	AnswerSection int `default:"3" require:"true"`
	VoteSection int `default:"3" require:"true"`
	QuestionSection int `default:"3" require:"true"`
}

type Config struct {
	MySQL MySQLConfig `envconfig:"MYSQL" required:"true"`
	Oogiri OogiriConfig `envconfig:"OOGIRI" require:"true"`
}

func Init() (*Config, error) {
	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
