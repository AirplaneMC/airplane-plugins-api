package events

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/events/plugin"
	"github.com/sirupsen/logrus"

	playerEvents "github.com/AirplaneMC/airplane-plugins-api/controller/events/player"
	serverPlayer "github.com/df-mc/dragonfly/server/player"
	lua "github.com/yuin/gopher-lua"
)

func CallOnLoadPE(log *logrus.Logger, l *lua.LState) {
	for _, e := range events[PLUGIN_EVENTS]["LOAD"] {
		log.Infof("Loading plugin %v v%v by %v\n", e.Plugin.Name, e.Plugin.Version, e.Plugin.Author)
		plugin.OnLoad(log, l, e.Self, e.Function, e.Plugin)
		log.Infof("Plugin %v v%v by %v loaded successfully!\n", e.Plugin.Name, e.Plugin.Version, e.Plugin.Author)
	}
}

func CallOnJoinPE(log *logrus.Logger, l *lua.LState, p *serverPlayer.Player) {
	for _, e := range events[PLAYER_EVENTS]["JOIN"] {
		playerEvents.OnJoin(l, e.Self, e.Function, e.Plugin, p, log)
	}
}
