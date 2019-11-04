package ctxlog

import (
	"context"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkLogSimpleStringWithoutCtx(b *testing.B) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debugf(ctx, "debug")
	}
	b.ReportAllocs()
}

func BenchmarkLogWithoutCtx(b *testing.B) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debugf(ctx, "debug %v", i)
	}
	b.ReportAllocs()
}

func BenchmarkPushfAndLog(b *testing.B) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := Pushf(ctx, "[context1]")
		ctx = Pushf(ctx, "[context2]")
		ctx = Pushf(ctx, "[context3]")
		Debugf(ctx, "debug %v", i)
	}
	b.ReportAllocs()
}

func BenchmarkPushAndLog(b *testing.B) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := Push(ctx, "[context1]")
		ctx = Push(ctx, "[context2]")
		ctx = Push(ctx, "[context3]")
		Debugf(ctx, "debug %v", i)
	}
	b.ReportAllocs()
}

func BenchmarkPushf(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := Pushf(ctx, "[context1]")
		func() {
			ctx := Pushf(ctx, "[context2]")
			func() {
				ctx := Pushf(ctx, "[context3]")
				func() {
					ctx := Pushf(ctx, "[context4]")
					func() {
						ctx := Pushf(ctx, "[context5]")
						_ = Ctx(ctx)
					}()
				}()
			}()
		}()
	}
	b.ReportAllocs()
}

func BenchmarkPush(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := Pushf(ctx, "[context1]")
		func() {
			ctx := Pushf(ctx, "[context2]")
			func() {
				ctx := Pushf(ctx, "[context3]")
				func() {
					ctx := Pushf(ctx, "[context4]")
					func() {
						ctx := Pushf(ctx, "[context5]")
						_ = Ctx(ctx)
					}()
				}()
			}()
		}()
	}
	b.ReportAllocs()
}

func BenchmarkCtxGet(b *testing.B) {
	ctx := context.Background()
	ctx = Pushf(ctx, "[context1]")
	ctx = Pushf(ctx, "[context2]")
	ctx = Pushf(ctx, "[context3]")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Ctx(ctx)
	}
	b.ReportAllocs()
}

func BenchmarkLogWithCtx(b *testing.B) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	ctx := context.Background()
	ctx = Pushf(ctx, "[context1]")
	ctx = Pushf(ctx, "[context2]")
	ctx = Pushf(ctx, "[context3]")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debugf(ctx, "debug %v", i)
	}
	b.ReportAllocs()
}
