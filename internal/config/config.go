package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port      int      `yaml:"port"`
	Interface string   `yaml:"interface"`
	Interval  string   `yaml:"interval"`
	URLs      []string `yaml:"urls"`
}

func ReadConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	// Set default values if not provided
	if config.Port == 0 {
		config.Port = 9090
	}
	if config.Interface == "" {
		config.Interface = "0.0.0.0"
	}

	return config, nil
}
