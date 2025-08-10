package http

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidator_Validate(t *testing.T) {
	type fields struct {
		validator *validator.Validate
	}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid input",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				i: struct {
					Empty bool   `json:"-" validate:"-"`
					Int1  int    `json:"int1" validate:"required,min=1,max=100"`
					Int2  int    `json:"int2" validate:"required,min=1,max=100"`
					Limit int    `json:"limit" validate:"required,min=1,max=1000"`
					Str1  string `json:"str1" validate:"required,min=1,max=50"`
					Str2  string `json:"str2" validate:"required,min=1,max=50"`
				}{
					Int1:  3,
					Int2:  5,
					Limit: 15,
					Str1:  "Fizz",
					Str2:  "Buzz",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid input - missing required field",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				i: struct {
					Int1  int    `json:"int1" validate:"required,min=1,max=100"`
					Int2  int    `json:"int2" validate:"required,min=1,max=100"`
					Limit int    `json:"limit" validate:"required,min=1,max=1000"`
					Str1  string `json:"str1" validate:"required,min=1,max=50"`
					Str2  string `json:"str2" validate:"required,min=1,max=50"`
				}{
					Int1:  0, // Int1 is missing
					Int2:  5,
					Limit: 15,
					Str1:  "Fizz",
					Str2:  "Buzz",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid input - must be less than",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				i: struct {
					Int1  int    `json:"int1" validate:"required,min=1,max=100"`
					Int2  int    `json:"int2" validate:"required,min=1,max=100"`
					Limit int    `json:"limit" validate:"required,min=1,max=1000"`
					Str1  string `json:"str1" validate:"required,min=1,max=50"`
					Str2  string `json:"str2" validate:"required,min=1,max=50"`
				}{
					Int1:  3,
					Int2:  105, // Int2 is out of range
					Limit: 15,
					Str1:  "Fizz",
					Str2:  "Buzz",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid input - string is required",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				i: struct {
					Int1  int    `json:"int1" validate:"-"`
					Int2  int    `json:"int2" validate:"required,min=1,max=100"`
					Limit int    `json:"limit" validate:"required,min=1,max=1000"`
					Str1  string `json:"str1" validate:"required,min=1,max=50"`
					Str2  string `json:"str2" validate:"required,min=1,max=50"`
				}{
					Int1:  3,
					Int2:  5,
					Limit: 15,
					Str2:  "Buzz",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid input - must be greater than",
			fields: fields{
				validator: validator.New(),
			},
			args: args{
				i: struct {
					Empty int `json:"-"`
					Limit int `json:"limit" validate:"required,min=1,max=1000"`
				}{
					Limit: 0, // Limit must be bigger than 0
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv := NewValidator()
			if err := cv.Validate(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
