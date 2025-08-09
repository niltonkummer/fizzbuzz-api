package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"strings"

	"github.com/cucumber/godog"
	"github.com/go-redis/redis/v8"
	"github.com/niltonkummer/fizzbuzz-api/config"
	"github.com/niltonkummer/fizzbuzz-api/internal/adapters/inbound/http"
	"github.com/niltonkummer/fizzbuzz-api/internal/adapters/outbound/repository"
	"github.com/niltonkummer/fizzbuzz-api/internal/application"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

type godogsResponseCtxKey struct{}

type apiFeature struct {
	router *http.Router
	repo   adapters.StatsRepository
	body   []byte
	resp   *httptest.ResponseRecorder
}

type response struct {
	status int
	body   any
}

func (a *apiFeature) resetResponse(sc *godog.Scenario) {

	a.resp = httptest.NewRecorder()

	if strings.Contains(sc.Name, "reset stats") {
		if err := a.repo.ResetStats(); err != nil {
			panic(fmt.Sprintf("failed to reset stats: %v", err))
		}
	}
}

func (a *apiFeature) iSendrequestTo(ctx context.Context, method, route string) (context.Context, error) {
	return a.iSendrequestToRequestBody(ctx, method, route, nil)
}

func (a *apiFeature) iSendrequestToRequestBody(ctx context.Context, method, route string, payloadDoc *godog.DocString) (context.Context, error) {
	var reqBody []byte

	if payloadDoc != nil {
		payloadMap := model.FizzBuzzRequest{}
		err := json.Unmarshal([]byte(payloadDoc.Content), &payloadMap)
		if err != nil {
			panic(err)
		}

		reqBody, _ = json.Marshal(payloadMap)
	}
	app := a.router.GetApp()
	req := httptest.NewRequest(method, route, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	var resp any

	if method == "POST" && route == "/fizzbuzz" {
		err := a.router.GetHandler().HandleFizzBuzzRequest(app.NewContext(req, a.resp))
		if err != nil {
			return ctx, fmt.Errorf("error handling request: %w", err)
		}
		var fizzBuzzResponse model.FizzBuzzResponse
		a.body, _ = io.ReadAll(a.resp.Body)
		json.NewDecoder(bytes.NewBuffer(a.body)).Decode(&fizzBuzzResponse)
		resp = fizzBuzzResponse.Response
	}
	if method == "GET" && route == "/stats" {
		err := a.router.GetHandler().HandleGetStats(app.NewContext(req, a.resp))
		if err != nil {
			return ctx, fmt.Errorf("error handling request: %w", err)
		}
		var statsResponse model.StatsResponse
		a.body, _ = io.ReadAll(a.resp.Body)
		json.NewDecoder(bytes.NewBuffer(a.body)).Decode(&statsResponse)
		resp = map[string]any{
			"int1":  statsResponse.Int1,
			"int2":  statsResponse.Int2,
			"limit": statsResponse.Limit,
			"str1":  statsResponse.Str1,
			"str2":  statsResponse.Str2,
			"hits":  statsResponse.Hits,
		}

	}
	actual := response{
		status: a.resp.Code,
		body:   resp,
	}

	return context.WithValue(ctx, godogsResponseCtxKey{}, actual), nil
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	var expected, actual interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.NewDecoder(bytes.NewBuffer(a.body)).Decode(&actual); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	config := config.LoadConfig("../etc/config/")
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: config.RedisAddress,
		})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		fmt.Println("Redis connection failed, using in-memory stats repository")
		panic(err)
	}
	repo := repository.NewRedisStatsRepository(redisClient)

	api := &apiFeature{
		router: application.InitServices(context.Background(), repo, fizzbuzz.WithCache(func() repository.CacheFizzbuzz {
			if config.UseFizzbuzzCache {
				return repository.NewCacheRedis(redisClient)
			}
			return repository.NewCacheFizzbuzzNoOp()
		}())),
		repo: repo,
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})

	ctx.Step(`^I send "(POST|PUT|DELETE)" request to "([^"]*)" with payload:$`, api.iSendrequestToRequestBody)
	ctx.Step(`^I send "(GET)" request to "([^"]*)":$`, api.iSendrequestTo)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response payload should match json:$`, api.theResponseShouldMatchJSON)
}
