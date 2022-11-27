package plugin

import (
	pController "github.com/AirplaneMC/airplane-plugins-api/controller"
	lua "github.com/yuin/gopher-lua"
)

func createPluginContext(l *lua.LState, p *pController.Plugin) *lua.LTable {
	t := l.NewTable()

	l.SetField(t, "Name", lua.LString(p.Name))
	l.SetField(t, "Version", lua.LString("v"+p.Version))
	l.SetField(t, "Author", lua.LString(p.Author))

	return t
}
