package logger

import (
	"ASL/utils/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Debug
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
	entry         LogEntry
	isDebug       bool
	isFileLogging bool
	logFile       *os.File
}

var defaultLogger = &Logger{entry: LogEntry{tag: "Logger"}}

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

func (l *Logger) SetDebug(isDebug bool) {
	l.isDebug = isDebug
}

func (l *Logger) SetFileLogging(isFileLogging bool) {
	l.isFileLogging = isFileLogging
}

func (l *Logger) Log() {
	var logLevel string
	switch l.entry.level {
	case Debug:
		logLevel = "D"
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

	if l.isFileLogging {
		log.SetOutput(l.logFile)
		if l.entry.err == nil {
			log.Printf("[%s] [%s] [%s] %s\n", logTime, logLevel, l.entry.tag, l.entry.msg)
		} else {
			log.Printf("[%s] [%s] [%s] %s: %s\n", logTime, logLevel, l.entry.tag, l.entry.msg, l.entry.err)
		}
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

func (l *Logger) Debug(format string, args ...interface{}) {
	if !l.isDebug {
		return
	}

	var msg string
	if format == "" {
		l.Warn("A format from Debug is empty")
	} else {
		msg = fmt.Sprintf(format, args...)
	}
	l.SetLevel(Debug)
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

func (l *Logger) Close() {
	if l.logFile != nil {
		l.logFile.Close()
	}
}

func NewLogger(tag string) *Logger {
	if config.IS_FILELOGGING {
		logFileName := time.Now().Format("20060102") + ".txt"
		logFilePath := filepath.Join(filepath.Dir(config.EXEC_PATH), "log", logFileName)
		os.Mkdir(filepath.Join(filepath.Dir(config.EXEC_PATH), "log"), 0755)
		logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			defaultLogger.Error("Error opening log file %s: %v", logFilePath, err)
			return nil
		}
		return &Logger{entry: LogEntry{tag: tag}, logFile: logFile, isDebug: config.IS_DEBUG, isFileLogging: config.IS_FILELOGGING}
	} else {
		return &Logger{entry: LogEntry{tag: tag}, isDebug: config.IS_DEBUG, isFileLogging: config.IS_FILELOGGING}
	}
}
