package logger

import (
	"github.com/go-kit/log"
	"os"
)

func InitLogger(port string) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)
	return logger
}
