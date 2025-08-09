package repository

import (
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

type StorageType string

func GetStatsRepository(create func() adapters.StatsRepository) adapters.StatsRepository {
	if create != nil {
		return create()
	} else {
		return NewInMemoryStatsRepository(make(map[model.FizzBuzzRequest]int))
	}

}
