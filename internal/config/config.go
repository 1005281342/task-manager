package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var conf []byte

func New() Config {
	var (
		c   Config
		err error
	)
	if err = yaml.Unmarshal(conf, &c); err != nil {
		panic(err)
	}
	return c
}

type Config struct {
	Gorm Gorm `json:"gorm"`
}

type Gorm struct {
	Driver string `yaml:"driver"`
	Dsn    string `yaml:"dsn"`
}
