package logger

import (
	"log"
)
// Create logger struct which includes App (it will be static for this app (go crud backend), we are gonna change it after), Module (like UserApi or EntryService, you will create that logger instances for each module), And date format, it will be static for now too. 
// After that when you tried to create logi log format should be. Log time | module name(like UserApi) - Your log message
var LogLevel uint8

func Fatal(v ...interface{}) {

	if LogLevel > 0 {
		log.Fatal(v...)
	}

}

func Info(v ...interface{}) {

	if LogLevel > 1 {
		log.Println(v...)
	}

}

func Error(v ...interface{}) {

	if LogLevel > 2 {
		log.Println(v...)
	}

}
