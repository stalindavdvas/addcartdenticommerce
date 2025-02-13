// main.go
package main

import (
	"addcart/database"
	"addcart/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Inicializar la conexión a Redis
	client := database.InitRedis()
	defer client.Close()

	// Crear un nuevo router
	r := mux.NewRouter()

	// Ruta para agregar productos al carrito
	r.HandleFunc("/api/addcart", handlers.AddToCart(client)).Methods("POST")

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://3.229.231.204:3000"},     // Permite solicitudes desde tu frontend
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},  // Métodos HTTP permitidos
		AllowedHeaders: []string{"Content-Type", "Authorization"}, // Encabezados permitidos
	})

	// Usar el middleware de CORS
	handler := corsHandler.Handler(r)

	// Iniciar el servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
