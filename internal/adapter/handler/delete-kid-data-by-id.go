package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// DeleteKidDataById godoc
// @Summary Delete kid data by ID
// @Description Delete kid data by given kid ID
// @Tags Kid
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param kidId path string true "Kid ID"
// @Success 200 {object} domain.DeleteKidDataByIdResponse "Kid data deleted successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /kids/{kidId} [delete]
func (h *handler) DeleteKidDataById(c *fiber.Ctx) error {
	kidId := c.Params("kidId")

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	req := domain.DeleteKidDataByIdRequest{
		AccessToken: accessToken,
		KidId:       kidId,
	}

	result, err := h.svc.DeleteKidDataById(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
