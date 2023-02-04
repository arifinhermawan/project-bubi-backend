package golang

import (
	// golang package
	"encoding/json"
	"testing"

	// external package
	"github.com/stretchr/testify/assert"
)

func TestGolang_JsonMarshal(t *testing.T) {
	mockResult, _ := json.Marshal("abc")
	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "success",
			args: args{input: "abc"},
			want: mockResult,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := &Golang{}
			got, err := g.JsonMarshal(test.args.input)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestGolang_JsonUnmarshal(t *testing.T) {
	type args struct {
		input []byte
		dest  interface{}
	}

	mockInput, _ := json.Marshal("abc")
	var dest string
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				input: mockInput,
				dest:  &dest,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := &Golang{}
			err := g.JsonUnmarshal(test.args.input, test.args.dest)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
