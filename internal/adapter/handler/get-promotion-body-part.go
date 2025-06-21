package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetPromotionBodyPart godoc
// @Summary Get promotion body part
// @Description Retrieve the promotion body part details
// @Tags Promotion
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.GetPromotionBodyPartResponse "Promotion body part retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /promotions/body-part [get]
func (h *handler) GetPromotionBodyPart(c *fiber.Ctx) error {
	result, err := h.svc.GetPromotionBodyPart()
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
