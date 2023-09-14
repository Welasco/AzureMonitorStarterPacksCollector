package logger

import (
	"log"
	"os"
)

// var logFile *os.File
//
//	type Logger struct {
//		log.Logger
//	}
//
// var l = log.New(os.Stdout, "", 0)
var l log.Logger

func Init(file string) {
	// open log file
	fileName := "collector.log"
	var err error
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}

	//logger = log.New(logFile, "INFO: ", log.Lshortfile|log.LstdFlags)
	l.SetOutput(logFile)
	l.SetFlags(log.Lshortfile | log.LstdFlags)

	l.Println("INIT Starting collector...")

	//defer logFile.Close()
}

func Debug(msg ...interface{}) {
	l.Println("DEBUG:", msg)
}

func Info(msg ...interface{}) {
	l.Println("INFO:", msg)
}

func Warning(msg ...interface{}) {
	l.Println("WARNING:", msg)
}

func Error(msg ...interface{}) {
	l.Println("ERROR:", msg)
}
