package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) absorption(l *lua.LState) int {
	n := p.player.Absorption()

	l.Push(lua.LNumber(n))

	return 1
}
