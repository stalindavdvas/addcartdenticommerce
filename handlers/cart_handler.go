// handlers/cart_handler.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func AddToCart(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := "user:1" // User key

		// Decode body
		var req struct {
			ProductID int     `json:"product_id"`
			Name      string  `json:"name"`     // Name
			Quantity  int     `json:"quantity"` // Stock
			Price     float64 `json:"price"`    // Price
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Valid data
		if req.ProductID <= 0 || req.Quantity <= 0 || req.Name == "" || req.Price <= 0 {
			http.Error(w, "El product_id, name, quantity y price deben ser válidos", http.StatusBadRequest)
			return
		}

		// JSON to save
		productData := map[string]interface{}{
			"name":     req.Name,
			"quantity": req.Quantity,
			"price":    req.Price,
		}

		// Convert to JSON
		productJSON, err := json.Marshal(productData)
		if err != nil {
			http.Error(w, "Error to process information product", http.StatusInternalServerError)
			return
		}

		// Save on redis
		err = client.HSet(ctx, userID, strconv.Itoa(req.ProductID), productJSON).Err()
		if err != nil {
			http.Error(w, "Failure to add cart", http.StatusInternalServerError)
			return
		}

		// time to expire
		client.Expire(ctx, userID, 24*time.Hour) // Expire 24 hours

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Product Add"))
	}
}
