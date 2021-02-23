package logger

import (
	"context"
	"io"
	"log"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevleError
	LevelFetal
	LevelPanic
)

func (l Level) string() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevleError:
		return "error"
	case LevelFetal:
		return "fetal"
	case LevelPanic:
		return "panic"
	}
	return ""

}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}
