package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"default"`
	Port    string `default:"9090"`
	Logger  struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"default.log"`
	}
	DB struct {
		Use      string `default:"postgres"`
		Enabled  bool   `default:"true"`
		Host     string `default:"localhost"`
		Port     string `default:"5432"`
		UserName string `default:"root"`
		Password string `default:"root"`
		Database string `default:"default"`
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "pkg/config/config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
