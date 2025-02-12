// main.go
package main

import (
	"addcart/database"
	"addcart/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar la conexión a Redis
	client := database.InitRedis()
	defer client.Close()

	// Crear un nuevo router
	r := mux.NewRouter()

	// Ruta para agregar productos al carrito
	r.HandleFunc("/api/cart", handlers.AddToCart(client)).Methods("POST")

	// Iniciar el servidor
	log.Println("Servidor iniciado en http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
