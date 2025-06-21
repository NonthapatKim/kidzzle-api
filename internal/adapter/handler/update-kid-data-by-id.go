package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UpdateKidDataById godoc
// @Summary Update kid data by ID
// @Description Update the data of a kid by their ID
// @Tags Kid
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param kidId path string true "Kid ID"
// @Param kidData body domain.UpdateKidDataByIdRequest true "Kid Data Update Request"
// @Success 200 {object} domain.UpdateKidDataByIdResponse "Kid data updated successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /kids/{kidId} [put]
func (h *handler) UpdateKidDataById(c *fiber.Ctx) error {
	var kidData domain.UpdateKidDataByIdRequest

	kidId := c.Params("kidId")

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	if err := c.BodyParser(&kidData); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UpdateKidDataByIdRequest{
		AccessToken:          accessToken,
		KidId:                kidId,
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

	result, err := h.svc.UpdateKidDataById(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
