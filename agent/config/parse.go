package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func ParseConfig(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}
	cfg := new(Config)

	_, err = toml.Decode(string(data), cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
