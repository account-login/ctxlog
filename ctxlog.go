package ctxlog

import (
	"context"
	"fmt"
	"log"
)

type keyType struct{}

var key keyType

// get logging context
func Ctx(ctx context.Context) string {
	v := ctx.Value(key)
	if v == nil {
		return ""
	} else {
		return v.(string)
	}
}

// set logging context
func Push(ctx context.Context, s string) context.Context {
	return context.WithValue(ctx, key, Ctx(ctx)+s)
}

func Pushf(ctx context.Context, format string, args ...interface{}) context.Context {
	return context.WithValue(ctx, key, Ctx(ctx)+fmt.Sprintf(format, args...))
}

// logging
func Print(ctx context.Context, v ...interface{}) {
	log.Print(append([]interface{}{Ctx(ctx)}, v...)...)
}

func Debugf(ctx context.Context, f string, a ...interface{}) {
	log.Printf("[DEBUG]  %s "+f, append([]interface{}{Ctx(ctx)}, a...)...)
}

func Infof(ctx context.Context, f string, a ...interface{}) {
	log.Printf("[INFO]   %s "+f, append([]interface{}{Ctx(ctx)}, a...)...)
}

func Warnf(ctx context.Context, f string, a ...interface{}) {
	log.Printf("[WARN]   %s "+f, append([]interface{}{Ctx(ctx)}, a...)...)
}

func Noticef(ctx context.Context, f string, a ...interface{}) {
	log.Printf("[NOTICE] %s "+f, append([]interface{}{Ctx(ctx)}, a...)...)
}

func Errorf(ctx context.Context, f string, a ...interface{}) {
	log.Printf("[ERROR]  %s "+f, append([]interface{}{Ctx(ctx)}, a...)...)
}

func Fatal(ctx context.Context, v ...interface{}) {
	log.Fatal(append([]interface{}{Ctx(ctx)}, v...)...)
}

// TODO: more
// TODO: logger wrapper
