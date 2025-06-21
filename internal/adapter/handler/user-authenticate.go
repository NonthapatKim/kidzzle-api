package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain"
	"github.com/NonthapatKim/kidzzle-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UserAuthenticate godoc
// @Summary Authenticate user from token
// @Description Authenticates the user based on the provided user token
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token" default(Bearer <token>)
// @Success 200 {object} domain.UserAuthenticateResponse
// @Failure 400 {object} domain.ErrorResponseExample400
// @Failure 500 {object} domain.ErrorResponseExample500
// @Router /users/authenticate [get]
func (h *handler) UserAuthenticate(c *fiber.Ctx) error {
	accessToken := c.Locals("accessToken").(string)

	req := domain.UserAuthenticateRequest{
		AccessToken: accessToken,
	}

	result, err := h.svc.UserAuthenticate(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
