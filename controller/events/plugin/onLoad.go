package plugin

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/types"
	"github.com/sirupsen/logrus"

	lua "github.com/yuin/gopher-lua"
)

func OnLoad(log *logrus.Logger, l *lua.LState, self, loadFunc lua.LValue, p *types.Plugin) {
	err := l.CallByParam(lua.P{
		Fn: loadFunc,

		NRet:    1,
		Protect: true,
	}, self, createPluginContext(l, p))
	if err != nil {
		log.Warningf("[%v] An error occurred while trying to call the event function. Error: %v\n", p.Name, err)
	}
}
