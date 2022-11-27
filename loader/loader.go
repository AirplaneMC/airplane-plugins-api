package loader

import (
	"io/ioutil"
	"os"

	"github.com/AirplaneMC/airplane-plugins-api/controller"
	"github.com/AirplaneMC/airplane-plugins-api/controller/events"
	"github.com/AirplaneMC/airplane-plugins-api/controller/types"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func Load(log *logrus.Logger) (*lua.LState, error) {
	log.Infoln("Loading plugins...")
	if _, err := os.Stat("./plugins"); err != nil {
		log.Debugln("Crating plugins folder...")
		if err := os.Mkdir("./plugins", 0666); err != nil {
			return nil, err
		}
		return nil, err
	}
	pluginsFolders, err := os.ReadDir("./plugins")
	if err != nil {
		return nil, err
	}

	l := lua.NewState()
	initLibs(l)

	for _, obj := range pluginsFolders {
		if obj.IsDir() {
			c := types.Plugin{}
			path := "./plugins/" + obj.Name()

			data, err := ioutil.ReadFile(path + "/config.yml")
			if err != nil {
				return nil, err
			}

			err = c.ReadConfig(data)
			if err != nil {
				return nil, err
			}

			err = controller.Init(log, l, &c, path)
			if err != nil {
				return nil, err
			}
		}
	}

	events.CallOnLoadPE(log, l)

	return l, nil
}
