package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// CreateAssessmentResult godoc
// @Summary Create a new assessment result
// @Description Create a new assessment result for a given kid and assessment question
// @Tags Assessment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param accessToken header string true "Access Token"
// @Param assessment body domain.CreateAssessmentResultRequest true "Assessment Result Request"
// @Success 200 {object} domain.CreateAssessmentResultResponse "Assessment result created successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /assessments/result [post]
func (h *handler) CreateAssessmentResult(c *fiber.Ctx) error {
	var assessment domain.CreateAssessmentResultRequest

	accessToken, ok := c.Locals("accessToken").(string)
	if !ok || accessToken == "" {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid access token", nil)
	}

	if err := c.BodyParser(&assessment); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.CreateAssessmentResultRequest{
		AccessToken:          accessToken,
		KidId:                assessment.KidId,
		AssessmentQuestionId: assessment.AssessmentQuestionId,
		IsPassed:             assessment.IsPassed,
	}

	result, err := h.svc.CreateAssessmentResult(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
