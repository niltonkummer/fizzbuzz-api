package tests

import (
	"context"
	config2 "niltonkummer/fizz-buzz/config"
	"testing"

	"github.com/cucumber/godog"
	"github.com/go-redis/redis/v8"
)

func resetRedis() {
	conf := config2.LoadConfig("../etc/config/")
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
