package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type C struct {
	URL      string
	Username string
	Password string
	Site     string
}

func ReadConfig(fileName string) (conf *C, err error) {
	_, err = toml.DecodeFile(fileName, &conf)
	if err != nil {
		return nil, fmt.Errorf("can't read config file: %w", err)
	}

	return conf, nil
}
