package router

import (
	"fmt"
	"os"

	_ "github.com/NonthapatKim/kidzzle-api/docs"
	"github.com/NonthapatKim/kidzzle-api/internal/adapter/handler"
	"github.com/NonthapatKim/kidzzle-api/internal/adapter/handler/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	app *fiber.App
}

const serviceBaseURL = "/api"

func NewRouter(h handler.Handler) (*Router, error) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // For Develop Only
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	basePath := app.Group(serviceBaseURL)
	basePathV1 := basePath.Group("/v1").Use(middleware.LoggerMiddleware())

	basePathV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("สวัสดี !")
	})

	// Swagger UI
	basePathV1.Get("/swagger/*", swagger.HandlerDefault)

	assessment := basePathV1.Group("/assessments")
	{
		assessment.Post("/result", middleware.Authorization(), h.CreateAssessmentResult)
		assessment.Get("/:kidId/:assessmentQuestionId/result", middleware.Authorization(), h.GetAssessmentResultById)
		assessment.Get("/:kidId/:ageRageId/:assessmentTypeId/result", middleware.Authorization(), h.GetAssessmentResultDataById)
		assessment.Get("/:assessmentType/question", h.GetAssessmentQuestion)
		assessment.Get("/:assessmentType/age-range", h.GetAssessmentAgeRange)
		assessment.Get("/:assessmentType/training", h.GetAssessmentTraining)
	}

	motherPregnant := basePathV1.Group("/mother-pregnants")
	{
		motherPregnant.Get("/", middleware.Authorization(), h.GetMotherPregnantById)
		motherPregnant.Post("/create", middleware.Authorization(), h.CreateMotherPregnant)
		motherPregnant.Put("/:pregnantId", middleware.Authorization(), h.UpdateMotherPregnantById)
		motherPregnant.Delete("/:pregnantId", middleware.Authorization(), h.DeleteMotherPregnantById)
	}

	kid := basePathV1.Group("/kids")
	{
		kid.Get("/:pregnantId", middleware.Authorization(), h.GetKidDataById)
		kid.Post("/create", middleware.Authorization(), h.CreateKidData)
		kid.Put("/:kidId", middleware.Authorization(), h.UpdateKidDataById)
		kid.Delete("/:kidId", middleware.Authorization(), h.DeleteKidDataById)
	}

	promotion := basePathV1.Group("/promotions")
	{
		promotion.Get("/body-part", h.GetPromotionBodyPart)
		promotion.Get("/color", h.GetPromotionColor)
	}

	user := basePathV1.Group("/users")
	{
		user.Get("/authenticate", middleware.Authorization(), h.UserAuthenticate)
		user.Post("/login", h.UserLogin)
		user.Post("/login/social", h.UserLoginBySocial)
		user.Post("/register", h.UserRegister)
		user.Post("/request-reset-password", h.UserRequestResetPassword)
		user.Put("/reset-password", middleware.Authorization(), h.UserResetPassword)
	}

	return &Router{app: app}, nil
}

func (r *Router) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port", port)
	return r.app.Listen(":" + port)
}
