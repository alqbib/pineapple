package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	App `yaml:"app"`
}

type App struct {
	Name    string `yaml:"name"`
	Port    string `yaml:"port"`
	RunMode string `yaml:"run_mode"`
}

func InitConfig(cfgFile string) (*Config, error) {
	cfg := &Config{}
	f, err := os.Open(cfgFile)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
