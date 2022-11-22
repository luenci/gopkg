package log

import (
	"fmt"
	"log"
	"runtime/debug"
)

const (
	info    = "[INFO]"
	trac    = "[TRAC]"
	erro    = "[ERROR]"
	warn    = "[WARN]"
	success = "[SUCCESS]"
)

// Trace see complete color rules
// document in https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func Trace(format string, a interface{}) {
	prefix := yellow(trac)
	log.Println(formatLog(prefix), fmt.Sprintf(format, a))
}

func Info(format string, a interface{}) {
	prefix := blue(info)
	log.Println(formatLog(prefix), fmt.Sprintf(format, a))
}

func Success(format string, a interface{}) {
	prefix := green(success)
	log.Println(formatLog(prefix), fmt.Sprintf(format, a))
}

func Warning(format string, a interface{}) {
	prefix := magenta(warn)
	log.Println(formatLog(prefix), fmt.Sprintf(format, a))
}

func Errors(format string, a interface{}) {
	prefix := red(erro)
	log.Println(formatLog(prefix), fmt.Sprintf(format, a))
}

func Panicf(format string, a interface{}) {
	buf := debug.Stack()
	log.Panicf(format, map[string]interface{}{
		"stack": string(buf),
		"err":   a,
	})
}
