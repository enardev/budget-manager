package validation

import (
	"bytes"
	"io"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestNewValidator(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "when create a new validator, then get a new validator with success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidator(); !assert.NotNil(t, got) {
				t.Errorf("NewValidator() = %v", got)
			}
		})
	}
}

func TestValidator_DecodeAndValidate(t *testing.T) {
	type testStruct struct {
		Id int `json:"id" validate:"required,number"`
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "given a valid struct, then get nil error",
			args: args{
				r: bytes.NewBuffer([]byte("{\"id\":20}")),
			},
			wantErr: false,
		},
		{
			name: "given an invalid struct, then get error",
			args: args{
				r: bytes.NewBuffer([]byte("{\"id\":20")),
			},
			wantErr: true,
		},
		{
			name: "given a valid struct with invalid field, then get error",
			args: args{
				r: bytes.NewBuffer([]byte("{\"test\":\"test\"}")),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result testStruct
			v := NewValidator()
			if err := v.DecodeAndValidate(tt.args.r, &result); (err != nil) != tt.wantErr {
				t.Errorf("Validator.DecodeAndValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
