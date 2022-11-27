package player

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller"
	"github.com/df-mc/dragonfly/server/player"
)

type PlayerHandler struct {
	player.NopHandler

	CurrentH controller.Handler
}
