package loglevel

import (
	"fmt"
	"io"
	"sync"
	"os"
	"time"
	"strings"
)

const (
	DEBUG = 3
	INFO = 2
	WARN = 1
	ERROR = 0
)

type Logger struct {
	Name  string
	Level *Level
	TimeFormat string
	PrintTime bool
	PrintName bool
	PrintLevel bool
	mu sync.Mutex
	out io.Writer
	buf []byte
}

type Level struct {
  Num int
  Name string
}

func (l *Level) SetLevel(level int) *Level {
  switch level {
    case 0:
      l.Num = 0
      l.Name = "ERROR"
    case 1:
      l.Num = 1
      l.Name = "WARN"
    case 2:
      l.Num = 2
      l.Name = "INFO"
    case 3:
      l.Num = 3
      l.Name = "DEBUG"
    default:
      l.Num = 1
      l.Name = "WARN"
  }
  return l
}

// Writes the specified string to std and colors it accordingly
func (l *Logger) Output(color, level, str string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	nam := ""
	lvl := ""
	tim := "" 
	if l.PrintName {
		nam = l.Name
	}
	if l.PrintLevel {
		lvl = level
	}
	if l.PrintTime {
		tim = time.Now().Format(l.TimeFormat)
	}
	l.buf = []byte(strings.Replace(strings.Trim(fmt.Sprintf("%s %s %s %s %s %s", color, tim, lvl, nam, str, CLR_N), " "), "  ", " ", -1))
	_, err := l.out.Write(l.buf)
	return err
}
// Return the level
func (l *Logger) GetLevel() *Level {
	return l.Level
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Debugf(format string, message ...interface{}) {
  if l.Level.Num >= DEBUG {
		l.Output(CLR_G, "DEBUG", fmt.Sprintf(format, message...))
  }
}

// Prints a debug message on a new line
func (l *Logger) Debug(message ...interface{}) {
  if l.Level.Num >= DEBUG {
		l.Output(CLR_G, "DEBUG", fmt.Sprintln(message...))
  }
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Infof(format string, message ...interface{}) {
  if l.Level.Num >= INFO {
		l.Output(CLR_W, "INFO", fmt.Sprintf(format, message...))
  }
}

// Prints n info message on a new line
func (l *Logger) Info(message ...interface{}) {
  if l.Level.Num >= INFO {
		l.Output(CLR_W, "INFO",fmt.Sprintln(message...))
  }
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Warnf(format string, message ...interface{}) {
	if l.Level.Num >= WARN {
		l.Output(CLR_Y, "WARN", fmt.Sprintf(format, message...))
	}
}

// Prints a warning message on a new line
func (l *Logger) Warn(message ...interface{}) {
  if l.Level.Num >= WARN {
		l.Output(CLR_Y, "WARN", fmt.Sprintln(message...))
  }
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Errorf(format string, message ...interface{}) {
	if l.Level.Num >= ERROR {
		l.Output(CLR_R, "ERROR", fmt.Sprintf(format, message...))
	}
}

// Prints an error message on a new line
func (l *Logger) Error(message ...interface{}) {
  if l.Level.Num >= ERROR {
    l.Output(CLR_R, "ERROR", fmt.Sprintln(message...))
  }
}

type Empty struct {}

// Create return logger
func New() *Logger {
	return &Logger{
		out: os.Stderr,
		Name:  "",
		Level: &Level{1, "INFO"},
		PrintTime: true,
		PrintName: true,
		PrintLevel: true,
		TimeFormat: time.RFC3339,
	}
}
