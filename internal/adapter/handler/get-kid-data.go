package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetKidDataById godoc
// @Summary Get kid data by ID
// @Description Retrieve kid data using the provided access token
// @Tags Kid
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param PregnantId path string true "Pregnant ID"
// @Param accessToken header string true "Access Token"
// @Success 200 {object} domain.GetKidDataResponse "Kid data retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /kids/{pregnantId} [get]
func (h *handler) GetKidDataById(c *fiber.Ctx) error {
	accessToken := c.Locals("accessToken").(string)

	pregnantId := c.Params("pregnantId")

	req := domain.GetKidDataRequest{
		AccessToken: accessToken,
		PregnantId:  pregnantId,
	}

	result, err := h.svc.GetKidData(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
