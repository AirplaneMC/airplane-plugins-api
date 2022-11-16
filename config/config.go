package config

import "gopkg.in/yaml.v3"

type Config struct {
	Name   string `yaml:"name"`
	Author string `yaml:"author"`

	LUA_V string `yaml:"lua_v"`
	API_V string `yaml:"api_v"`
}

func (c *Config) ReadConfig(data []byte) error {
	return yaml.Unmarshal(data, &c)
}
