package log

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Fields map[string]interface{}

type Logger interface {
	Debug(ctx context.Context, msg string, fields Fields)
	Error(ctx context.Context, msg string, err error)
	Log(ctx context.Context, msg string, fields Fields)
}

type NullLogger struct{}

func (nl NullLogger) Debug(ctx context.Context, msg string, fields Fields) {}
func (nl NullLogger) Error(ctx context.Context, msg string, err error)     {}
func (nl NullLogger) Log(ctx context.Context, msg string, fields Fields)   {}

type JSONLogger struct {
	EnableDebug bool
	Output      io.Writer
}

func (jl JSONLogger) Debug(ctx context.Context, msg string, fields Fields) {
	if !jl.EnableDebug {
		return
	}
	jl.log(ctx, msg, fields)
}

func (jl JSONLogger) Error(ctx context.Context, msg string, err error) {
	jl.log(ctx, msg, Fields{"Error": err.Error()})
}

func (jl JSONLogger) Log(ctx context.Context, msg string, fields Fields) {
	jl.log(ctx, msg, fields)
}

func (jl JSONLogger) log(ctx context.Context, msg string, fields Fields) {
	body := map[string]interface{}{"Message": msg}
	if len(fields) != 0 {
		body["Fields"] = fields
	}
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "log: json marshaling failed: %v: %#v\n", err, body)
	}
	fmt.Fprintln(jl.Output, string(b))
}
