// database/redis.go
package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	// Configuración de la conexión a Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Dirección de Redis
		Password: "",               // Contraseña (si aplica)
		DB:       0,                // Base de datos Redis
	})

	// Verificar la conexión
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error al conectar a Redis: %v", err)
	}

	fmt.Println("Conexión exitosa a Redis")
	return client
}
