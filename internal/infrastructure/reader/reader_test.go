package reader

import (
	// golang package
	"io"
	"testing"

	// external package
	"github.com/stretchr/testify/assert"
)

func TestNewReader(t *testing.T) {
	want := &Reader{}
	got := NewReader()
	assert.Equal(t, want, got)
}

func TestReader_ReadAll(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "success",
			args: args{
				input: io.LimitReader(&io.LimitedReader{}, 10),
			},
			want: []byte{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := &Reader{}
			got, err := r.ReadAll(test.args.input)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
