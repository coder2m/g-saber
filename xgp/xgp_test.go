package xgp

import (
	"github.com/coder2z/g-saber/xconsole"
	"testing"
	"time"
)

func TestPG(t *testing.T) {
	Go(func() {
		xconsole.Red("test1")
		time.Sleep(2 * time.Second)
		xconsole.Red("test1 Ok")
	})

	Go(func() {
		xconsole.Red("test2")
		time.Sleep(3 * time.Second)
		xconsole.Red("test2 Ok")
	})

	Go(func() {
		xconsole.Red("test3")
		time.Sleep(4 * time.Second)
		xconsole.Red("test3 Ok")
	})

	Wait()
}
