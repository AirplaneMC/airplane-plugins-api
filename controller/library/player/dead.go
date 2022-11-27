package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) dead(l *lua.LState) int {
	return 0
}
