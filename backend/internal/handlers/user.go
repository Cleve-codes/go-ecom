package handlers

import (
	"database/sql"
	"ecommerce-backend/internal/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

// @Summary List all users
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Security BearerAuth
// @Router /api/admin/users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	rows, err := h.db.Query(`SELECT id, email, full_name, role, created_at, updated_at FROM users WHERE disabled = false`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.FullName, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan user"})
		}
		users = append(users, user)
	}
	return c.Status(201).JSON(users)
}

// @Summary Get user by ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/admin/users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	query := `SELECT id, email, full_name, role, created_at, updated_at FROM users WHERE id = $1 AND disabled = false`
	err := h.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.FullName, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	return c.Status(201).JSON(user)
}

// @Summary Update user info/role
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UpdateUserRequest true "User data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /api/admin/users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Build update query dynamically
	setClauses := []string{}
	args := []interface{}{}
	argIdx := 1
	if req.FullName != nil {
		setClauses = append(setClauses, "full_name = $"+string(rune(argIdx)))
		args = append(args, *req.FullName)
		argIdx++
	}
	if req.Role != nil {
		setClauses = append(setClauses, "role = $"+string(rune(argIdx)))
		args = append(args, *req.Role)
		argIdx++
	}
	if req.Disabled != nil {
		setClauses = append(setClauses, "disabled = $"+string(rune(argIdx)))
		args = append(args, *req.Disabled)
		argIdx++
	}
	if len(setClauses) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No fields to update"})
	}
	setQuery := "SET " + strings.Join(setClauses, ", ")
	args = append(args, id)
	query := "UPDATE users " + setQuery + ", updated_at = NOW() WHERE id = $" + string(rune(argIdx)) + " RETURNING id, email, full_name, role, created_at, updated_at"

	var user models.User
	err := h.db.QueryRow(query, args...).Scan(&user.ID, &user.Email, &user.FullName, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	return c.Status(201).JSON(user)
}

// @Summary Delete (soft) user
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /api/admin/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	query := `UPDATE users SET disabled = true, updated_at = NOW() WHERE id = $1`
	res, err := h.db.Exec(query, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to check result"})
	}
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.SendStatus(204)
}
