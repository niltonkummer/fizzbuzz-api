package application

import (
	"context"

	httpIn "github.com/niltonkummer/fizzbuzz-api/internal/adapters/inbound/http"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/services/fizzbuzz"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/services/stats"
)

func InitServices(ctx context.Context, repo adapters.StatsRepository, opts ...fizzbuzz.Option) *httpIn.Router {
	fizzBuzzService := fizzbuzz.NewFizzBuzzService(repo, opts...)
	statsService := stats.NewStats(repo)

	handler := httpIn.NewHandler(fizzBuzzService, statsService)

	router := httpIn.NewRouter(ctx)
	router.RegisterRoutes(handler)
	return router
}
