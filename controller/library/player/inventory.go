package player

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/library/inventory"
	lua "github.com/yuin/gopher-lua"
)

func (p *Player) inventory(l *lua.LState) int {
	inv := p.player.Inventory()

	l.Push(inventory.GetInventory(inv, l))

	return 1
}
