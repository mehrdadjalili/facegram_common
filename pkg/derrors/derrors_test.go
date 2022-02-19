package derrors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/matiniiuu/mcommon/pkg/translator/messages"
)

func TestHttpError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		name  string
		args  args
		want1 string
		want2 int
	}{
		{
			name: "server error",
			args: args{
				err: New(KindNotFound, "page not found"),
			},
			want1: "page not found",
			want2: http.StatusNotFound,
		},
		{
			name: "server error",
			args: args{
				err: fmt.Errorf("something: %w", New(KindNotFound, "page not found")),
			},
			want1: "page not found",
			want2: http.StatusNotFound,
		},
		{
			name: "other errors",
			args: args{
				err: errors.New("something"),
			},
			want1: messages.GeneralError,
			want2: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := HttpError(tt.args.err)
			if got1 != tt.want1 {
				t.Errorf("HttpError() got = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("HttpError() got1 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
