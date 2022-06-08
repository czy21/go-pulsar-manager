package bootstrap

import (
	"github.com/czyhome/go-pulsar-manager/exception"
	"github.com/spf13/viper"
)

func bootConfig() {
	viper.SetConfigFile("./app.yaml")
	err := viper.ReadInConfig()
	exception.Check(err)
	viper.WatchConfig()
}
