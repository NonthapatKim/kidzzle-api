package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetAssessmentResultDataById godoc
// @Summary Get assessment result data by ID
// @Description Get assessment result data for a given kid, age range, and assessment type
// @Tags Assessment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param kidId path string true "Kid ID"
// @Param ageRageId path string true "Age Range ID"
// @Param assessmentTypeId path string true "Assessment Type ID"
// @Success 200 {object} domain.GetAssessmentResultDataByIdResponse "Assessment result data retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /assessments/{kidId}/{ageRageId}/{assessmentTypeId}/result [get]
func (h *handler) GetAssessmentResultDataById(c *fiber.Ctx) error {
	kidId := c.Params("kidId")
	ageRageId := c.Params("ageRageId")
	assessmentTypeId := c.Params("assessmentTypeId")

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	req := domain.GetAssessmentResultDataByIdRequest{
		AccessToken:      accessToken,
		KidId:            kidId,
		AgeRageId:        ageRageId,
		AssessmentTypeId: assessmentTypeId,
	}

	result, err := h.svc.GetAssessmentResultDataById(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
