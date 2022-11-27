package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) abortBreaking(l *lua.LState) int {
	p.player.AbortBreaking()

	return 0
}
