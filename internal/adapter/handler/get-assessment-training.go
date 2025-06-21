package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetAssessmentTraining godoc
// @Summary Get assessment training methods
// @Description Get training methods for a given assessment type
// @Tags Assessment
// @Accept  json
// @Produce  json
// @Param assessmentType path string true "Assessment Type ID"
// @Success 200 {object} domain.GetAssessmentTrainingMethodsResponse "Training methods retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /assessments/{assessmentType}/training [get]
func (h *handler) GetAssessmentTraining(c *fiber.Ctx) error {
	assessmentType := c.Params("assessmentType")

	req := domain.GetAssessmentTrainingMethodsRequest{
		AssessmentTypeId: assessmentType,
	}

	result, err := h.svc.GetAssessmentTraining(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
