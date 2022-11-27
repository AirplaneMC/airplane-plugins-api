package player

import lua "github.com/yuin/gopher-lua"

func (p *Player) addExperience(l *lua.LState) int {
	exp := l.ToInt(1)
	n := p.player.AddExperience(exp)

	l.Push(lua.LNumber(n))

	return 1
}
