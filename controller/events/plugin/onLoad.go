package plugin

import (
	"github.com/sirupsen/logrus"

	pController "github.com/AirplaneMC/airplane-plugins-api/controller"
	lua "github.com/yuin/gopher-lua"
)

func OnLoad(l *lua.LState, self, loadFunc lua.LValue, p *pController.Plugin, log *logrus.Logger) {
	err := l.CallByParam(lua.P{
		Fn: loadFunc,

		NRet:    1,
		Protect: true,
	}, self, createPluginContext(l, p))
	if err != nil {
		log.Warningf("[%v] An error occurred while trying to call the event function. Error: %v\n", p.Name, err)
	}
}
