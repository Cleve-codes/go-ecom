package handlers

import (
	"database/sql"
	"ecommerce-backend/internal/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	db         *sql.DB
	uploadPath string
}

func NewProductHandler(db *sql.DB, uploadPath string) *ProductHandler {
	return &ProductHandler{db: db, uploadPath: uploadPath}
}

// @Summary Get all products
// @Tags Products
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/products [get]
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	rows, err := h.db.Query(`SELECT id, name, description, price, stock, category, image_url, created_at, updated_at FROM products ORDER BY created_at DESC`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Category, &p.ImageURL, &p.CreatedAt, &p.UpdatedAt); err == nil {
			products = append(products, p)
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"products": products,
		"total":    len(products),
		"page":     1,
		"limit":    len(products),
	})
}

// @Summary Get a product by ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Product
	err := h.db.QueryRow(`SELECT id, name, description, price, stock, category, image_url, created_at, updated_at FROM products WHERE id = $1`, id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Category, &p.ImageURL, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.Status(201).JSON(p)
}

// @Summary Create a new product
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Param product formData models.Product true "Product data"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /api/products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	price := c.FormValue("price")
	stock := c.FormValue("stock")
	category := c.FormValue("category")

	// Debug: log received form values
	fmt.Printf("Received: name=%q, price=%q, stock=%q, category=%q\n", name, price, stock, category)
	// Validate required fields
	if name == "" || price == "" || stock == "" || category == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name, price, stock, and category are required"})
	}

	// Parse price and stock
	var priceVal float64
	var stockVal int
	n1, err1 := fmt.Sscanf(price, "%f", &priceVal)
	n2, err2 := fmt.Sscanf(stock, "%d", &stockVal)
	if n1 != 1 || n2 != 1 || err1 != nil || err2 != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Price and stock must be valid numbers"})
	}

	// Handle image upload
	imageUrl := ""
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		filePath := h.uploadPath + "/" + file.Filename
		if saveErr := c.SaveFile(file, filePath); saveErr == nil {
			imageUrl = filePath
		}
	}

	var p models.Product
	err = h.db.QueryRow(
		`INSERT INTO products (name, description, price, stock, category, image_url) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at`,
		name, description, priceVal, stockVal, category, imageUrl,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create product"})
	}
	// Populate the rest of the product struct
	p.Name = name
	p.Description = description
	p.Price = priceVal
	p.Stock = stockVal
	p.Category = category
	p.ImageURL = imageUrl
	return c.Status(201).JSON(p)
}

// @Summary Update a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body models.Product true "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.Product
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if p.Name == "" || p.Price <= 0 || p.Stock < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Name, positive price, and non-negative stock are required"})
	}
	_, err := h.db.Exec(
		`UPDATE products SET name=$1, description=$2, price=$3, stock=$4, category=$5, image_url=$6, updated_at=NOW() WHERE id=$7`,
		p.Name, p.Description, p.Price, p.Stock, p.Category, p.ImageURL, id,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update product"})
	}
	// Return updated product with 201 status
	prodResp := h.GetProduct(c)
	c.Status(201)
	return prodResp
}

// @Summary Delete a product
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Router /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := h.db.Exec(`DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete product"})
	}
	return c.SendStatus(204)
}

// @Summary Upload product image
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Product ID"
// @Param image formData file true "Product image"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /api/products/{id}/image [post]
func (h *ProductHandler) UploadProductImage(c *fiber.Ctx) error {
	id := c.Params("id")
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Image file is required"})
	}
	filePath := h.uploadPath + "/" + file.Filename
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save image"})
	}
	_, err = h.db.Exec(`UPDATE products SET image_url = $1, updated_at = NOW() WHERE id = $2`, filePath, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update product image"})
	}
	// Return updated product with 201 status
	prodResp := h.GetProduct(c)
	c.Status(201)
	return prodResp
}
