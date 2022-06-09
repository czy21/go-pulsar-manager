package bootstrap

import (
	"github.com/czyhome/go-pulsar-manager/controller"
	"github.com/czyhome/go-pulsar-manager/repository"
)

func Boot() {
	bootConfig()
	bootLog()
	repository.Boot()
	controller.Boot()
}
