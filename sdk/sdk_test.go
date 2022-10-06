package sdk_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/haleyrc/cheevos/sdk"
)

func TestSignUp(t *testing.T) {
	ctx := context.Background()

	c := sdk.Client{C: &http.Client{
		Timeout: 5 * time.Second,
	}}

	err := c.SignUp(ctx, "ryan3", "12345678")
	if err != nil {
		t.Errorf("expected err to be nil, but it wasn't")
	}
}
