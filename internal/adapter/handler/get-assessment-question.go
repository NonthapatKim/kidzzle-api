package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetAssessmentQuestion godoc
// @Summary Get assessment question
// @Description Get assessment question by assessment type
// @Tags Assessment
// @Accept  json
// @Produce  json
// @Param assessmentType path string true "Assessment Type"
// @Success 200 {object} domain.GetAssessmentQuestionResponse "Assessment question retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /assessments/{assessmentType}/question/ [get]
func (h *handler) GetAssessmentQuestion(c *fiber.Ctx) error {
	assessmentType := c.Params("assessmentType")

	req := domain.GetAssessmentQuestionRequest{
		AssessmentTypeId: assessmentType,
	}

	result, err := h.svc.GetAssessmentQuestion(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
