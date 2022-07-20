package cheevos

import "context"

type Database interface {
	Call(context.Context, func(ctx context.Context, tx Transaction) error) error
}

type Transaction interface{}
