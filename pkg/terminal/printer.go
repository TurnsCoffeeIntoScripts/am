package terminal

import (
	"fmt"
	"github.com/TwiN/go-color"
)

func InfoMessage(format string, arg ...any) {
	msg := fmt.Sprintf(format, arg...)
	println(color.InGreen(msg))
}

func ErrorMessage(format string, arg ...any) {
	msg := fmt.Sprintf(format, arg...)
	println(color.InRed(msg))
}
