package repository

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

var redisClient *redis.Client

func init() {
	// Initialize the Redis client here if needed for tests
	// This is just a placeholder, actual initialization should be done in the test setup
	redisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"), // Adjust the address as needed
	})
}

func TestRedisStatsRepository_GetMostFrequentRequest(t *testing.T) {
	type fields struct {
		client *redis.Client
	}
	redisClient.FlushAll(context.Background())

	tests := []struct {
		name       string
		fields     fields
		beforeCall func()
		wantStats  *model.StatsResult
		wantErr    bool
	}{
		{
			name: "No requests",
			fields: fields{
				client: redisClient,
			},
			wantStats: nil,
			wantErr:   false,
		},
		{
			name: "Single request",
			fields: fields{
				client: redisClient,
			},
			beforeCall: func() {
				// Simulate a single request in Redis
				redisClient.ZAdd(redisClient.Context(), "fizzbuzz:stats", &redis.Z{
					Score:  10,
					Member: "3,5,15,Fizz,Buzz",
				})
			},
			wantStats: &model.StatsResult{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
				Hits:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisStatsRepository(tt.fields.client)
			if tt.beforeCall != nil {
				tt.beforeCall()
			}
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

func TestRedisStatsRepository_IncrementRequestCount(t *testing.T) {
	type fields struct {
		client *redis.Client
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
				client: redisClient,
			},
			args: args{
				int1:  3,
				int2:  5,
				limit: 15,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisStatsRepository(tt.fields.client)
			if err := r.IncrementRequestCount(tt.args.int1, tt.args.int2, tt.args.limit, tt.args.str1, tt.args.str2); (err != nil) != tt.wantErr {
				t.Errorf("IncrementRequestCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisStatsRepository_ResetStats(t *testing.T) {
	type fields struct {
		client *redis.Client
	}
	tests := []struct {
		name        string
		fields      fields
		checkStatus func()
		wantErr     bool
	}{
		{
			name: "Reset stats",
			fields: fields{
				client: redisClient,
			},
			checkStatus: func() {
				if len(redisClient.ZRevRangeWithScores(context.TODO(), "fizzbuzz:stats", 0, 0).Val()) > 0 {
					t.Errorf("Expected no stats after reset, but found some")
				}
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedisStatsRepository(tt.fields.client)

			if err := r.ResetStats(); (err != nil) != tt.wantErr {
				t.Errorf("ResetStats() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.checkStatus()
		})
	}
}
