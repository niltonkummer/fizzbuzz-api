package repository

import (
	"reflect"
	"testing"

	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

func TestInMemoryStatsRepository_GetMostFrequentRequest(t *testing.T) {
	type fields struct {
		stats map[model.FizzBuzzRequest]int
	}
	tests := []struct {
		name      string
		fields    fields
		wantStats *model.StatsResult
		wantErr   bool
	}{
		{
			name: "No requests",
			fields: fields{
				stats: make(map[model.FizzBuzzRequest]int),
			},
			wantStats: nil,
			wantErr:   false,
		},
		{
			name: "Single request",
			fields: fields{
				stats: map[model.FizzBuzzRequest]int{
					{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}: 10,
				},
			},
			wantStats: &model.StatsResult{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
				Hits:  10,
			},
			wantErr: false,
		},
		{
			name: "Multiple requests",
			fields: fields{
				stats: map[model.FizzBuzzRequest]int{
					{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}: 10,
					{Int1: 2, Int2: 7, Limit: 20, Str1: "Foo", Str2: "Bar"}:   5,
					{Int1: 4, Int2: 6, Limit: 30, Str1: "Qux", Str2: "Quux"}:  15,
				},
			},
			wantStats: &model.StatsResult{
				Int1:  4,
				Int2:  6,
				Limit: 30,
				Str1:  "Qux",
				Str2:  "Quux",
				Hits:  15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewInMemoryStatsRepository(tt.fields.stats)
			gotStats, err := r.GetMostFrequentRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMostFrequentRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStats, tt.wantStats) {
				t.Errorf("GetMostFrequentRequest() gotStats = %v, want %v", gotStats, tt.wantStats)
			}
		})
	}
}

func TestInMemoryStatsRepository_IncrementRequestCount(t *testing.T) {
	type fields struct {
		stats map[model.FizzBuzzRequest]int
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
		wantErr bool
	}{
		{
			name: "Increment existing request",
			fields: fields{
				stats: map[model.FizzBuzzRequest]int{
					{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}: 10,
				},
			},
			args: args{
				int1:  3,
				int2:  5,
				limit: 15,
				str1:  "Fizz",
				str2:  "Buzz",
			},
		},
		{
			name: "Increment new request",
			fields: fields{
				stats: make(map[model.FizzBuzzRequest]int),
			},
			args: args{
				int1:  2,
				int2:  7,
				limit: 20,
				str1:  "Foo",
				str2:  "Bar",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewInMemoryStatsRepository(tt.fields.stats)
			if err := r.IncrementRequestCount(tt.args.int1, tt.args.int2, tt.args.limit, tt.args.str1, tt.args.str2); (err != nil) != tt.wantErr {
				t.Errorf("IncrementRequestCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryStatsRepository_ResetStats(t *testing.T) {
	type fields struct {
		stats map[model.FizzBuzzRequest]int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Reset stats",
			fields: fields{
				stats: map[model.FizzBuzzRequest]int{
					{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}: 10,
					{Int1: 2, Int2: 7, Limit: 20, Str1: "Foo", Str2: "Bar"}:   5,
					{Int1: 4, Int2: 6, Limit: 30, Str1: "Qux", Str2: "Quux"}:  15,
				},
			},
			wantErr: false,
		},
		{
			name:    "Reset empty stats",
			fields:  fields{stats: make(map[model.FizzBuzzRequest]int)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewInMemoryStatsRepository(tt.fields.stats)
			if err := r.ResetStats(); (err != nil) != tt.wantErr {
				t.Errorf("ResetStats() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
