package item

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

var items map[string]world.Item = map[string]world.Item{
	"amethystShard": item.AmethystShard{},
	"apple":         item.Apple{},
	"arrow":         item.Arrow{},

	"woodAxe":      item.Axe{Tier: item.ToolTierWood},
	"stoneAxe":     item.Axe{Tier: item.ToolTierStone},
	"ironAxe":      item.Axe{Tier: item.ToolTierIron},
	"goldAxe":      item.Axe{Tier: item.ToolTierGold},
	"diamondAxe":   item.Axe{Tier: item.ToolTierDiamond},
	"netheriteAxe": item.Axe{Tier: item.ToolTierNetherite},
}

func GetStack(itemName string, n int) (item.Stack, error) {
	v, ok := items[itemName]
	if !ok {
		return item.Stack{}, fmt.Errorf("item %v does not exist", itemName)
	}

	return item.NewStack(v, n), nil
}
