package logging

import (
	"github.com/op/go-logging"
	"log"
	"os"
	"telegramxbot/internal/config"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func Configure(cfg *config.Config) {
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	formatter := logging.NewBackendFormatter(backend, format)

	level, err := logging.LogLevel(cfg.Log.Level)
	if err != nil {
		log.Panicln(err)
	}
	backendLeveled := logging.AddModuleLevel(formatter)
	backendLeveled.SetLevel(level, "")

	logging.SetBackend(backendLeveled)

}
