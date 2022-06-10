package bootstrap

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/spf13/viper"
	"os"
)

func bootConfig() {
	appConf := os.Getenv("GPM_CONFIG_FILE")
	if appConf == "" {
		viper.SetConfigFile("app.yaml")
	} else {
		viper.SetConfigFile(appConf)
	}
	err := viper.ReadInConfig()
	exception.Check(err)
	viper.WatchConfig()
}
