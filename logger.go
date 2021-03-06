// Package logger implements a logger that takes a writer and a func to prefix its output.
package logger

import (
	"fmt"
	"io"
	"sync"
)

// A Logger is a mutex-protected logger.
type Logger struct {
	m sync.Mutex
	w io.Writer
	f func() string
}

// New returns a new Logger.
func New(w io.Writer, f func() string) *Logger {
	return &Logger{f: f, w: w}
}

// Printf prefixes the Logger's output with the string returned from the Logger's f func.
func (l *Logger) Printf(format string, a ...interface{}) {
	l.m.Lock()
	fmt.Fprintf(l.w, format, prepend(l.f(), a)...)
	l.m.Unlock()
}

// Println prefixes the Logger's output with the string returned from the Loggers' f func.
func (l *Logger) Println(a ...interface{}) {
	l.m.Lock()
	fmt.Fprintln(l.w, prepend(l.f(), a)...)
	l.m.Unlock()
}

func prepend(v interface{}, a []interface{}) []interface{} {
	return append([]interface{}{v}, a...)
}
