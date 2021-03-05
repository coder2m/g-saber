package xlog

import (
	"github.com/coder2m/g-saber/xlog/rotate"
	"io"
)

func newRotate(o *options) io.Writer {
	rotateLog := rotate.NewLogger()
	rotateLog.Filename = o.Filename()
	rotateLog.MaxSize = o.MaxSize // MB
	rotateLog.MaxAge = o.MaxAge   // days
	rotateLog.MaxBackups = o.MaxBackup
	rotateLog.Interval = o.Interval
	rotateLog.LocalTime = true
	rotateLog.Compress = false
	return rotateLog
}
