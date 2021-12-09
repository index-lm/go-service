package log

import "testing"

func TestLog(t *testing.T) {
	Init()
	Sugar.Info("这是一个打印")
}
