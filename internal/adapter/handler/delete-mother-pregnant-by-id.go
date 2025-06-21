package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// DeleteMotherPregnantById godoc
// @Summary Delete a mother pregnant record by ID
// @Description Delete a mother pregnant record by its ID
// @Tags MotherPregnant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param pregnantId path string true "Pregnant ID"
// @Success 200 {object} domain.DeleteMotherPregnantByIdResponse "Mother pregnant record deleted successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /mother-pregnant/{pregnantId} [delete]
func (h *handler) DeleteMotherPregnantById(c *fiber.Ctx) error {
	pregnantId := c.Params("pregnantId")

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	req := domain.DeleteMotherPregnantByIdRequest{
		AccessToken: accessToken,
		PregnantId:  pregnantId,
	}

	result, err := h.svc.DeleteMotherPregnantById(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
