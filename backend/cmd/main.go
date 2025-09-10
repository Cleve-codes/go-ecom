package main

import (
	"log"
	"os"

	_ "ecommerce-backend/docs"
	"ecommerce-backend/internal/config"
	"ecommerce-backend/internal/database"
	"ecommerce-backend/internal/handlers"
	"ecommerce-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db := database.NewConnection(cfg.GetDBConnectionString())
	defer db.Close()

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll(cfg.UploadPath, 0755); err != nil {
		log.Fatal("Failed to create uploads directory:", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
		BodyLimit: 10 * 1024 * 1024, // 10MB
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:5173", // Add your frontend URLs
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Global debug middleware: log all requests and bodies
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("[DEBUG] %s %s", c.Method(), c.OriginalURL())
		if c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPut {
			// Try to read form and body
			if c.Is("multipart/form-data") {
				form, err := c.MultipartForm()
				if err == nil && form != nil {
					log.Printf("[DEBUG] Multipart form: %+v", form.Value)
					log.Printf("[DEBUG] Multipart files: %+v", form.File)
				} else {
					log.Printf("[DEBUG] Multipart form parse error: %v", err)
				}
			} else {
				log.Printf("[DEBUG] Body: %s", string(c.Body()))
			}
		}
		return c.Next()
	})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db.DB, cfg.JWTSecret)
	userHandler := handlers.NewUserHandler(db.DB)
	productHandler := handlers.NewProductHandler(db.DB, cfg.UploadPath)
	orderHandler := handlers.NewOrderHandler(db.DB)
	mpesaHandler := handlers.NewMpesaHandler(db.DB)

	// API routes
	api := app.Group("/api")
	// Endpoint to list all products
	api.Get("/products", productHandler.GetProducts)
	api.Get("/products/:id", productHandler.GetProduct)
	api.Post("/products", productHandler.CreateProduct)
	api.Delete("/products/:id", productHandler.DeleteProduct)
	// Endpoint to list and create orders
	api.Get("/orders", middleware.AuthRequired(cfg.JWTSecret), orderHandler.GetUserOrders)
	api.Post("/orders", middleware.AuthRequired(cfg.JWTSecret), orderHandler.CreateOrder)
	api.Get("/admin/orders", middleware.AuthRequired(cfg.JWTSecret), middleware.AdminRequired(), orderHandler.GetAllOrders)

	// M-Pesa payment simulation routes
	api.Post("/mpesa/stkpush", middleware.AuthRequired(cfg.JWTSecret), mpesaHandler.InitiateSTKPush)
	api.Get("/mpesa/transaction/:id", middleware.AuthRequired(cfg.JWTSecret), mpesaHandler.GetTransactionStatus)
	api.Post("/mpesa/webhook", mpesaHandler.Webhook)

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"message": "E-commerce API is running",
		})
	})

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Get("/profile", middleware.AuthRequired(cfg.JWTSecret), authHandler.GetProfile)

	// Admin user management routes
	adminUsers := api.Group("/admin/users", middleware.AuthRequired(cfg.JWTSecret), middleware.AdminRequired())
	adminUsers.Get("/", userHandler.GetAllUsers)
	adminUsers.Get(":id", userHandler.GetUserByID)
	adminUsers.Put(":id", userHandler.UpdateUser)
	adminUsers.Delete(":id", userHandler.DeleteUser)

	// Serve uploaded files
	app.Static("/uploads", cfg.UploadPath)

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running")
	})

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Printf("Environment: %s", cfg.Env)
	log.Fatal(app.Listen(":" + cfg.Port))
}
