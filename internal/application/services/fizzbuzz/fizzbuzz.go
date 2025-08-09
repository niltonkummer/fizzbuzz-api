package fizzbuzz

import (
	"fmt"

	"github.com/niltonkummer/fizzbuzz-api/internal/adapters/outbound/repository"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/fizzbuzz"
)

type Option func(*Service)

type Service struct {
	fizzbuzz *fizzbuzz.FizzBuzz
	stat     adapters.StatsRepository
	cache    adapters.CacheFizzbuzz
}

// WithCache allows setting a cache for the FizzBuzz service
func WithCache(cache adapters.CacheFizzbuzz) Option {
	return func(s *Service) {
		s.cache = cache
	}
}

func NewFizzBuzzService(sts adapters.StatsRepository, opts ...Option) *Service {
	service := &Service{
		fizzbuzz: fizzbuzz.NewFizzBuzz(),
		stat:     sts,
		cache:    repository.NewCacheFizzbuzzNoOp(), // Default to no-op cache
	}
	for _, opt := range opts {
		opt(service)
	}

	return service
}

func (fb *Service) GenerateFizzBuzz(int1, int2, limit int, str1, str2 string) (string, error) {

	res, err := fb.calculateFizzBuzzOrGetFromCache(int1, int2, limit, str1, str2)
	if err != nil {
		return "", fmt.Errorf("error calculating or getting from cache: %w", err)
	}

	if err = fb.stat.IncrementRequestCount(int1, int2, limit, str1, str2); err != nil {
		return "", fmt.Errorf("error incrementing request count: %w", err)
	}
	return res, nil
}

func (fb *Service) calculateFizzBuzzOrGetFromCache(int1, int2, limit int, str1, str2 string) (string, error) {
	res, err := fb.cache.Get(fmt.Sprintf("%d,%d,%d,%s,%s", int1, int2, limit, str1, str2))
	if res != "" {
		if err != nil {
			return "", fmt.Errorf("error getting from cache: %w", err)
		}
		return res, nil
	}

	res, err = fb.fizzbuzz.Calculate(int1, int2, limit, str1, str2)
	if err != nil {
		return "", fmt.Errorf("error calculating fizzbuzz: %w", err)
	}

	if err = fb.cache.Set(fmt.Sprintf("%d,%d,%d,%s,%s", int1, int2, limit, str1, str2), res); err != nil {
		return "", fmt.Errorf("error setting cache: %w", err)
	}
	return res, nil
}
