package bootstrap

import (
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
)

func bootLog() {
	var LogWriter io.Writer = os.Stdout
	logFile := viper.GetString("log.file")
	if logFile != "" {
		f, _ := os.Create(logFile)
		LogWriter = io.MultiWriter(f)
	}
	logger := log.Default()
	logger.SetPrefix("[SYS] ")
	logger.SetOutput(LogWriter)
}
