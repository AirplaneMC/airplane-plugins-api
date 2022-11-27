package controller

import (
	"fmt"

	"github.com/AirplaneMC/airplane-plugins-api/controller/events"
	"github.com/AirplaneMC/airplane-plugins-api/controller/types"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

var (
	plugins []*types.Plugin
)

func Init(log *logrus.Logger, l *lua.LState, p *types.Plugin, path string) error {
	if err := l.DoFile(path + "/src/index.lua"); err != nil {
		return err
	}

	mainV := l.GetGlobal(p.Main)
	if mainV.Type() != lua.LTTable {
		return fmt.Errorf("main(%v) value in plugin config %v is wrong", p.Main, p.Name)
	}
	main, ok := mainV.(*lua.LTable)
	if !ok {
		return fmt.Errorf("%v values cannot be used as a table", p.Main)
	}

	events.InitEvent(log, p, main)

	plugins = append(plugins, p)

	return nil
}
