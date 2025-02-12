# Etapa 1: Construcción de la aplicación
FROM golang:1.20-alpine AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos necesarios
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación (sin dependencias del sistema)
RUN CGO_ENABLED=0 GOOS=linux go build -o addcart .

# Etapa 2: Imagen final
FROM alpine:latest

# Instalar dependencias mínimas
RUN apk add --no-cache bash

# Establecer el directorio de trabajo
WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/addcart .

# Exponer el puerto en el que corre la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./addcart"]