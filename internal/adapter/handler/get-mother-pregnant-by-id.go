package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetMotherPregnantById godoc
// @Summary Get mother pregnant by ID
// @Description Get mother pregnant information by ID
// @Tags MotherPregnant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param id path string true "Mother ID"
// @Success 200 {object} domain.GetMotherPregnantByIdResponse "Mother pregnant information retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /mother-pregnants/ [get]
func (h *handler) GetMotherPregnantById(c *fiber.Ctx) error {
	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	req := domain.GetMotherPregnantByIdRequest{
		AccessToken: accessToken,
	}

	result, err := h.svc.GetMotherPregnantById(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
