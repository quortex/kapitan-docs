package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/quortex/kapitan-docs/pkg/document"
	"github.com/quortex/kapitan-docs/pkg/flags"
	"github.com/quortex/kapitan-docs/pkg/kapitan"
	log "github.com/sirupsen/logrus"
)

var opts flags.Options

func main() {
	// Flags parsing
	if err := flags.Parse(&opts); err != nil {
		log.Error(fmt.Errorf("Invalid flags: %w", err))
		os.Exit(1)
	}
	log.SetLevel(opts.LogLevel.Level)

	// Project directory parsing
	log.Debugf("Parsing project directory: %s", opts.Positional.Directory)
	p, err := kapitan.NewProject(opts.Positional.Directory)
	if err != nil {
		log.Error(fmt.Errorf("Cannot parse project: %w", err))
		os.Exit(1)
	}

	// Read given template file
	f, err := ioutil.ReadFile(opts.TemplateFile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Error(fmt.Errorf("Read file error: %w", err))
			os.Exit(1)
		}
	}

	// Render markdown documentation
	rendered, err := document.RenderAsMarkdown(p, string(f))
	if err != nil {
		log.Error(fmt.Errorf("Cannot render documentation: %w", err))
		os.Exit(1)
	}

	fmt.Println(rendered)
}
