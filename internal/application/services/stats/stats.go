package stats

import (
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

type StatsService struct {
	repository adapters.StatsRepository
}

// NewStats creates a new Stats instance
func NewStats(repo adapters.StatsRepository) *StatsService {
	return &StatsService{
		repository: repo,
	}
}

// GetStats returns the statistics of the application
func (s *StatsService) GetStats() (stats *model.StatsResult, err error) {
	stats, err = s.repository.GetMostFrequentRequest()
	if err != nil {
		return nil, err
	}

	if stats == nil {
		return nil, model.ErrNoRequestsFound
	}
	return stats, nil
}

// ResetStats resets the statistics data
func (s *StatsService) ResetStats() error {
	return s.repository.ResetStats()
}
