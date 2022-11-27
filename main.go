package main

import (
	"fmt"
	"os"

	"github.com/AirplaneMC/airplane-plugins-api/controller/events"
	"github.com/AirplaneMC/airplane-plugins-api/loader"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"

	playerEvents "github.com/AirplaneMC/airplane-plugins-api/controller/events/player"
	serverPlayer "github.com/df-mc/dragonfly/server/player"
)

func main() {
	log := logrus.New()

	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	l, err := loader.Load(log)
	if err != nil {
		log.Fatalln("An error occurred while trying to load plugins. Error:", err)
		return
	}

	srv.Listen()

	for srv.Accept(func(p *serverPlayer.Player) {
		p.Handle(&playerEvents.PlayerHandler{})
		events.CallOnJoinPE(log, l, p)
	}) {
	}
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	c.Server.Name = "Airplane Server"
	c.World.Folder = "worlds/world"

	var zero server.Config

	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}

		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}

		return c.Config(log)
	}

	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}

	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}

	return c.Config(log)
}
