package loader

import (
	"io/ioutil"
	"os"

	"github.com/AirplaneMC/airplane-plugins-api/controller"
	"github.com/AirplaneMC/airplane-plugins-api/controller/events"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func Load(log *logrus.Logger) error {
	log.Infoln("Loading plugins...")
	if _, err := os.Stat("./plugins"); err != nil {
		log.Debugln("Crating plugins folder...")
		if err := os.Mkdir("./plugins", 0666); err != nil {
			return err
		}
		return err
	}
	pluginsFolders, err := os.ReadDir("./plugins")
	if err != nil {
		return err
	}

	controller.GEvents = make(map[string]map[string][]*controller.CalledEvent)
	controller.Plugin_VM = lua.NewState()

	initLibs(controller.Plugin_VM)

	for _, obj := range pluginsFolders {
		if obj.IsDir() {
			c := controller.Plugin{}
			path := "./plugins/" + obj.Name()

			data, err := ioutil.ReadFile(path + "/config.yml")
			if err != nil {
				return err
			}

			err = c.ReadConfig(data)
			if err != nil {
				return err
			}

			err = controller.Init(&c, path, log)
			if err != nil {
				return err
			}
		}
	}

	events.CallOnLoadPE(log)

	return nil
}
