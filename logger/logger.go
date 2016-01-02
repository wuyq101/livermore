package logger

import (
	"fmt"
	"log"
	"os"
)

var lg = log.New(os.Stdout, "", log.LstdFlags)

const (
	FATAL = 0
	PANIC = 1
	ERROR = 2
	WARN  = 3
	INFO  = 4
	DEBUG = 5
)

var Level int64 = INFO

func Info(format string, v ...interface{}) {
	if Level >= INFO {
		lg.Output(2, fmt.Sprintf("[Info] "+format+"\n", v...))
	}
}

func Error(format string, v ...interface{}) {
	if Level >= ERROR {
		lg.Output(2, fmt.Sprintf("[Error] "+format+"\n", v...))
	}
}
