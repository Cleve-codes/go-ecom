package handlers

import (
	"database/sql"
	"ecommerce-backend/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MpesaHandler struct {
	db *sql.DB
}

func NewMpesaHandler(db *sql.DB) *MpesaHandler {
	return &MpesaHandler{db: db}
}

// @Summary Simulate M-Pesa STK Push
// @Description Simulates an M-Pesa payment for an order
// @Tags Mpesa
// @Accept json
// @Produce json
// @Param request body models.MpesaSTKPushRequest true "STK Push data"
// @Success 200 {object} models.MpesaSTKPushResponse
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /api/mpesa/stkpush [post]
func (h *MpesaHandler) InitiateSTKPush(c *fiber.Ctx) error {
	var req models.MpesaSTKPushRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	// Simulate payment processing
	transactionID := uuid.New()
	ref := "MPESA" + transactionID.String()[:8]
	resp := models.MpesaSTKPushResponse{
		TransactionID: transactionID,
		MpesaRef:      ref,
		Status:        "pending",
		Amount:        req.Amount,
		CreatedAt:     time.Now(),
	}
	// Optionally: Save transaction to DB here
	return c.JSON(resp)
}

// @Summary Get M-Pesa transaction status
// @Description Returns the status of a simulated M-Pesa transaction
// @Tags Mpesa
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.MpesaSTKPushResponse
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/mpesa/transaction/{id} [get]
func (h *MpesaHandler) GetTransactionStatus(c *fiber.Ctx) error {
	// Simulate always successful for demo
	id := c.Params("id")
	resp := models.MpesaSTKPushResponse{
		TransactionID: uuid.MustParse(id),
		MpesaRef:      "MPESA" + id[:8],
		Status:        "success",
		Amount:        100.0,
		CreatedAt:     time.Now(),
	}
	return c.JSON(resp)
}

// @Summary M-Pesa payment webhook (simulated)
// @Description Receives payment notifications (simulated)
// @Tags Mpesa
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/mpesa/webhook [post]
func (h *MpesaHandler) Webhook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Webhook received (simulated)"})
}
