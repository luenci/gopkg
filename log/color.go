package log

import (
	"fmt"
	"time"
)

const (
	Red = uint8(iota + 91)
	Green
	Yellow
	Blue
	Magenta // 洋红

)

func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", Red, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", Green, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", Yellow, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", Blue, s)
}

func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", Magenta, s)
}

func formatLog(prefix string) string {
	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
}
