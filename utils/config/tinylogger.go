package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warn
	Error
)

type LogEntry struct {
	tag   string
	level LogLevel
	msg   string
	err   error
}

type Logger struct {
	entry LogEntry
}

func (l *Logger) SetTag(tag string) {
	l.entry.tag = tag
}

func (l *Logger) SetLevel(level LogLevel) {
	l.entry.level = level
}

func (l *Logger) SetMsg(msg string) {
	l.entry.msg = msg
}

func (l *Logger) SetError(err error) {
	l.entry.err = err
}

func (l *Logger) Log() {
	var logLevel string
	switch l.entry.level {
	case Warn:
		logLevel = "W"
	case Error:
		logLevel = "E"
	default:
		logLevel = "I"
	}

	logTime := time.Now().Format("15:04:05")
	if l.entry.err == nil {
		fmt.Printf("[%s] [%s] [%s] %s\n", logTime, logLevel, l.entry.tag, l.entry.msg)
	} else {
		fmt.Printf("[%s] [%s] [%s] %s: %s\n", logTime, logLevel, l.entry.tag, l.entry.msg, l.entry.err)
	}

}

func (l *Logger) Info(format string, args ...interface{}) {
	var msg string
	if format == "" {
		l.Warn("A format from Info is empty")
	} else {
		msg = fmt.Sprintf(format, args...)
	}
	l.SetLevel(Info)
	l.SetMsg(msg)
	l.Log()
}

func (l *Logger) Warn(format string, args ...interface{}) {
	var msg string
	if format == "" {
		l.Warn("A format from Warn is empty")
	} else {
		msg = fmt.Sprintf(format, args...)
	}
	l.SetLevel(Warn)
	l.SetMsg(msg)
	l.Log()
}

func (l *Logger) Error(format string, args ...interface{}) error {
	var msg string
	var err error
	if len(args) > 0 {
		if e, ok := args[len(args)-1].(error); ok {
			err = e
			args = args[:len(args)-1]
		}
	}
	msg = fmt.Sprintf(format, args...)

	l.SetLevel(Error)
	l.SetMsg(msg)
	l.SetError(err)
	l.Log()
	return err // 可以尝试去除返回？
}

func NewTinyLogger(tag string) *Logger {
	outputFolder := "log"
	logFileName := time.Now().Format("20060102") + ".txt"
	logFilePath := filepath.Join(logFileName, outputFolder)
	logFile, _ := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // 备注：记得处理错误
	defer logFile.Close()
	return &Logger{entry: LogEntry{tag: tag}}
}
