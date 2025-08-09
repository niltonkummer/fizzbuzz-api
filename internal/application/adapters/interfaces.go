package adapters

import "github.com/niltonkummer/fizzbuzz-api/internal/domain/model"

//go:generate mockgen -source=interfaces.go -destination=mock_interfaces.go -package=adapters

type StatsRepository interface {
	// GetMostFrequentRequest returns the most frequent request parameters and their hit count
	GetMostFrequentRequest() (stats *model.StatsResult, err error)
	// IncrementRequestCount increments the count for a specific request parameters
	IncrementRequestCount(int1, int2, limit int, str1, str2 string) error
	// ResetStats resets the statistics data
	ResetStats() error
}

type CacheFizzbuzz interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
