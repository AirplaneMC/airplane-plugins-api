package inventory

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/library/item"
	"github.com/df-mc/dragonfly/server/item/inventory"
	lua "github.com/yuin/gopher-lua"
)

type Inventory struct {
	inventory *inventory.Inventory
}

func (i *Inventory) addItem(l *lua.LState) int {
	iName := l.ToString(1)
	count := l.ToInt(2)

	item, err := item.GetStack(iName, count)
	if err != nil {
		l.Push(lua.LNumber(0))
		l.Push(lua.LString(err.Error()))
		return 2
	}
	n, err := i.inventory.AddItem(item)
	if err != nil {
		l.Push(lua.LNumber(n))
		l.Push(lua.LString(err.Error()))
		return 2
	}

	l.Push(lua.LNumber(n))
	l.Push(lua.LNil)

	return 2
}
