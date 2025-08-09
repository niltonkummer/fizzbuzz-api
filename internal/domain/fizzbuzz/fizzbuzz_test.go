package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz_Calculate(t *testing.T) {
	type args struct {
		int1  int
		int2  int
		limit int
		str1  string
		str2  string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		error assert.ErrorAssertionFunc
		skip  bool
	}{
		{
			name: "division by zero case",
			args: args{
				int1: 0, int2: 0, limit: 10, str1: "Fizz", str2: "Buzz",
			},
			want:  "",
			error: assert.Error,
		},
		{
			name:  "limit is zero",
			args:  args{int1: 3, int2: 5, limit: 0, str1: "Fizz", str2: "Buzz"},
			want:  "",
			error: assert.Error,
		},
		{
			name:  "negative limit",
			args:  args{int1: 3, int2: 5, limit: -10, str1: "Fizz", str2: "Buzz"},
			want:  "",
			error: assert.Error,
		},
		{
			name:  "int1 equals int2",
			args:  args{int1: 3, int2: 3, limit: 5, str1: "Fizz", str2: "Buzz"},
			want:  "1,2,FizzBuzz,4,5",
			error: assert.NoError,
		},
		{
			name:  "empty strings",
			args:  args{int1: 3, int2: 5, limit: 5, str1: "", str2: ""},
			want:  "1,2,,4,",
			error: assert.NoError,
		},
		{
			name:  "high limit",
			args:  args{int1: 3, int2: 5, limit: 100000, str1: "Fizz", str2: "Buzz"},
			want:  "...",
			skip:  true,
			error: assert.NoError,
		},
		{
			name:  "unicode strings",
			args:  args{int1: 3, int2: 5, limit: 15, str1: "🎉", str2: "🚀"},
			want:  "1,2,🎉,4,🚀,🎉,7,8,🎉,🚀,11,🎉,13,14,🎉🚀",
			error: assert.NoError,
		},
		{
			name:  "limit is one",
			args:  args{int1: 1, int2: 2, limit: 1, str1: "a", str2: "b"},
			want:  "a",
			error: assert.NoError,
		},
		{
			name:  "multiple of all",
			args:  args{int1: 1, int2: 2, limit: 10, str1: "a", str2: "b"},
			want:  "a,ab,a,ab,a,ab,a,ab,a,ab",
			error: assert.NoError,
		},
		{
			name:  "original fizzbuzz",
			args:  args{int1: 3, int2: 5, limit: 21, str1: "Fizz", str2: "Buzz"},
			want:  "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz,16,17,Fizz,19,Buzz,Fizz",
			error: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := NewFizzBuzz()
			got, err := fb.Calculate(tt.args.int1, tt.args.int2, tt.args.limit, tt.args.str1, tt.args.str2)
			tt.error(t, err)
			if tt.skip {
				t.Skip()
			}
			assert.Equal(t, tt.want, got, "GenerateFizzBuzz() = %v, want %v", got, tt.want)

		})
	}
}
