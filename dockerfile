# Usar una imagen base oficial de Go
FROM golang:1.23.2-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum y descargar las dependencias
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copiar el código fuente del proyecto
COPY . .

# Compilar la aplicación desde la carpeta cmd
RUN go build -o main ./cmd

# Exponer el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
