package notifier

import (
	"context"
	"testing"
)

func Test_Throw(t *testing.T) {
	ctx := context.Background()
	if err := Listen(ctx, 0); err != nil {
		t.Fatalf("%+v", err)
	}
}
