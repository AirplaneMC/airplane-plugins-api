package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) hasCooldown(l *lua.LState) int {
	return 0
}
