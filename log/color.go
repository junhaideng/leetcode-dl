package log

import "log"

// 简单的包装一下log
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

