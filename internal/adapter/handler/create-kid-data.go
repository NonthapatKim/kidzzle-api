package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// CreateKidData godoc
// @Summary Create a new kid data
// @Description Create a new kid data with the provided information
// @Tags Kid
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param kidData body domain.CreateKidDataRequest true "Kid Data Request"
// @Success 200 {object} domain.CreateKidDataResponse "Kid data created successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /kids/create [post]
func (h *handler) CreateKidData(c *fiber.Ctx) error {
	var kidData domain.CreateKidDataRequest

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	if err := c.BodyParser(&kidData); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.CreateKidDataRequest{
		AccessToken:          accessToken,
		PregnantId:           kidData.PregnantId,
		KidName:              kidData.KidName,
		KidBirthday:          kidData.KidBirthday,
		KidGender:            kidData.KidGender,
		KidGestationalAge:    kidData.KidGestationalAge,
		KidOxygen:            kidData.KidOxygen,
		KidBirthWeight:       kidData.KidBirthWeight,
		KidBodyLength:        kidData.KidBodyLength,
		KidBloodType:         kidData.KidBloodType,
		KidCongenitalDisease: kidData.KidCongenitalDisease,
	}

	result, err := h.svc.CreateKidData(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
