package ctxlog

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	buf := bytes.Buffer{}
	log.SetFlags(0)
	log.SetOutput(&buf)

	ctx := Push(context.Background(), "[c1]")
	Infof(ctx, "info %v", 1)
	ctx = Pushf(ctx, "[c:%v]", 2)
	Errorf(ctx, "err %v", 3)
	Debugf(ctx, "dbg")

	expected := `[INFO]   [c1] info 1
[ERROR]  [c1][c:2] err 3
[DEBUG]  [c1][c:2] dbg
`
	assert.Equal(t, expected, buf.String())
}

func TestCtxChain(t *testing.T) {
	buf := bytes.Buffer{}
	log.SetFlags(0)
	log.SetOutput(&buf)

	ctx := Push(context.Background(), "[c1]")
	ctx, _ = context.WithTimeout(ctx, time.Second)
	ctx = Push(ctx, "[c2]")
	ctx = context.WithValue(ctx, 1, 2)
	assert.Equal(t, "[c1][c2]", Ctx(ctx))
}
