package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// GetAssessmentAgeRange godoc
// @Summary Get assessment age range
// @Description Get the age range for a given assessment type
// @Tags Assessment
// @Accept  json
// @Produce  json
// @Param assessmentType path string true "Assessment Type ID"
// @Success 200 {object} domain.GetAssessmentAgeRangeResponse "Age range retrieved successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /assessments/{assessmentType}/age-range/ [get]
func (h *handler) GetAssessmentAgeRange(c *fiber.Ctx) error {
	assessmentType := c.Params("assessmentType")

	req := domain.GetAssessmentAgeRangeRequest{
		AssessmentTypeId: assessmentType,
	}

	result, err := h.svc.GetAssessmentAgeRange(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
