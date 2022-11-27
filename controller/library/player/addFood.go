package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) addFood(l *lua.LState) int {
	n := l.ToInt(1)

	p.player.AddFood(n)
	return 0
}
