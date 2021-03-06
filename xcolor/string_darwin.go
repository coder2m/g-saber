// +build darwin

package xcolor

import (
	"fmt"
)

// Yellow ...
func Yellow(msg string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", msg)
}

// Yellowf ...
func Yellowf(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m %+v\n", msg, arg)
}

// Red ...
func Red(msg string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", msg)
}

// Redf ...
func Redf(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m %+v\n", msg, arg)
}

// Blue ...
func Blue(msg string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", msg)
}

// Bluef ...
func Bluef(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m %+v\n", msg, arg)
}

// Green ...
func Green(msg string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", msg)
}

// Greenf ...
func Greenf(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m %+v\n", msg, arg)
}
