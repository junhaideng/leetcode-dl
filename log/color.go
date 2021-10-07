// Package log 简单的包装一下标准 log
package log

import (
	"io"
	"log"
	"os"
)

func SetOutput(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println("Create log file failed, use console as output")
		return
	}
	log.SetOutput(io.MultiWriter(f, os.Stdout))
}

func Infof(format string, v ...interface{}) {
	log.SetPrefix("\033[0;36m")
	log.Printf(format, v...)
	log.SetPrefix("\033[00m")
}

func Warn(v ...interface{}) {
	log.SetPrefix("\033[1;33m")
	log.Println(v...)
	log.SetPrefix("\033[00m")
}

func Errorf(format string, v ...interface{}) {
	log.SetPrefix("\033[0;31m")
	log.Printf(format, v...)
	log.SetPrefix("\033[00m")
}

func Successf(format string, v ...interface{}) {
	log.SetPrefix("\033[0;32m")
	log.Printf(format, v...)
	log.SetPrefix("\033[00m")
}
