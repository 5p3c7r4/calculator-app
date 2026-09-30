package logger

import (
	"io"
	"log/slog"
	"os"
)

var handler slog.Handler

type logger struct {
	logLevelVar *slog.LevelVar
	out         io.Writer
	options     *slog.HandlerOptions
}

func (l *logger) GetJSONLogger() *slog.Logger {
	handler = slog.NewJSONHandler(l.out, l.options)
	logger := slog.New(handler)
	return logger
}

func New(envFile ...string) *logger {
	var levelVar = new(slog.LevelVar)
	logger := &logger{
		logLevelVar: levelVar,
		out:         os.Stderr,
		options: &slog.HandlerOptions{
			Level:     levelVar,
			AddSource: true,
		},
	}
	return logger
}
