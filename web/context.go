package web

import "context"

type contextKey int

const (
	currentUserKey = iota
)

func GetCurrentUser(ctx context.Context) string {
	tmp := ctx.Value(currentUserKey)
	if tmp == nil {
		panic("critical: GetCurrentUser called without a suitable context")
	}
	s, ok := tmp.(string)
	if !ok {
		panic("critical: GetCurrentUser got a current user that isn't a string")
	}
	return s
}
