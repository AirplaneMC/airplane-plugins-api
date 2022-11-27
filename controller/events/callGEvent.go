package events

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller"
	"github.com/AirplaneMC/airplane-plugins-api/controller/events/player"
	"github.com/AirplaneMC/airplane-plugins-api/controller/events/plugin"
	dPlayer "github.com/df-mc/dragonfly/server/player"
	"github.com/sirupsen/logrus"
)

func CallOnLoadPE(log *logrus.Logger) {
	for _, e := range controller.GEvents[controller.PLUGIN_EVENTS]["LOAD"] {
		log.Infof("Loading plugin %v v%v by %v", e.Plugin.Name, e.Plugin.Version, e.Plugin.Author)
		plugin.OnLoad(controller.Plugin_VM, e.Self, e.Function, e.Plugin, log)
		log.Infof("Plugin %v v%v by %v loaded successfully!", e.Plugin.Name, e.Plugin.Version, e.Plugin.Author)
	}
}

func CallOnJoinPE(log *logrus.Logger, p *dPlayer.Player) {
	for _, e := range controller.GEvents[controller.PLAYER_EVENTS]["JOIN"] {
		player.OnJoin(controller.Plugin_VM, e.Self, e.Function, e.Plugin, p, log)
	}
}
