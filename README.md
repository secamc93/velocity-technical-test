# velocity-technical-test
# Prueba Técnica, Desarrollador GO, Velocity Fullcommerce

¡Bienvenido! Este proyecto consiste en una aplicación Go que integra una API con MySQL y Redis, ofreciendo una interfaz de documentación vía Swagger para explorar fácilmente todos sus endpoints.

## Contenido

- [Requerimientos Previos](#requerimientos-previos)
- [Formas de Ejecución](#formas-de-ejecución)
  - [Opción 1: Docker Compose](#opción-1-docker-compose)
  - [Opción 2: Ejecución desde Código Fuente](#opción-2-ejecución-desde-código-fuente)
- [Variables de Entorno](#variables-de-entorno)
- [Swagger - Documentación de la API](#swagger---documentación-de-la-api)

---

## Requerimientos Previos

- **Go** (1.18 o superior)
- **Docker** y **Docker Compose** (para la opción de ejecución con contenedores)

Asegúrate de tener estas herramientas instaladas antes de continuar.

---

## Formas de Ejecución

### Opción 1: Docker Compose

1. Dirígete a la carpeta `compose/` ubicada en la raíz del proyecto.
2. Ejecuta:
   ```bash
   docker-compose up -d


### Opción 2: Ejecución desde Código Fuente
1. Primero, levanta la base de datos MySQL y Redis con Docker Compose:
        cd compose
        docker-compose up -d mysql redis
2. Regresa a la raíz del proyecto y ejecuta la aplicación Go directamente:
        go run cmd/main.go


## Variables de Entorno
1. Para ejecutar la aplicación localmente (sin el contenedor de la aplicación), debes definir las siguientes variables de entorno:

export DB_HOST=localhost
export DB_USER=root
export DB_PASSWORD=secret
export DB_NAME=mydatabase
export REDIS_HOST=localhost:6379
export SERVER_PORT=60000

## Swagger - Documentación de la API
Una vez la aplicación esté corriendo (ya sea vía Docker Compose o localmente), puedes acceder a la documentación de la API en:

http://localhost:60000/swagger/index.html#/

Aquí encontrarás:

Una lista completa de endpoints.
Parámetros de entrada y salida.
Ejemplos de peticiones y respuestas.