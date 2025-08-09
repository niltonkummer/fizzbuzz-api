package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/services/fizzbuzz"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/services/stats"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

type Handler struct {
	fizzBuzzService *fizzbuzz.Service
	statsService    *stats.StatsService
}

func NewHandler(fizzBuzz *fizzbuzz.Service, sts *stats.StatsService) *Handler {
	return &Handler{
		fizzBuzzService: fizzBuzz,
		statsService:    sts,
	}
}

// HandleFizzBuzzRequest handles the FizzBuzz request
func (h *Handler) HandleFizzBuzzRequest(ctx echo.Context) error {

	var request model.FizzBuzzRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
			"code":    "invalid_request",
		})
	}

	response, err := h.fizzBuzzService.GenerateFizzBuzz(request.Int1, request.Int2, request.Limit, request.Str1, request.Str2)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process FizzBuzz request: "+err.Error())
	}

	resp := model.FizzBuzzResponse{
		Response: response,
	}

	return ctx.JSON(http.StatusOK, resp)
}

// HandleGetStats handles the statistics request
func (h *Handler) HandleGetStats(ctx echo.Context) error {
	sts, err := h.statsService.GetStats()
	if err != nil {
		if errors.Is(err, model.ErrNoRequestsFound) {
			return ctx.JSON(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get statistics: "+err.Error())
	}

	statsResponse := model.StatsResponse{
		Int1:  sts.Int1,
		Int2:  sts.Int2,
		Limit: sts.Limit,
		Str1:  sts.Str1,
		Str2:  sts.Str2,
		Hits:  sts.Hits,
	}

	return ctx.JSON(http.StatusOK, statsResponse)
}
