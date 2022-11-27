package controller

import (
	lua "github.com/yuin/gopher-lua"
)

type HType string

type Handler struct {
	Type HType
	Main lua.LValue
}

const (
	WorldHandler     HType = "world"
	InventoryHandler HType = "inventory"
	PlayerHandler    HType = "player"
)
