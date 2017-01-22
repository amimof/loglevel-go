package loglevel

import (
	"fmt"
	"io"
	"sync"
	"os"
	"time"
)

const (
	DEBUG = 3
	INFO = 2
	WARN = 1
	ERROR = 0
)

type Logger struct {
	Name  string
	Level int
	TimeFormat string
	PrintTime bool
	PrintName bool
	PrintLevel bool
	UseColors bool
	mu sync.Mutex
	out io.Writer
	buf []byte
}

func (l *Logger) SetLevel(level int) *Logger {
  l.Level = level
  return l
}

// Writes the specified string to std and colors it accordingly
func (l *Logger) Output(color, level, str string) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	if l.UseColors {
		l.buf = append(l.buf, color...)	
	}
	if l.PrintTime {
		l.buf = append(l.buf, time.Now().Format(l.TimeFormat)...)
		l.buf = append(l.buf, ' ')
	}
	if l.PrintLevel {
		l.buf = append(l.buf, level...)
		l.buf = append(l.buf, ' ')
	}
	if l.PrintName && len(l.Name) > 0 {
		l.buf = append(l.buf, l.Name...)
		l.buf = append(l.buf, ' ')
	}
	l.buf = append(l.buf, str...)
	if l.UseColors {
		l.buf = append(l.buf, CLR_N...)
	}
	if len(str) == 0 || str[len(str)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	n, err = l.out.Write(l.buf)
	return
}

// Writes str to stdout
func (l *Logger) Out(str string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = []byte(str)
	_, err := l.out.Write(l.buf)
	return err
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Debugf(format string, message ...interface{}) {
  if l.Level >= DEBUG {
		l.Output(CLR_G, "DEBUG", fmt.Sprintf(format, message...))
  }
}

// Prints a debug message on a new line
func (l *Logger) Debug(message ...interface{}) {
  if l.Level >= DEBUG {
		l.Output(CLR_G, "DEBUG", fmt.Sprint(message...))
  }
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Infof(format string, message ...interface{}) {
  if l.Level >= INFO {
		l.Output(CLR_W, "INFO", fmt.Sprintf(format, message...))
  }
}

// Prints n info message on a new line
func (l *Logger) Info(message ...interface{}) {
  if l.Level >= INFO {
		l.Output(CLR_W, "INFO",fmt.Sprint(message...))
  }
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Warnf(format string, message ...interface{}) {
	if l.Level >= WARN {
		l.Output(CLR_Y, "WARN", fmt.Sprintf(format, message...))
	}
}

// Prints a warning message on a new line
func (l *Logger) Warn(message ...interface{}) {
  if l.Level >= WARN {
		l.Output(CLR_Y, "WARN", fmt.Sprint(message...))
  }
}

// Prints according to fmt.Sprintf format specifier and returns the resulting string
func (l *Logger) Errorf(format string, message ...interface{}) {
	if l.Level >= ERROR {
		l.Output(CLR_R, "ERROR", fmt.Sprintf(format, message...))
	}
	os.Exit(1)
}

// Prints an error message on a new line followed by Exit
func (l *Logger) Error(message ...interface{}) {
	if l.Level >= ERROR {
		l.Output(CLR_R, "ERROR", fmt.Sprint(message...))
	}
	os.Exit(1)
}

// Same as fmt.Sprint
func (l *Logger) Print(message ...interface{}) {
	l.Out(fmt.Sprint(message...))
}

// Same as fmt.Sprintln
func (l *Logger) Println(message ...interface{}) {
	l.Out(fmt.Sprintln(message...))
}

// Same as fmt.Sprintf
func (l *Logger) Printf(format string, message ...interface{}) {
	l.Out(fmt.Sprintf(format, message...))
}

// Same as fmt.Sprint followed by panic
func (l *Logger) Panic(message ...interface{}) {
	l.Out(fmt.Sprint(message...))
	panic(message)
}

// Same as fmt.Sprintln followed by panic
func (l *Logger) Panicln(message ...interface{}) {
	l.Out(fmt.Sprintln(message...))
	panic(message)
}

// Same as fmt.Sprintf followed by panic
func (l *Logger) Panicf(format string, message ...interface{}) {
	l.Out(fmt.Sprintf(format, message...))
	panic(message)
}

// Create return logger
func New() *Logger {
	return &Logger{
		out: os.Stderr,
		Name:  "",
		Level: INFO,
		PrintTime: true,
		PrintName: true,
		PrintLevel: true,
		UseColors: true,
		TimeFormat: time.RFC3339,
	}
}
