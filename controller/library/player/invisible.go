package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) invisible(l *lua.LState) int {
	return 0
}
