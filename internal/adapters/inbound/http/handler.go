package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/niltonkummer/fizzbuzz-api/internal/application/adapters"
	"github.com/niltonkummer/fizzbuzz-api/internal/domain/model"
)

type Handler struct {
	fizzBuzzService adapters.FizzBuzzService
	statsService    adapters.StatsService
}

func NewHandler(fizzBuzz adapters.FizzBuzzService, sts adapters.StatsService) *Handler {
	return &Handler{
		fizzBuzzService: fizzBuzz,
		statsService:    sts,
	}
}

// HandleFizzBuzzRequest handles the FizzBuzz request
func (h *Handler) HandleFizzBuzzRequest(ctx echo.Context) error {

	var request model.FizzBuzzRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
			"code":    "invalid_payload",
		})
	}

	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
			"code":    "invalid_request",
		})
	}

	response, err := h.fizzBuzzService.GenerateFizzBuzz(request.Int1, request.Int2, request.Limit, request.Str1, request.Str2)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to generate FizzBuzz response: " + err.Error(),
			"code":    "internal_error",
		})
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

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve statistics: " + err.Error(),
			"code":    "internal_error",
		})
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
