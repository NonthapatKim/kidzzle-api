package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UpdateMotherPregnantById godoc
// @Summary Update mother pregnant information by ID
// @Description Update the information of a mother pregnant by the given ID
// @Tags MotherPregnant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param pregnantId path string true "Pregnant ID"
// @Param motherPregnant body domain.UpdateMotherPregnantByIdRequest true "Mother Pregnant Update Request"
// @Success 200 {object} domain.UpdateMotherPregnantByIdResponse "Mother pregnant information updated successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /mother-pregnants/{pregnantId} [put]
func (h *handler) UpdateMotherPregnantById(c *fiber.Ctx) error {
	var motherPregnant domain.UpdateMotherPregnantByIdRequest

	pregnantId := c.Params("pregnantId")

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	if err := c.BodyParser(&motherPregnant); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UpdateMotherPregnantByIdRequest{
		AccessToken:               accessToken,
		PregnantId:                pregnantId,
		MotherName:                motherPregnant.MotherName,
		MotherBirthday:            motherPregnant.MotherBirthday,
		PregnantCongenitalDisease: motherPregnant.PregnantCongenitalDisease,
		PregnantComplications:     motherPregnant.PregnantComplications,
		PregnantDrugHistory:       motherPregnant.PregnantDrugHistory,
	}

	result, err := h.svc.UpdateMotherPregnantById(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
