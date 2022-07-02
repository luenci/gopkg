package gopkg

import (
	"context"
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {

	type args struct {
		ctx     context.Context
		handler func()
	}
	tests := []struct {
		name string
		args args
	}{
		{"Go1", args{
			ctx:     context.Background(),
			handler: func() { panic("test panic") }},
		},
		{"Go2", args{
			ctx:     context.Background(),
			handler: func() { fmt.Println("Success") }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Go(tt.args.ctx, tt.args.handler)
		})
	}
}
