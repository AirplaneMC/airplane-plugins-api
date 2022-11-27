package effect

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/entity/effect"
)

var effects map[string]effect.LastingType = map[string]effect.LastingType{
	"absorption":     effect.Absorption{},
	"blindness":      effect.Blindness{},
	"conduitPower":   effect.ConduitPower{},
	"darkness":       effect.Darkness{},
	"fatalPoison":    effect.FatalPoison{},
	"fireResistance": effect.FireResistance{},
	"haste":          effect.Haste{},
	"healthBoost":    effect.HealthBoost{},
	"hunger":         effect.Hunger{},
	"invisibility":   effect.Invisibility{},
	"jumpBoost":      effect.JumpBoost{},
	"levitation":     effect.Levitation{},
	"miningFatigue":  effect.MiningFatigue{},
	"nausea":         effect.Nausea{},
	"nightVision":    effect.NightVision{},
	"poison":         effect.Poison{},
	"regeneration":   effect.Regeneration{},
	"resistance":     effect.Resistance{},
	"saturation":     effect.Saturation{},
	"slowFalling":    effect.SlowFalling{},
	"slowness":       effect.Slowness{},
	"speed":          effect.Speed{},
	"strength":       effect.Strength{},
	"waterBreathing": effect.WaterBreathing{},
	"eeakness":       effect.Weakness{},
	"wither":         effect.Wither{},
}

func GetEffect(effect string) (effect.LastingType, error) {
	v, ok := effects[effect]
	if !ok {
		return nil, fmt.Errorf("effect %v does not exist", effect)
	}

	return v, nil
}
