package repository

import (
	"fmt"
	"strconv"
	"strings"

	redis "github.com/go-redis/redis/v8"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

var _ adapters.StatsRepository = (*RedisStatsRepository)(nil)

const (
	// RedisKeyStats is the key used to store statistics in Redis
	RedisKeyStats = "fizzbuzz:stats"
)

type RedisStatsRepository struct {
	client *redis.Client
}

func NewRedisStatsRepository(redis *redis.Client) *RedisStatsRepository {
	return &RedisStatsRepository{
		client: redis,
	}
}

// GetMostFrequentRequest returns the most frequent request parameters and their hit count
func (r *RedisStatsRepository) GetMostFrequentRequest() (stats *model.StatsResult, err error) {
	ctx := r.client.Context()
	cmd := r.client.ZRevRangeWithScores(ctx, RedisKeyStats, 0, 0)
	if cmd.Err() != nil {
		return stats, cmd.Err()
	}
	if len(cmd.Val()) == 0 {
		return stats, nil
	}

	mostFrequent := cmd.Val()[0]
	parts := strings.Split(mostFrequent.Member.(string), ",")

	int1, err := strconv.Atoi(parts[0])
	if err != nil {
		return stats, fmt.Errorf("failed to parse int1: %w", err)
	}
	int2, err := strconv.Atoi(parts[1])
	if err != nil {
		return stats, fmt.Errorf("failed to parse int2: %w", err)
	}
	limit, err := strconv.Atoi(parts[2])
	if err != nil {
		return stats, fmt.Errorf("failed to parse limit: %w", err)
	}
	str1 := parts[3]
	str2 := parts[4]

	var hitsFloat float64
	hitsFloat = mostFrequent.Score
	hits := int(hitsFloat)
	return &model.StatsResult{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
		Hits:  hits,
	}, nil
}

// IncrementRequestCount increments the count for a specific request parameters
func (r *RedisStatsRepository) IncrementRequestCount(int1, int2, limit int, str1, str2 string) error {
	ctx := r.client.Context()
	key := fmt.Sprintf("%d,%d,%d,%s,%s", int1, int2, limit, str1, str2)
	cmd := r.client.ZIncrBy(ctx, RedisKeyStats, 1, key)
	return cmd.Err()
}

// ResetStats resets the statistics data
func (r *RedisStatsRepository) ResetStats() error {
	ctx := r.client.Context()
	cmd := r.client.Del(ctx, RedisKeyStats)
	return cmd.Err()
}
