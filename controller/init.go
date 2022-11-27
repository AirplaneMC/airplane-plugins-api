package controller

import (
	"fmt"

	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var (
	Plugin_VM *lua.LState

	plugins []*Plugin
	GEvents map[string]map[string][]*CalledEvent
)

func Init(p *Plugin, path string, log *logrus.Logger) error {
	if err := Plugin_VM.DoFile(path + "/src/index.lua"); err != nil {
		return err
	}

	mainV := Plugin_VM.GetGlobal(p.Main)
	if mainV.Type() != lua.LTTable {
		return fmt.Errorf("main(%v) value in plugin config %v is wrong", p.Main, p.Name)
	}
	main, ok := mainV.(*lua.LTable)
	if !ok {
		return fmt.Errorf("%v values cannot be used as a table", p.Main)
	}

	registerEvents(main, p, log)

	plugins = append(plugins, p)

	return nil
}

func registerEvents(main *lua.LTable, p *Plugin, log *logrus.Logger) {
	eventsV := main.RawGetString("events")
	if eventsV.Type() != lua.LTTable {
		return
	}

	eventsL := eventsV.(*lua.LTable)

	eventsL.ForEach(func(key, value lua.LValue) {
		switch key.String() {
		case PLUGIN_EVENTS:
			registerCalledFunction(log, p, value, PLUGIN_EVENTS, "LOAD", "Load")
			registerCalledFunction(log, p, value, PLUGIN_EVENTS, "UNLOAD", "Unload")
		case PLAYER_EVENTS:
			registerCalledFunction(log, p, value, PLAYER_EVENTS, "JOIN", "Join")
			registerCalledFunction(log, p, value, PLAYER_EVENTS, "QUIT", "Quit")
		}
	})
}

func registerCalledFunction(log *logrus.Logger, p *Plugin, main lua.LValue, event, eventName, calledFunction string) {
	if main.Type() != lua.LTTable && main.Type() != lua.LTNil {
		log.Warningf("[%v] pluginEvents is not a table.\n", p.Name)
		return
	}

	eventsF, ok := main.(*lua.LTable)
	if !ok {
		log.Warningf("[%v] %v values cannot be used as a table", p.Name, main.String())
		return
	}

	calledF := eventsF.RawGetString(calledFunction)
	if calledF.Type() != lua.LTFunction && calledF.Type() != lua.LTNil {
		log.Warning("[%v] %v values cannot be used as a function.", p.Name, main.String())
		return
	}
	if calledF.Type() == lua.LTNil {
		return
	}

	if GEvents[event] == nil {
		GEvents[event] = make(map[string][]*CalledEvent)
		GEvents[event][eventName] = make([]*CalledEvent, 0)
	}

	functions := GEvents[PLUGIN_EVENTS][event]
	functions = append(functions, &CalledEvent{
		Plugin:   p,
		Function: calledF,
	})

	GEvents[event][eventName] = functions
}
