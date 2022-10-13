package main

import (
	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)

	//getwd, err := os.Getwd()
	//if err != nil {
	//	return
	//}
	//fmt.Println(getwd)
	//return

	path := "/Users/gongyao/workspace/goproject/src/mygo2/chapter9log/var/vmlet.log"
	logger:=&lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,
		MaxSize:    20, // megabytes
		MaxBackups: 5,
		MaxAge:     30,    //days
		Compress:   false, // disabled by default
	}
	writers := []io.Writer{
		logger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)
}

func main() {
	//getwd, err := os.Getwd()
	//if err != nil {
	//	return
	//}
	//fmt.Println(getwd)
	//return

	for {
		log.Error("xixixixi")
		log.Info("hello, world!")
	}
}