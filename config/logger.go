package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	fatal   *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, ">DEBUG: ", logger.Flags()),
		info:    log.New(writer, ">INFO: ", logger.Flags()),
		warning: log.New(writer, ">WARNING: ", logger.Flags()),
		err:     log.New(writer, ">ERROR: ", logger.Flags()),
		fatal:   log.New(writer, ">FATAL: ", logger.Flags()),
		writer:  writer,
	}
}

// create non-formatted logs
func (l *Logger) Debug(msg string) {
	l.debug.Println(msg)
}

func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

func (l *Logger) Warning(msg string) {
	l.warning.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.err.Println(msg)
}

func (l *Logger) Fatal(msg string) {
	l.fatal.Println(msg)
}

// create format enabled logs
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.debug.Printf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.info.Printf(format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.warning.Printf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.err.Printf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.fatal.Printf(format, args...)
}
