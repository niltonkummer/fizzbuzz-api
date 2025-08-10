package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
	"go.uber.org/mock/gomock"
)

func newEchoContext(method, path string, body io.Reader, validator echo.Validator) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = validator
	req := httptest.NewRequest(method, path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return ctx, rec
}

type errorValidator struct{}

func (v *errorValidator) Validate(i interface{}) error {
	return errors.New("validation failed")
}

func TestHandler_HandleFizzBuzzRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validReq := model.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}
	validReqBody, _ := json.Marshal(validReq)

	tests := []struct {
		name           string
		mockService    func(*adapters.MockFizzBuzzService)
		validator      echo.Validator
		body           []byte
		wantStatusCode int
	}{
		{
			name: "success",
			mockService: func(m *adapters.MockFizzBuzzService) {
				m.EXPECT().GenerateFizzBuzz(3, 5, 15, "Fizz", "Buzz").Return("1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz", nil)
			},
			validator:      NewValidator(),
			body:           validReqBody,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "invalid json",
			mockService:    nil,
			validator:      nil,
			body:           []byte("{"),
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "validation error",
			mockService:    nil,
			validator:      &errorValidator{},
			body:           validReqBody,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "service error",
			mockService: func(m *adapters.MockFizzBuzzService) {
				m.EXPECT().GenerateFizzBuzz(3, 5, 15, "Fizz", "Buzz").Return("", errors.New("service fail"))
			},
			validator:      NewValidator(),
			body:           validReqBody,
			wantStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFizzBuzz := adapters.NewMockFizzBuzzService(ctrl)
			if tt.mockService != nil {
				tt.mockService(mockFizzBuzz)
			}

			h := NewHandler(mockFizzBuzz, nil)
			ctx, rec := newEchoContext(http.MethodPost, "/fizzbuzz", bytes.NewReader(tt.body), tt.validator)
			_ = h.HandleFizzBuzzRequest(ctx)

			if rec.Code != tt.wantStatusCode {
				t.Errorf("expected %d, got %d", tt.wantStatusCode, rec.Code)
			}
		})
	}
}

func TestHandler_HandleGetStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statResult := &model.StatsResult{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz", Hits: 10}

	tests := []struct {
		name           string
		mockService    func(*adapters.MockStatsService)
		wantStatusCode int
	}{
		{
			name: "success",
			mockService: func(m *adapters.MockStatsService) {
				m.EXPECT().GetStats().Return(statResult, nil)
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "no stats found",
			mockService: func(m *adapters.MockStatsService) {
				m.EXPECT().GetStats().Return(nil, model.ErrNoRequestsFound)
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "service error",
			mockService: func(m *adapters.MockStatsService) {
				m.EXPECT().GetStats().Return(nil, errors.New("db fail"))
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStats := adapters.NewMockStatsService(ctrl)
			if tt.mockService != nil {
				tt.mockService(mockStats)
			}
			h := NewHandler(nil, mockStats)
			ctx, rec := newEchoContext(http.MethodGet, "/stats", nil, nil)
			_ = h.HandleGetStats(ctx)

			if rec.Code != tt.wantStatusCode {
				t.Errorf("expected %d, got %d", tt.wantStatusCode, rec.Code)
			}
		})
	}
}
