package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// CreateMotherPregnant godoc
// @Summary Create a new mother pregnant record
// @Description Create a new mother pregnant record with the given details
// @Tags MotherPregnant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param motherPregnant body domain.CreateMotherPregnantRequest true "Mother Pregnant Request"
// @Success 200 {object} domain.CreateMotherPregnantResponse "Mother pregnant record created successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /mother-pregnants/create [post]
func (h *handler) CreateMotherPregnant(c *fiber.Ctx) error {
	var motherPregnant domain.CreateMotherPregnantRequest

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	if err := c.BodyParser(&motherPregnant); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.CreateMotherPregnantRequest{
		AccessToken:               accessToken,
		MotherName:                motherPregnant.MotherName,
		MotherBirthday:            motherPregnant.MotherBirthday,
		PregnantCongenitalDisease: motherPregnant.PregnantCongenitalDisease,
		PregnantComplications:     motherPregnant.PregnantComplications,
		PregnantDrugHistory:       motherPregnant.PregnantDrugHistory,
	}

	result, err := h.svc.CreateMotherPregnant(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
