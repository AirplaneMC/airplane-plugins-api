package player

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/types"
	"github.com/df-mc/dragonfly/server/player"
)

type PlayerHandler struct {
	player.NopHandler

	Handler *types.Handler
}
