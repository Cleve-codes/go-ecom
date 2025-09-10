package handlers

import (
	"database/sql"
	"ecommerce-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderHandler struct {
	db *sql.DB
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{db: db}
}

// @Summary Create a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body models.OrderRequest true "Order data"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Router /api/orders [post]
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var req models.OrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if len(req.Items) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Order must have at least one item"})
	}

	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tx, err := h.db.Begin()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to start transaction"})
	}
	defer tx.Rollback()

	var total float64
	orderID := ""

	// Insert order
	err = tx.QueryRow(
		`INSERT INTO orders (user_id, status, total_amount, shipping_address, phone_number) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		userID, "pending", 0.0, req.ShippingAddress, req.PhoneNumber,
	).Scan(&orderID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create order"})
	}

	// Insert order items and calculate total
	for _, item := range req.Items {
		// Use unit_price from the request (frontend/cart)
		_, err := tx.Exec(
			`INSERT INTO order_items (order_id, product_id, quantity, unit_price) VALUES ($1, $2, $3, $4)`,
			orderID, item.ProductID, item.Quantity, item.UnitPrice,
		)
		if err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{"error": "Failed to add order item"})
		}
		total += item.UnitPrice * float64(item.Quantity)
	}

	// Update order total
	_, err = tx.Exec(`UPDATE orders SET total_amount = $1 WHERE id = $2`, total, orderID)
	if err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update order total"})
	}

	if err := tx.Commit(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to commit order"})
	}

	// Fetch and return the created order (simplified)
	var order models.Order
	err = h.db.QueryRow(
		`SELECT id, user_id, status, total_amount, created_at, updated_at FROM orders WHERE id = $1`, orderID,
	).Scan(&order.ID, &order.UserID, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch created order"})
	}

	// Fetch order items
	rows, err := h.db.Query(`SELECT id, order_id, product_id, quantity, unit_price FROM order_items WHERE order_id = $1`, orderID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var item models.OrderItem
			if err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price); err == nil {
				order.Items = append(order.Items, item)
			}
		}
	}

	return c.Status(201).JSON(order)
}

// @Summary Get all orders for a user
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/orders [get]
func (h *OrderHandler) GetUserOrders(c *fiber.Ctx) error {
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok || userIDStr == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	rows, err := h.db.Query(`SELECT id, user_id, status, total_amount, created_at, updated_at FROM orders WHERE user_id = $1 ORDER BY created_at DESC`, userUUID)
	if err != nil {
		// Log the actual SQL error for debugging
		println("[GetUserOrders SQL ERROR]", err.Error())
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch orders", "details": err.Error()})
	}
	defer rows.Close()
	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt); err == nil {
			// Fetch order items
			order.Items = []models.OrderItem{}
			itemRows, err := h.db.Query(`SELECT oi.id, oi.order_id, oi.product_id, oi.quantity, oi.unit_price, p.name FROM order_items oi JOIN products p ON oi.product_id = p.id WHERE oi.order_id = $1`, order.ID)
			if err == nil {
				for itemRows.Next() {
					var item models.OrderItem
					var productName string
					if err := itemRows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price, &productName); err == nil {
						item.Product = &models.Product{Name: productName}
						order.Items = append(order.Items, item)
					}
				}
				itemRows.Close()
			}
			if order.Items == nil {
				order.Items = []models.OrderItem{}
			}
			orders = append(orders, order)
		}
	}
	return c.Status(200).JSON(orders)
}

// @Summary Get an order by ID
// @Tags Orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} map[string]string
// @Router /api/orders/{id} [get]
func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	err := h.db.QueryRow(`SELECT id, user_id, status, total_amount, created_at, updated_at FROM orders WHERE id = $1`, id).
		Scan(&order.ID, &order.UserID, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}
	order.Items = []models.OrderItem{}
	itemRows, err := h.db.Query(`SELECT oi.id, oi.order_id, oi.product_id, oi.quantity, oi.unit_price, p.name FROM order_items oi JOIN products p ON oi.product_id = p.id WHERE oi.order_id = $1`, id)
	if err == nil {
		for itemRows.Next() {
			var item models.OrderItem
			var productName string
			if err := itemRows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price, &productName); err == nil {
				item.Product = &models.Product{Name: productName}
				order.Items = append(order.Items, item)
			}
		}
		itemRows.Close()
	}
	if order.Items == nil {
		order.Items = []models.OrderItem{}
	}
	return c.Status(200).JSON(order)
}

// @Summary Get all orders (admin)
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/admin/orders [get]
func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	rows, err := h.db.Query(`SELECT id, user_id, status, total_amount, created_at, updated_at FROM orders ORDER BY created_at DESC`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch orders"})
	}
	defer rows.Close()
	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.Status, &order.TotalAmount, &order.CreatedAt, &order.UpdatedAt); err == nil {
			// Fetch order items
			itemRows, err := h.db.Query(`SELECT oi.id, oi.order_id, oi.product_id, oi.quantity, oi.unit_price, p.name FROM order_items oi JOIN products p ON oi.product_id = p.id WHERE oi.order_id = $1`, order.ID)
			if err == nil {
				for itemRows.Next() {
					var item models.OrderItem
					var productName string
					if err := itemRows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price, &productName); err == nil {
						item.Product = &models.Product{Name: productName}
						order.Items = append(order.Items, item)
					}
				}
				itemRows.Close()
			}
			orders = append(orders, order)
		}
	}
	return c.Status(201).JSON(orders)
}

// @Summary Update order status (admin)
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param status body models.UpdateOrderStatusRequest true "Order status"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]string
// @Router /api/admin/orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var req models.UpdateOrderStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if req.Status == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Status is required"})
	}
	_, err := h.db.Exec(`UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2`, req.Status, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update order status"})
	}
	// Return updated order with 201 status
	orderResp := h.GetOrder(c)
	c.Status(201)
	return orderResp
}
