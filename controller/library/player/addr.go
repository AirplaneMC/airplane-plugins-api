package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) addr(l *lua.LState) int {
	addr := p.player.Addr().String()

	l.Push(lua.LString(addr))
	return 1
}
