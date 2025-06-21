package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetAssessmentResultById godoc
// @Summary Get an assessment result by ID
// @Description Get an assessment result for a given kid and assessment question
// @Tags Assessment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param kidId path string true "Kid ID"
// @Param assessmentQuestionId path string true "Assessment Question ID"
// @Success 200 {object} domain.GetAssessmentResultByIdResponse "Assessment result retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /assessments/{kidId}/{assessmentQuestionId}/result [get]
func (h *handler) GetAssessmentResultById(c *fiber.Ctx) error {
	kidId := c.Params("kidId")
	assessmentQuestionId := c.Params("assessmentQuestionId")

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	req := domain.GetAssessmentResultByIdRequest{
		AccessToken:          accessToken,
		KidId:                kidId,
		AssessmentQuestionId: assessmentQuestionId,
	}

	result, err := h.svc.GetAssessmentResultById(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
