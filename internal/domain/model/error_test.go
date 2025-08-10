package model

import "testing"

func TestError_Error(t *testing.T) {
	type fields struct {
		Code    string
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "no requests found error",
			fields: fields{
				Code:    ErrNoRequestsFound.Code,
				Message: ErrNoRequestsFound.Message,
			},
			want: ErrNoRequestsFound.Message,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &Error{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
