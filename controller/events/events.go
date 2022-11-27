package events

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/events/plugin"
	"github.com/AirplaneMC/airplane-plugins-api/controller/types"
	"github.com/sirupsen/logrus"

	playerEvents "github.com/AirplaneMC/airplane-plugins-api/controller/events/player"
	lua "github.com/yuin/gopher-lua"
)

type CalledEvent struct {
	Plugin   *types.Plugin
	Function lua.LValue
	Self     lua.LValue
}

var events map[string]map[string][]*CalledEvent = make(map[string]map[string][]*CalledEvent)

func InitEvent(log *logrus.Logger, p *types.Plugin, main *lua.LTable) {
	eventsList := main.RawGetString("events")
	if eventsList.Type() != lua.LTTable {
		return
	}

	eventsL := eventsList.(*lua.LTable)

	eventsL.ForEach(func(key, value lua.LValue) {
		switch key.String() {
		case PLUGIN_EVENTS:
			for k, v := range plugin.PluginGE {
				registerEvent(log, value, p, PLUGIN_EVENTS, k, v)
			}
		case PLAYER_EVENTS:
			for k, v := range playerEvents.PlayerGE {
				registerEvent(log, value, p, PLAYER_EVENTS, k, v)
			}
		}
	})
}

func registerEvent(log *logrus.Logger, eventTable lua.LValue, p *types.Plugin, eventType, eventName, eventFunction string) {
	if eventTable.Type() != lua.LTTable && eventTable.Type() != lua.LTNil {
		log.Warningf("[%v] %v is not a table.\n", eventType, p.Name)
		return
	}

	eventsF, ok := eventTable.(*lua.LTable)
	if !ok {
		log.Warningf("[%v] %v values cannot be used as a table", p.Name, eventTable.String())
		return
	}

	calledF := eventsF.RawGetString(eventFunction)
	if calledF.Type() != lua.LTFunction && calledF.Type() != lua.LTNil {
		log.Warning("[%v] %v values cannot be used as a function.", p.Name, eventTable.String())
		return
	}
	if calledF.Type() == lua.LTNil {
		return
	}

	if events[eventType] == nil {
		events[eventType] = make(map[string][]*CalledEvent)
		events[eventType][eventName] = make([]*CalledEvent, 0)
	}

	functions := events[PLUGIN_EVENTS][eventType]
	functions = append(functions, &CalledEvent{
		Plugin:   p,
		Function: calledF,
	})

	events[eventType][eventName] = functions
}
