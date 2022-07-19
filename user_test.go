package cheevos_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos"
)

func TestCreatingAValidUserWithSucceeds(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.UserService{}
	_, err := svc.SignUp(ctx, cheevos.SignUpRequest{
		Username: "test",
		Password: "testtest",
	})
	if err != nil {
		t.Fatal(err)
	}
}
