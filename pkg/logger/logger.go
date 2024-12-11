package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type ILogger interface {
	Debug(msg string, params ...interface{})
	Info(msg string, params ...interface{})
	Success(msg string, params ...interface{})
	Warn(msg string, params ...interface{})
	Error(msg string, params ...interface{})
	Fatal(msg string, params ...interface{})
}

type logger struct {
	colorReset  string
	colorCyan   string
	colorRed    string
	colorYellow string
	colorGreen  string
}

func NewLogger() ILogger {
	return &logger{
		colorReset:  "\033[0m",
		colorCyan:   "\033[36m",
		colorRed:    "\033[31m",
		colorYellow: "\033[33m",
		colorGreen:  "\033[32m",
	}
}

func (l *logger) log(level, color, msg string, params ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	filename := filepath.Base(file)

	// Reemplaza el nombre del archivo por el nombre deseado
	var origin string
	if filename == "db.go" {
		origin = fmt.Sprintf("[config.ConnectDB:%d]", line)
	} else {
		// Ajusta según sea necesario para otros archivos
		origin = fmt.Sprintf("[%s:%d]", strings.TrimSuffix(filename, filepath.Ext(filename)), line)
	}

	logMsg := fmt.Sprintf("[%s] %s [%s]", level, origin, fmt.Sprintf(msg, params...))
	log.Printf("%s%s%s", color, logMsg, l.colorReset)
}

func (l *logger) Debug(msg string, params ...interface{}) {
	l.log("DEBU", l.colorCyan, msg, params...)
}

func (l *logger) Info(msg string, params ...interface{}) {
	l.log("INFO", l.colorCyan, msg, params...)
}

func (l *logger) Success(msg string, params ...interface{}) {
	l.log("SUCC", l.colorGreen, msg, params...)
}

func (l *logger) Warn(msg string, params ...interface{}) {
	l.log("WARN", l.colorYellow, msg, params...)
}

func (l *logger) Error(msg string, params ...interface{}) {
	l.log("ERRO", l.colorRed, msg, params...)
}

func (l *logger) Fatal(msg string, params ...interface{}) {
	l.log("FATA", l.colorRed, msg, params...)
	log.Panicf("[%s]", fmt.Sprintf(msg, params...))
}
