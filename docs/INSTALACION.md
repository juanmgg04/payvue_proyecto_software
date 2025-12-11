# Guía de Instalación

## Requisitos del Sistema

### Software Requerido

| Software | Versión Mínima | Descarga |
|----------|---------------|----------|
| Git | 2.30+ | [git-scm.com](https://git-scm.com) |
| Docker | 20.10+ | [docker.com](https://docker.com) |
| Docker Compose | 2.0+ | Incluido con Docker Desktop |
| Node.js | 18+ | [nodejs.org](https://nodejs.org) (opcional) |
| Go | 1.21+ | [go.dev](https://go.dev) (opcional) |

### Verificar Instalaciones

```bash
# Git
git --version
# Output: git version 2.x.x

# Docker
docker --version
# Output: Docker version 24.x.x

# Docker Compose
docker compose version
# Output: Docker Compose version v2.x.x

# Node.js (opcional)
node --version
# Output: v18.x.x

# Go (opcional)
go version
# Output: go version go1.21.x
```

---

## Instalación Rápida (Docker)

### 1. Clonar el Repositorio

```bash
git clone https://github.com/juanmgg04/payvue_proyecto_software.git
cd payvue_proyecto_software
```

### 2. Iniciar con Docker Compose

```bash
docker compose up -d --build
```

### 3. Verificar que está funcionando

```bash
# Ver contenedores
docker compose ps

# Ver logs
docker compose logs -f

# Probar API
curl http://localhost:8080/health
```

### 4. Acceder a la Aplicación

- **Frontend:** http://localhost:3000
- **Backend API:** http://localhost:8080

---

## Instalación Manual (Sin Docker)

### Backend (Go)

```bash
# Entrar al directorio backend
cd backend

# Descargar dependencias
go mod download

# Crear directorio de datos
mkdir -p data uploads

# Ejecutar servidor
go run cmd/server/main.go
```

El servidor iniciará en `http://localhost:8080`

### Frontend (React)

```bash
# En otra terminal, entrar al directorio frontend
cd frontend

# Instalar dependencias
npm install

# Crear archivo de entorno
echo "REACT_APP_API_URL=http://localhost:8080" > .env

# Iniciar servidor de desarrollo
npm start
```

La aplicación abrirá en `http://localhost:3000`

---

## Variables de Entorno

### Backend

Crear archivo `.env` en `backend/`:

```env
PORT=8080
DATABASE_PATH=./data/payvue.db
ENVIRONMENT=development
```

### Frontend

Crear archivo `.env` en `frontend/`:

```env
REACT_APP_API_URL=http://localhost:8080
```

---

## Estructura de Directorios

```
payvue_proyecto_software/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go          # Punto de entrada
│   ├── pkg/
│   │   ├── domain/              # Lógica de negocio
│   │   ├── repository/          # Acceso a datos
│   │   └── rest/                # Handlers HTTP
│   ├── data/                    # Base de datos (creado automáticamente)
│   ├── uploads/                 # Archivos subidos
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── components/          # Componentes React
│   │   ├── pages/               # Páginas
│   │   └── config/              # Configuración
│   ├── public/
│   ├── package.json
│   └── Dockerfile
├── docs/                        # Documentación
├── docker-compose.yml
└── README.md
```

---

## Comandos Útiles

### Docker

```bash
# Iniciar servicios
docker compose up -d

# Iniciar con rebuild
docker compose up -d --build

# Ver logs en tiempo real
docker compose logs -f

# Ver logs de un servicio específico
docker compose logs -f api
docker compose logs -f frontend

# Detener servicios
docker compose down

# Detener y eliminar volúmenes
docker compose down -v

# Reiniciar un servicio
docker compose restart api

# Ver estado de servicios
docker compose ps

# Entrar a un contenedor
docker compose exec api sh
docker compose exec frontend sh
```

### Backend (Go)

```bash
# Ejecutar en modo desarrollo
go run cmd/server/main.go

# Compilar binario
go build -o bin/server cmd/server/main.go

# Ejecutar binario
./bin/server

# Ejecutar tests
go test ./...

# Ver cobertura de tests
go test -cover ./...
```

### Frontend (React)

```bash
# Instalar dependencias
npm install

# Iniciar desarrollo
npm start

# Compilar para producción
npm run build

# Ejecutar tests
npm test

# Lint
npm run lint
```

---

## Solución de Problemas

### Error: Puerto 8080 en uso

```bash
# Encontrar proceso usando el puerto
lsof -i :8080

# Matar proceso
kill -9 <PID>

# O cambiar puerto en docker-compose.yml
```

### Error: Puerto 3000 en uso

```bash
# Encontrar proceso
lsof -i :3000

# Matar proceso
kill -9 <PID>
```

### Error: Permission denied en uploads

```bash
# Dar permisos al directorio
chmod -R 755 backend/uploads
chmod -R 755 backend/data
```

### Error: Cannot connect to Docker daemon

```bash
# En Linux, agregar usuario al grupo docker
sudo usermod -aG docker $USER

# Reiniciar sesión o ejecutar
newgrp docker
```

### Error: node_modules no encontrado

```bash
# Limpiar e instalar de nuevo
rm -rf node_modules package-lock.json
npm install
```

### Error: Database locked

```bash
# Detener todos los procesos que usan la BD
docker compose down

# Eliminar archivo de BD (perderás datos)
rm -f backend/data/payvue.db

# Reiniciar
docker compose up -d
```

---

## Primer Usuario de Prueba

Después de instalar, puedes registrar un usuario de prueba:

### Usando curl

```bash
# Registrar
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"123456"}'

# Login
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"123456"}'
```

### Usando la interfaz web

1. Ir a http://localhost:3000
2. Click en "¿No tienes cuenta? Crea una"
3. Ingresar email y contraseña
4. Iniciar sesión

---

## Actualización

### Con Docker

```bash
# Obtener últimos cambios
git pull origin main

# Reconstruir y reiniciar
docker compose down
docker compose up -d --build
```

### Sin Docker

```bash
# Backend
cd backend
git pull origin main
go mod download
# Reiniciar servidor

# Frontend
cd frontend
git pull origin main
npm install
# Reiniciar servidor
```

---

## Desinstalación

### Completa (con datos)

```bash
# Detener y eliminar contenedores, redes, volúmenes
docker compose down -v

# Eliminar imágenes
docker rmi payvue-api payvue-frontend

# Eliminar directorio
cd ..
rm -rf payvue_proyecto_software
```

### Solo contenedores (mantener datos)

```bash
docker compose down
```

---

## Soporte

Si encuentras problemas:

1. Revisa los logs: `docker compose logs -f`
2. Verifica las variables de entorno
3. Consulta la documentación
4. Abre un issue en GitHub

