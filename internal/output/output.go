package output

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func Error(line string) {
	fmt.Println(aurora.Red(line))
}

func Errorf(line string, args ...interface{}) {
	Error(fmt.Sprintf(line, args...))
}

func Info(line string) {
	fmt.Println(aurora.Green(line))
}

func Infof(line string, args ...interface{}) {
	Info(fmt.Sprintf(line, args...))
}

func Warn(line string) {
	fmt.Println(aurora.Brown(line))
}

func Warnf(line string, args ...interface{}) {
	Warn(fmt.Sprintf(line, args...))
}

func NoFormat(line string) {
	fmt.Println(line)
}

func NoFormatf(line string, args ...interface{}) {
	NoFormat(fmt.Sprintf(line, args...))
}
