package fizzbuzz

import (
	"errors"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestService_GenerateFizzBuzz(t *testing.T) {
	var ctrl = gomock.NewController(t)

	type fields struct {
		stat func() adapters.StatsRepository
	}
	type args struct {
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "division by zero case",
			fields: fields{
				stat: func() adapters.StatsRepository {
					return adapters.NewMockStatsRepository(ctrl)
				},
			},
			args: args{
				int1:  0,
				int2:  0,
				limit: 10,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error incrementing request count",
			fields: fields{
				stat: func() adapters.StatsRepository {
					m := adapters.NewMockStatsRepository(ctrl)
					m.EXPECT().IncrementRequestCount(3, 5, 15, "Fizz", "Buzz").Return(errors.New("failed")).Times(1)
					return m
				},
			},
			args: args{
				int1:  3,
				int2:  5,
				limit: 15,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid case",
			fields: fields{
				stat: func() adapters.StatsRepository {
					m := adapters.NewMockStatsRepository(ctrl)
					m.EXPECT().IncrementRequestCount(3, 5, 15, "Fizz", "Buzz").Return(nil).Times(1)
					return m
				},
			},
			args: args{
				int1:  3,
				int2:  5,
				limit: 15,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			want:    "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := NewFizzBuzzService(
				tt.fields.stat(),
			)
			got, err := fb.GenerateFizzBuzz(tt.args.int1, tt.args.int2, tt.args.limit, tt.args.str1, tt.args.str2)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateFizzBuzz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateFizzBuzz() got = %v, want %v", got, tt.want)
			}
		})
	}
}
