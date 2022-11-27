package player

import (
	"time"

	e "github.com/AirplaneMC/airplane-plugins-api/controller/library/effect"
	"github.com/df-mc/dragonfly/server/entity/effect"
	lua "github.com/yuin/gopher-lua"
)

func (p *Player) addEffect(l *lua.LState) int {
	eName := l.ToString(1)
	eLvl := l.ToInt(2)
	eTime := l.ToInt(3)

	eType, err := e.GetEffect(eName)
	if err != nil {
		l.Push(lua.LString(err.Error()))
		return 1
	}

	p.player.AddEffect(effect.New(
		eType, eLvl, time.Duration(eTime)*time.Second,
	))

	return 0
}
