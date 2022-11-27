package plugin

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/types"
	lua "github.com/yuin/gopher-lua"
)

func createPluginContext(l *lua.LState, p *types.Plugin) *lua.LTable {
	t := l.NewTable()

	l.SetField(t, "Name", lua.LString(p.Name))
	l.SetField(t, "Version", lua.LString("v"+p.Version))
	l.SetField(t, "Author", lua.LString(p.Author))

	return t
}
