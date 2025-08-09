package repository

import (
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

var _ adapters.StatsRepository = (*InMemoryStatsRepository)(nil)

// InMemoryStatsRepository is an in-memory implementation of StatsRepository
type InMemoryStatsRepository struct {
	stats map[model.FizzBuzzRequest]int
}

// NewInMemoryStatsRepository creates a new InMemoryStatsRepository instance
func NewInMemoryStatsRepository(start map[model.FizzBuzzRequest]int) *InMemoryStatsRepository {
	return &InMemoryStatsRepository{
		stats: start,
	}
}

// GetMostFrequentRequest returns the most frequent request parameters and their hit count
func (r *InMemoryStatsRepository) GetMostFrequentRequest() (stats *model.StatsResult, err error) {

	var maxHits int
	var key model.FizzBuzzRequest
	for request, hits := range r.stats {
		if hits > maxHits {
			key = request
			maxHits = hits
		}
	}
	if maxHits > 0 {
		return &model.StatsResult{
			Int1:  key.Int1,
			Int2:  key.Int2,
			Limit: key.Limit,
			Str1:  key.Str1,
			Str2:  key.Str2,
			Hits:  maxHits,
		}, nil
	}

	return stats, nil
}

// IncrementRequestCount increments the count for a specific request parameters
func (r *InMemoryStatsRepository) IncrementRequestCount(int1, int2, limit int, str1, str2 string) error {
	request := model.FizzBuzzRequest{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}
	r.stats[request]++
	return nil
}

// ResetStats resets the statistics data
func (r *InMemoryStatsRepository) ResetStats() error {
	r.stats = make(map[model.FizzBuzzRequest]int)
	return nil
}
