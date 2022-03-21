package main

import (
	"fmt"

	"github.com/quortex/kapitan-docs/pkg/flags"
	log "github.com/sirupsen/logrus"
)

var opts flags.Options

func main() {
	// Parse flags
	if err := flags.Parse(&opts); err != nil {
		log.Fatal(fmt.Errorf("Invalid flags: %w", err))
	}
	log.SetLevel(opts.LogLevel.Level)
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
}
