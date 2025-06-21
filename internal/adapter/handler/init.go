package handler

import (
	"github.com/NonthapatKim/kidzzle-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	// Assessment
	CreateAssessmentResult(c *fiber.Ctx) error
	GetAssessmentAgeRange(c *fiber.Ctx) error
	GetAssessmentQuestion(c *fiber.Ctx) error
	GetAssessmentResultById(c *fiber.Ctx) error
	GetAssessmentResultDataById(c *fiber.Ctx) error
	GetAssessmentTraining(c *fiber.Ctx) error

	// Mother Pregnant
	CreateMotherPregnant(c *fiber.Ctx) error
	DeleteMotherPregnantById(c *fiber.Ctx) error
	GetMotherPregnantById(c *fiber.Ctx) error
	UpdateMotherPregnantById(c *fiber.Ctx) error

	// Kid
	CreateKidData(c *fiber.Ctx) error
	DeleteKidDataById(c *fiber.Ctx) error
	GetKidDataById(c *fiber.Ctx) error
	UpdateKidDataById(c *fiber.Ctx) error

	// Promotion
	GetPromotionBodyPart(c *fiber.Ctx) error
	GetPromotionColor(c *fiber.Ctx) error

	// User
	UserAuthenticate(c *fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
	UserLoginBySocial(c *fiber.Ctx) error
	UserRegister(c *fiber.Ctx) error
	UserRequestResetPassword(c *fiber.Ctx) error
	UserResetPassword(c *fiber.Ctx) error
}

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}
