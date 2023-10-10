package ares

import (
	"context"
	"fmt"
)

var logger Logger

func init() {
	logger = fmtLogger{}
}

func Info(format string, a ...any) {
	logger.Info(format, a)
}

type Logger interface {
	Info(format string, a ...any)
	CtxInfo(ctx context.Context, format string, a ...any)
}

type fmtLogger struct {
}

func (fl fmtLogger) Info(format string, a ...any) {
	fmt.Println(fmt.Sprintf(format, a))
}

func (fl fmtLogger) CtxInfo(ctx context.Context, format string, a ...any) {
	fmt.Println(fmt.Sprintf(format, a))
}
