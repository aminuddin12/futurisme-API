package server

import (
	"log"

	"futurisme-api/config"
	"futurisme-api/internal/middleware"
	authHttp "futurisme-api/internal/modules/auth/delivery/http"
	authUseCase "futurisme-api/internal/modules/auth/usecase"
	userHttp "futurisme-api/internal/modules/user/delivery/http" // Import User Handler
	userRepo "futurisme-api/internal/modules/user/repository"
	userUseCase "futurisme-api/internal/modules/user/usecase" // Import User UseCase
	"futurisme-api/pkg/database"
	"futurisme-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// RunServer adalah fungsi yang dipanggil oleh CLI command
func RunServer(isDev bool) {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// OVERRIDE Env jika flag --dev aktif
	if isDev {
		cfg.App.Env = "dev"
		cfg.App.Debug = true
		log.Println("🔧 Starting in DEVELOPMENT mode (Flag --dev detected)")
	}

	// 2. Connect Database
	db := database.NewPostgresDatabase(cfg)

	// 3. Setup Dependency Injection (Wiring)
	// Repositories
	userRepository := userRepo.NewUserRepository(db)

	// UseCases
	authUC := authUseCase.NewAuthUseCase(userRepository, cfg)
	userUC := userUseCase.NewUserUseCase(userRepository) // Logic Profile

	// Handlers
	authHandler := authHttp.NewAuthHandler(authUC)
	userHandler := userHttp.NewUserHandler(userUC) // Handler Profile

	// 4. Init Fiber App
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		BodyLimit:    10 * 1024 * 1024,
		ErrorHandler: customErrorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// 5. Routes
	api := app.Group("/api")

	// V1 Group (Dilindungi Layer 1: App Key)
	v1 := api.Group("/v1", middleware.AppLayerAuth(cfg))

	v1.Get("/", func(c *fiber.Ctx) error {
		return response.Success(c, fiber.StatusOK, "Futurisme API v1 is Running", nil)
	})

	// Auth Routes (Public - Tidak butuh JWT)
	authRoutes := v1.Group("/auth")
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)

	// User Routes (Protected - Butuh JWT Layer 2)
	// Middleware.JWTProtected disuntikkan di sini
	userRoutes := v1.Group("/users", middleware.JWTProtected(cfg))
	userRoutes.Get("/profile", userHandler.GetProfile)

	// 6. Start Server
	log.Printf("🚀 Server running on port %s in %s mode", cfg.App.Port, cfg.App.Env)
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return response.Error(c, code, err.Error(), nil)
}
