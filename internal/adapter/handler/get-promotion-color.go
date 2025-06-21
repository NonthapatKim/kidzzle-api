package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetPromotionColor godoc
// @Summary Get promotion color
// @Description Retrieve the promotion color
// @Tags Promotion
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.GetPromotionColorResponse "Promotion color retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /promotions/color [get]
func (h *handler) GetPromotionColor(c *fiber.Ctx) error {
	result, err := h.svc.GetPromotionColor()
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
