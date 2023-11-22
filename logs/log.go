package logs

import (
	"context"
	"fmt"
)

func CtxFatal(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func CtxNotice(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func Fatal(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func Error(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func Warn(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func Notice(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func Info(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func Debug(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}

func Trace(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v))
}
