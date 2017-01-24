package loglevel

import (
	"testing"
	"fmt"
)

func TestOut(t *testing.T) {
	l := New()
	l.Out(fmt.Sprintln("Out message"))
}

func TestOutput(t *testing.T) {
	l := New()
	l.Output(CLR_W, "OUTPUT", fmt.Sprint("Output message"))
}

// Debug
func TestDebug(t *testing.T) {
	l := New().SetLevel(DEBUG)
	l.Debug("Simple message")
}

func TestDebugf(t *testing.T) {
	l := New().SetLevel(DEBUG)
	l.Debugf("%s", "Format message")
}

// Info
func TestInfo(t *testing.T) {
	l := New().SetLevel(INFO)
	l.Info("Simple message")
}

func TestInfof(t *testing.T) {
	l := New().SetLevel(INFO)
	l.Infof("%s", "Format message")
}

// Warn
func TestWarn(t *testing.T) {
	l := New().SetLevel(WARN)
	l.Warn("Simple message")
}

func TestWarnf(t *testing.T) {
	l := New().SetLevel(WARN)
	l.Warnf("%s", "Format message")
}

// Error
func TestError(t *testing.T) {
	l := New().SetLevel(ERROR)
	l.Error("Simple message")
}

func TestErrorf(t *testing.T) {
	l := New().SetLevel(ERROR)
	l.Errorf("%s", "Format message")
}