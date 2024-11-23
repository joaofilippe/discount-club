package logger

import (
	"log/slog"
	"os"
	"runtime"
)

type Logger struct {
	Logger  *slog.Logger
}
// NewLogger returns a new logger
func NewLogger() *Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Logger{
		Logger:  logger,
	}
}

func (l *Logger) Errorf(err error) {
	_, _, line, _ := runtime.Caller(1)

	l.Logger.Error(err.Error(), "line=", line)
}

func (l *Logger) Fatalf(err error) {
	_, _, line, _ := runtime.Caller(1)

	l.Logger.Error(err.Error(), "line=", line)
	panic(err)
}

func (l *Logger) Infof(message string) {
	_, file, line, _ := runtime.Caller(1)

	l.Logger.Info(message, "line", line, "file", file)
}