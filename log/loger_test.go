package log

import (
	"testing"
)

func TestInfo(t *testing.T) {
	type args struct {
		format string
		a      interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "InfoLogger", args: args{
			format: "infoLog: %+v",
			a:      nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.format, tt.args.a)
		})
	}
}
