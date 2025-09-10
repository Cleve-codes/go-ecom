package handlers

import (
	"database/sql"
	"ecommerce-backend/internal/models"
	"ecommerce-backend/internal/services"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db        *sql.DB
	jwtSecret string
}

func NewAuthHandler(db *sql.DB, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

// Register creates a new user account
// @Summary Register a new user
// @Description Registers a new user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "User registration data"
// @Success 201 {object} models.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" || req.FullName == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email, password, and full name are required",
		})
	}

	// Basic email validation
	if !strings.Contains(req.Email, "@") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid email format",
		})
	}

	// Password validation
	if len(req.Password) < 8 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Password must be at least 8 characters long",
		})
	}

	// Check if user already exists
	var existingID string
	err := h.db.QueryRow("SELECT id FROM users WHERE email = $1", req.Email).Scan(&existingID)
	if err == nil {
		return c.Status(409).JSON(fiber.Map{
			"error": "User with this email already exists",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create user
	userID := uuid.New()
	query := `
		INSERT INTO users (id, email, password_hash, full_name, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, email, full_name, role, created_at, updated_at
	`

	var user models.User
	err = h.db.QueryRow(query, userID, req.Email, string(hashedPassword), req.FullName, "customer").
		Scan(&user.ID, &user.Email, &user.FullName, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Generate JWT token
	token, err := services.GenerateJWT(user.ID.String(), user.Role, h.jwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(201).JSON(models.LoginResponse{
		Message: "User registered successfully",
		Token:   token,
		User:    &user,
	})
}

// Login authenticates a user and returns a JWT token
// @Summary Login
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "User login data"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// Get user from database
	var user models.User
	query := `
		SELECT id, email, password_hash, full_name, role, created_at, updated_at
		FROM users WHERE email = $1
	`
	err := h.db.QueryRow(query, req.Email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FullName,
		&user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Generate JWT token
	token, err := services.GenerateJWT(user.ID.String(), user.Role, h.jwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(models.LoginResponse{
		Message: "Login successful",
		Token:   token,
		User:    &user,
	})
}

// GetProfile returns the profile of the authenticated user
// @Summary Get user profile
// @Description Returns the profile of the authenticated user
// @Tags Auth
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} map[string]string
// @Security BearerAuth
// @Router /api/auth/profile [get]
func (h *AuthHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var user models.User
	query := `
		SELECT id, email, full_name, role, created_at, updated_at
		FROM users WHERE id = $1
	`
	err := h.db.QueryRow(query, userID).Scan(
		&user.ID, &user.Email, &user.FullName,
		&user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}
