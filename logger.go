package logger

import (
	"fmt"
	"log"
	"os"
)

var LogLevelMap = map[string]int{
	"debug":   0,
	"info":    1,
	"warn":    2,
	"warning": 2,
	"error":   3,
	"fatal":   4,
}

var (
	LogLevel    int
	DebugLogger = log.New(os.Stdout, "[D] ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger  = log.New(os.Stdout, "[I] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger  = log.New(os.Stdout, "[W] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[E] ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(os.Stdout, "[F] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func InitLogger(level, defaultLevel string) {
	val, ok := LogLevelMap[level]
	if !ok {
		LogLevel = LogLevelMap[defaultLevel]
	} else {
		LogLevel = val
	}
}

func Debug(format string, v ...interface{}) {
	if 0 >= LogLevel {
		DebugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Info(format string, v ...interface{}) {
	if 1 >= LogLevel {
		InfoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Warn(format string, v ...interface{}) {
	if 2 >= LogLevel {
		WarnLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Error(format string, v ...interface{}) {
	if 3 >= LogLevel {
		ErrorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Fatal(format string, v ...interface{}) {
	if 4 >= LogLevel {
		FatalLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Debugln(v ...interface{}) {
	if 0 >= LogLevel {
		DebugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Infoln(v ...interface{}) {
	if 1 >= LogLevel {
		InfoLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Warnln(v ...interface{}) {
	if 2 >= LogLevel {
		WarnLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Errorln(v ...interface{}) {
	if 3 >= LogLevel {
		ErrorLogger.Output(2, fmt.Sprintln(v...))
	}
}

func Fatalln(v ...interface{}) {
	if 4 >= LogLevel {
		FatalLogger.Output(2, fmt.Sprintln(v...))
	}
}
