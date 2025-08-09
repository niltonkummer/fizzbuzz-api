package tests

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/go-redis/redis/v8"
	"github.com/niltonkummer/fizzbuzz-api/config"
)

func resetRedis() {
	conf := config.LoadConfig("../etc/config/")
	// Reset Redis database
	client := redis.NewClient(&redis.Options{
		Addr: conf.RedisAddress,
	})
	err := client.FlushDB(context.Background()).Err()
	if err != nil {
		panic(err)
	}

}

func TestFeatures(t *testing.T) {
	resetRedis()
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run features tests")
	}
}
