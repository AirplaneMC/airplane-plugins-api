package inventory

import (
	"github.com/df-mc/dragonfly/server/item/inventory"
	lua "github.com/yuin/gopher-lua"
)

func GetInventory(i *inventory.Inventory, l *lua.LState) *lua.LTable {
	inv := Inventory{inventory: i}

	var exports map[string]lua.LGFunction = map[string]lua.LGFunction{
		"AddItem": inv.addItem,
	}

	t := l.SetFuncs(l.NewTable(), exports)

	return t
}

func GetArmour(i *inventory.Armour) *lua.LTable {

	return nil
}
