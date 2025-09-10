package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"ecommerce-backend/internal/models"
)

func SeedShoes(db *sql.DB) {
	shoes := []models.Product{
		{
			Name:        "Nike Air Max 270",
			Description: "Nike's iconic Air Max 270 with a large air unit for comfort.",
			Price:       149.99,
			ImageURL:    "https://fakestoreapi.com/img/airmax270.png",
			Category:    "shoes",
			Stock:       50,
			CreatedAt:   time.Now(),
		},
		{
			Name:        "Adidas Ultraboost 21",
			Description: "Responsive running shoes with Boost cushioning.",
			Price:       179.99,
			ImageURL:    "https://fakestoreapi.com/img/ultraboost21.png",
			Category:    "shoes",
			Stock:       40,
			CreatedAt:   time.Now(),
		},
		{
			Name:        "Converse Chuck Taylor All Star",
			Description: "Classic canvas sneakers for everyday style.",
			Price:       59.99,
			ImageURL:    "https://fakestoreapi.com/img/chucktaylor.png",
			Category:    "shoes",
			Stock:       100,
			CreatedAt:   time.Now(),
		},
		{
			Name:        "Vans Old Skool",
			Description: "Skate-inspired low-top shoes with a retro look.",
			Price:       64.99,
			ImageURL:    "https://fakestoreapi.com/img/vansoldskool.png",
			Category:    "shoes",
			Stock:       80,
			CreatedAt:   time.Now(),
		},
		{
			Name:        "Puma RS-X",
			Description: "Chunky sneakers with bold color blocking.",
			Price:       109.99,
			ImageURL:    "https://fakestoreapi.com/img/pumarsx.png",
			Category:    "shoes",
			Stock:       60,
			CreatedAt:   time.Now(),
		},
		{
			Name:        "New Balance 574",
			Description: "Retro running shoes with ENCAP midsole cushioning.",
			Price:       89.99,
			ImageURL:    "https://fakestoreapi.com/img/nb574.png",
			Category:    "shoes",
			Stock:       70,
			CreatedAt:   time.Now(),
		},
	}

	for _, shoe := range shoes {
		_, err := db.ExecContext(context.Background(),
			`INSERT INTO products (name, description, price, image_url, category, stock, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (name) DO NOTHING`,
			shoe.Name, shoe.Description, shoe.Price, shoe.ImageURL, shoe.Category, shoe.Stock, shoe.CreatedAt,
		)
		if err != nil {
			log.Printf("Failed to seed shoe: %s, error: %v", shoe.Name, err)
		}
	}

	log.Println("Shoe products seeded successfully.")
}
