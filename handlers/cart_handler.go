// handlers/cart_handler.go
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func AddToCart(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simulación de ID de usuario (puedes obtenerlo de un token JWT o cookies)
		userID := "user:1" // Clave única para el carrito del usuario

		// Decodificar el cuerpo de la solicitud
		var req struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Solicitud inválida", http.StatusBadRequest)
			return
		}

		// Validar los datos
		if req.ProductID <= 0 || req.Quantity <= 0 {
			http.Error(w, "El product_id y quantity deben ser mayores que 0", http.StatusBadRequest)
			return
		}

		// Agregar logs para depurar
		fmt.Printf("Intentando agregar producto: product_id=%d, quantity=%d\n", req.ProductID, req.Quantity)

		// Agregar o actualizar el producto en el carrito
		err = client.HIncrBy(ctx, userID, strconv.Itoa(req.ProductID), int64(req.Quantity)).Err()
		if err != nil {
			fmt.Printf("Error al ejecutar HIncrBy: %v\n", err) // Log de error
			http.Error(w, "Error al agregar al carrito", http.StatusInternalServerError)
			return
		}

		// Configurar tiempo de expiración (opcional)
		client.Expire(ctx, userID, 24*time.Hour) // El carrito caduca después de 24 horas

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Producto agregado al carrito"))
	}
}
