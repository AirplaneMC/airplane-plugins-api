package player

import (
	"github.com/sirupsen/logrus"

	pController "github.com/AirplaneMC/airplane-plugins-api/controller"
	"github.com/AirplaneMC/airplane-plugins-api/controller/library/player"
	dPlayer "github.com/df-mc/dragonfly/server/player"
	lua "github.com/yuin/gopher-lua"
)

func OnJoin(l *lua.LState, self, loadFunc lua.LValue, p *pController.Plugin, dPlayer *dPlayer.Player, log *logrus.Logger) {
	err := l.CallByParam(lua.P{
		Fn: loadFunc,

		NRet:    1,
		Protect: true,
	}, self, player.GetPlayer(l, dPlayer))
	if err != nil {
		log.Warningf("[%v] An error occurred while trying to call the event function. Error: %v\n", p.Name, err)
	}
}
