package ctxlog

import (
	"context"
	"fmt"
	"log"
)

type ctxKey struct{}

var ckey interface{} = ctxKey{}

type logCtx struct {
	context.Context
	val string // NOTE: val is not interface{}, saves 1 alloc from context.WithValue()
}

func (lc *logCtx) Value(key interface{}) interface{} {
	if ckey == key {
		return lc.val
	} else {
		return lc.Context.Value(key)
	}
}

// get logging context
func Ctx(ctx context.Context) string {
	if vctx, ok := ctx.(*logCtx); ok {
		return vctx.val
	}
	if pval := ctx.Value(ckey); pval != nil {
		return pval.(string)
	}
	return ""
}

// set logging context
func Push(ctx context.Context, s string) context.Context {
	return &logCtx{ctx, Ctx(ctx) + s}
}

func Pushf(ctx context.Context, format string, args ...interface{}) context.Context {
	return Push(ctx, fmt.Sprintf(format, args...))
}

// logging
func Print(ctx context.Context, v ...interface{}) {
	log.Print(append([]interface{}{Ctx(ctx)}, v...)...)
}

//func Debugf(ctx context.Context, f string, a ...interface{}) {
//	log.Printf("[DEBUG]  %s "+f, append([]interface{}{Ctx(ctx)}, a...)...)
//}

// NOTE: merge format string into args slice to save 1 alloc
func Debugf(ctx context.Context, a ...interface{}) {
	f := "[DEBUG]  %s " + a[0].(string)
	a[0] = Ctx(ctx)
	log.Printf(f, a...)
}

func Infof(ctx context.Context, a ...interface{}) {
	f := "[INFO]   %s " + a[0].(string)
	a[0] = Ctx(ctx)
	log.Printf(f, a...)
}

func Warnf(ctx context.Context, a ...interface{}) {
	f := "[WARN]   %s " + a[0].(string)
	a[0] = Ctx(ctx)
	log.Printf(f, a...)
}

func Noticef(ctx context.Context, a ...interface{}) {
	f := "[NOTICE] %s " + a[0].(string)
	a[0] = Ctx(ctx)
	log.Printf(f, a...)
}

func Errorf(ctx context.Context, a ...interface{}) {
	f := "[ERROR]  %s " + a[0].(string)
	a[0] = Ctx(ctx)
	log.Printf(f, a...)
}

func Fatal(ctx context.Context, v ...interface{}) {
	log.Fatal(append([]interface{}{Ctx(ctx)}, v...)...)
}

// TODO: more
// TODO: logger wrapper
