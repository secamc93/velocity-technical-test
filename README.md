# velocity-technical-test

Prueba Técnica, Desarrollador GO, Velocity Fullcommerce
Este proyecto es una aplicación que incluye una API con acceso a una base de datos MySQL y un servicio Redis. Además, cuenta con documentación a través de Swagger para facilitar la exploración de los endpoints disponibles.

Formas de Levantar la Aplicación
Opción 1: Usar Docker Compose
Asegúrese de tener instalado Docker y Docker Compose.

En la raíz del proyecto, ubique la carpeta compose/.

Dentro de compose/ encontrará el archivo docker-compose.yaml.

Ejecute el siguiente comando para levantar la base de datos, Redis y la aplicación:

bash
Copiar código
docker-compose up -d
Este comando iniciará todos los servicios en contenedores. Una vez finalizado, la aplicación estará disponible y en ejecución.

Opción 2: Ejecutar la Aplicación Desde el Código Fuente
Asegúrese de tener instalados Go (versión 1.18 o superior).

Inicie la base de datos MySQL y Redis utilizando Docker Compose:

bash
Copiar código
cd compose
docker-compose up -d mysql redis
Esto inicia solo la base de datos y Redis, dejando fuera el contenedor de la aplicación.

En la raíz del proyecto, ejecute la aplicación directamente con Go:

bash
Copiar código
go run cmd/main.go
De esta forma, la aplicación se levantará localmente, utilizando la base de datos y Redis previamente iniciados con Docker.

Variables de Entorno
Si desea correr la aplicación localmente sin el contenedor de la aplicación, deberá proporcionar ciertas variables de entorno. Por ejemplo:

bash
Copiar código
export DB_HOST=localhost
export DB_USER=root
export DB_PASSWORD=secret
export DB_NAME=mydatabase
export REDIS_HOST=localhost:6379
export SERVER_PORT=60000
Ajuste estos valores según la configuración de su entorno.

Acceso a la Documentación Swagger
Una vez que la aplicación esté en línea (ya sea vía Docker Compose o ejecutando go run), podrá acceder a la documentación Swagger en:

http://localhost:60000/swagger/index.html#/

Aquí encontrará información detallada sobre los endpoints disponibles, los parámetros de entrada/salida y ejemplos de peticiones y respuestas.