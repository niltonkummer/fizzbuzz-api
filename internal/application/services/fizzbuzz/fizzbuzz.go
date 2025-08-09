package fizzbuzz

import (
	"fmt"

	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/fizzbuzz"
)

type Service struct {
	fizzbuzz *fizzbuzz.FizzBuzz
	stat     adapters.StatsRepository
}

func NewFizzBuzzService(sts adapters.StatsRepository) *Service {
	service := &Service{
		fizzbuzz: fizzbuzz.NewFizzBuzz(),
		stat:     sts,
	}

	return service
}

func (fb *Service) GenerateFizzBuzz(int1, int2, limit int, str1, str2 string) (string, error) {

	res, err := fb.fizzbuzz.Calculate(int1, int2, limit, str1, str2)
	if err != nil {
		return "", fmt.Errorf("error calculating fizzbuzz: %w", err)
	}

	if err = fb.stat.IncrementRequestCount(int1, int2, limit, str1, str2); err != nil {
		return "", fmt.Errorf("error incrementing request count: %w", err)
	}
	return res, nil
}
