package xconsole

import (
	"fmt"
	"github.com/coder2z/g-saber/xcolor"
)

var (
	debug = true
)

func ResetDebug(ok bool) {
	debug = ok
}

// Yellow ...
func Yellow(msg string) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Yellow(fmt.Sprintf("[xconsole]\t%v\n", msg)))
}

// Redf ...
func Yellowf(msg string, arg interface{}) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Yellowf(fmt.Sprintf("[xconsole]\t%-30v", msg), arg))
}

// Red ...
func Red(msg string) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Red(fmt.Sprintf("[xconsole]\t%v\n", msg)))
}

// Redf ...
func Redf(msg string, arg interface{}) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Redf(fmt.Sprintf("[xconsole]\t%-30v", msg), arg))
}

// Blue ...
func Blue(msg string) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Blue(fmt.Sprintf("[xconsole]\t%v\n", msg)))
}

// Greenf ...
func Bluef(msg string, arg interface{}) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Bluef(fmt.Sprintf("[xconsole]\t%-30v", msg), arg))
}

// Green ...
func Green(msg string) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Green(fmt.Sprintf("[xconsole]\t%v\n", msg)))
}

// Greenf ...
func Greenf(msg string, arg interface{}) {
	if !debug {
		return
	}
	fmt.Print(xcolor.Greenf(fmt.Sprintf("[xconsole]\t%-30v", msg), arg))
}
