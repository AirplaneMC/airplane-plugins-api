package types

import (
	"gopkg.in/yaml.v3"
)

type Plugin struct {
	Name    string `yaml:"name"`
	Author  string `yaml:"author"`
	Main    string `yaml:"main"`
	Version string `yaml:"version"`

	LUA_V string `yaml:"lua_v"`
	API_V string `yaml:"api_v"`
}

func (p *Plugin) ReadConfig(data []byte) error {
	return yaml.Unmarshal(data, &p)
}
