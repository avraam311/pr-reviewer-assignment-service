// Package logger provides logger
package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Zerolog is a logger type exported by the logger package.
type Zerolog = zerolog.Logger

// Logger is a global logger instance.
var Logger Zerolog

// Init initializes the global logger with output to the console and file.
func Init() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	fileWriter, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open log file")
	}
	multi := io.MultiWriter(fileWriter, consoleWriter)

	Logger = zerolog.New(multi).With().Timestamp().Logger()
}
