package log

import (
	"errors"
	"fmt"
	"lib/base"
	"game/config"
	"log"
	"os"
	"path"
	"strings"
	"syscall"
	"time"
)

// levels
const (
	debugLevel   = 0
	releaseLevel = 1
	errorLevel   = 2
	fatalLevel   = 3
)

const (
	printDebugLevel   = "[debug  ] "
	printReleaseLevel = "[release] "
	printErrorLevel   = "[error  ] "
	printFatalLevel   = "[fatal  ] "
)

type Logger struct {
	level      int
	baseLogger *log.Logger
	baseFile   *os.File
}

func New(strLevel string, pathname string, flag int) (*Logger, error) {
	// level
	var level int
	switch strings.ToLower(strLevel) {
	case "debug":
		level = debugLevel
	case "release":
		level = releaseLevel
	case "error":
		level = errorLevel
	case "fatal":
		level = fatalLevel
	default:
		return nil, errors.New("unknown level: " + strLevel)
	}

	// logger
	var baseLogger *log.Logger
	var baseFile *os.File
	if pathname != "" {
		filename := get_log_name()
		base.EnsureDir(pathname)
		filepath := path.Join(pathname, filename)
		file, err := os.OpenFile(filepath, syscall.O_APPEND|syscall.O_CREAT, 0666)
		if err != nil {
			return nil, err
		}

		baseLogger = log.New(file, "", flag)
		baseFile = file
	} else {
		baseLogger = log.New(os.Stdout, "", flag)
	}

	// new
	logger := new(Logger)
	logger.level = level
	logger.baseLogger = baseLogger
	logger.baseFile = baseFile

	return logger, nil
}

func get_log_name() string {
	now := time.Now()
	return fmt.Sprintf("%s_%s_%d_%d%02d%02d.log",
		config.GetGameCode(),
		config.GetAgentCode(),
		config.GetServerID(),
		now.Year(),
		now.Month(),
		now.Day())
}

// It's dangerous to call the method on logging
func (logger *Logger) Close() {
	if logger.baseFile != nil {
		logger.baseFile.Close()
	}

	logger.baseLogger = nil
	logger.baseFile = nil
}

func (logger *Logger) doPrintf(level int, printLevel string, format string, a ...interface{}) {
	if level < logger.level {
		return
	}
	if logger.baseLogger == nil {
		panic("logger closed")
	}

	format = printLevel + format
	logger.baseLogger.Output(3, fmt.Sprintf(format, a...))

	if level == fatalLevel {
		os.Exit(1)
	}
}

func (logger *Logger) Debug(format string, a ...interface{}) {
	logger.doPrintf(debugLevel, printDebugLevel, format, a...)
}

func (logger *Logger) Release(format string, a ...interface{}) {
	logger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
}

func (logger *Logger) Error(format string, a ...interface{}) {
	logger.doPrintf(errorLevel, printErrorLevel, format, a...)
}

func (logger *Logger) Fatal(format string, a ...interface{}) {
	logger.doPrintf(fatalLevel, printFatalLevel, format, a...)
}

var gLogger, _ = New(config.GetAgentCode(), "./log", log.LstdFlags|log.Lshortfile|log.Llongfile)

// It's dangerous to call the method on logging
func Export(logger *Logger) {
	if logger != nil {
		gLogger = logger
	}
}

func Debug(format string, a ...interface{}) {
	fmt.Println(format, a)
	gLogger.doPrintf(debugLevel, printDebugLevel, format+base.GetLineBreak(), a...)
}

func Release(format string, a ...interface{}) {
	fmt.Println(format, a)
	gLogger.doPrintf(releaseLevel, printReleaseLevel, format+base.GetLineBreak(), a...)
}

func Error(format string, a ...interface{}) {
	fmt.Println(format, a)
	gLogger.doPrintf(errorLevel, printErrorLevel, format+base.GetLineBreak(), a...)
}

func Fatal(format string, a ...interface{}) {
	fmt.Println(format, a)
	gLogger.doPrintf(fatalLevel, printFatalLevel, format+base.GetLineBreak(), a...)
}

func Close() {
	gLogger.Close()
}
