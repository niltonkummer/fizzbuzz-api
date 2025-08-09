package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/niltonkummer/fizzbuzz-api/config"
	"github.com/niltonkummer/fizzbuzz-api/internal/adapters/outbound/repository"
	"github.com/niltonkummer/fizzbuzz-api/internal/application"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/services/fizzbuzz"
)

var (
	conf     config.Config
	log      = slog.Default()
	stopTime = 15 * time.Second // 5 minutes

)

func init() {
	conf = config.LoadConfig("./etc/config/")
}

func main() {
	// Setup signal context
	mainCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddress,
		Password: conf.RedisPassword,
	})
	err := client.Ping(client.Context()).Err()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	ongoingCtx, stopGracefully := context.WithCancel(context.Background())
	statsRepo := repository.GetStatsRepository(func() adapters.StatsRepository {
		return repository.NewRedisStatsRepository(client)
	})

	router := application.InitServices(ongoingCtx, statsRepo,
		fizzbuzz.WithCache(func() adapters.CacheFizzbuzz {
			if conf.UseFizzbuzzCache {
				return repository.NewCacheRedis(client)
			}
			return repository.NewCacheFizzbuzzNoOp()
		}()))

	go func() {
		if err := router.Start(conf.HTTPServerHost); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("Failed to start HTTP server: " + err.Error())
		}
	}()

	// Wait for shutdown signal
	<-mainCtx.Done()

	// Shutdown gracefully
	log.InfoContext(mainCtx, "Shutting down server...")

	stopTimeCtx, cancel := context.WithTimeout(context.Background(), stopTime)
	defer cancel()
	if err := router.Shutdown(stopTimeCtx); err != nil {
		log.ErrorContext(mainCtx, "Failed to shutdown server", "error", err)
	}
	err = client.Close()
	if err != nil {
		return
	}
	stopGracefully()

	log.InfoContext(mainCtx, "Server shutdown complete")
}
