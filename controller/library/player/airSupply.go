package player

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func (p *Player) airSupply(l *lua.LState) int {
	t := p.player.AirSupply()

	l.Push(timeToTable(l, t))
	return 1
}

func timeToTable(l *lua.LState, t time.Duration) *lua.LTable {
	table := l.NewTable()

	l.SetField(table, "Microseconds", lua.LNumber(t.Microseconds()))
	l.SetField(table, "Seconds", lua.LNumber(t.Seconds()))
	l.SetField(table, "Minutes", lua.LNumber(t.Minutes()))
	l.SetField(table, "Hours", lua.LNumber(t.Hours()))

	return table
}
