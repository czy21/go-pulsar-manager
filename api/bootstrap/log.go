package bootstrap

import (
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
)

func bootLog() {
	var logWriter io.Writer = os.Stdout
	logFile := viper.GetString("log.file")
	if logFile != "" {
		f, _ := os.Create(logFile)
		logWriter = io.MultiWriter(f)
	}
	logger := log.Default()
	logger.SetPrefix("[SYS] ")
	logger.SetOutput(logWriter)
}
