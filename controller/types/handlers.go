package types

import (
	lua "github.com/yuin/gopher-lua"
)

type HType string

type Handler struct {
	Type HType
	Main lua.LValue
}

/*
type PlayerHandler struct {
	player.NopHandler

	CurrentHandler *Handler
}

type WorldHandler struct {
	world.NopHandler

	CurrentHandler *Handler
}

type InventoryHandler struct {
	inventory.NopHandler

	CurrentHandler *Handler
}*/

const (
	PlayerHandlerT    HType = "world"
	WorldHandlerT     HType = "inventory"
	InventoryHandlerT HType = "player"
)
