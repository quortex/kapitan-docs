package flags

import (
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

// LogLevel is a wrapper for logger log level.
type LogLevel struct {
	logrus.Level
}

// UnmarshalFlag is the go-flags Value UnmarshalFlag implementation for LogLevel.
func (l *LogLevel) UnmarshalFlag(arg string) error {
	level, err := logrus.ParseLevel(arg)
	if err != nil {
		return err
	}
	l.Level = level
	return nil
}

// MarshalFlag is the go-flags Value MarshalFlag implementation for LogLevel.
func (l *LogLevel) MarshalFlag() (string, error) {
	return l.Level.String(), nil
}

// Options wraps all flags.
type Options struct {
	DryRun       bool     `long:"dry-run" short:"d" description:"Don't render any markdown file, just print in the console."`
	LogLevel     LogLevel `long:"log-level" short:"l" description:"Level of logs that should printed, one of (panic, fatal, error, warning, info, debug, trace)." default:"error"`
	TemplateFile string   `long:"template-file" short:"t" description:"gotemplate file path from which documentation will be generated." default:"README.md.gotmpl"`
	Positional   struct {
		Directory string `description:"Kapitan project directory." default:"."`
	} `positional-args:"yes"`
}

// Parse parses flags into give Option.
func Parse(opts *Options) error {
	parser := flags.NewParser(opts, flags.Default)
	_, err := parser.Parse()
	return err
}
