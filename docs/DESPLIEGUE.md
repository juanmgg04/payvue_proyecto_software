# Proceso de Despliegue

## Descripción General

PayVue utiliza una estrategia de despliegue separada para frontend y backend:

- **Frontend**: Vercel (hosting de aplicaciones React)
- **Backend**: Render (hosting de contenedores Docker)
- **Control de Versiones**: GitHub (ramas separadas)

---

## Arquitectura de Despliegue

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Arquitectura de Despliegue

actor "Desarrollador" as Dev
node "GitHub" {
  [main] as Main
  [frontend] as FB
  [backend] as BB
}

cloud "Vercel" {
  [CDN\nReact Build] as CDN
}

cloud "Render" {
  [Docker\nGo API] as Docker
  database "SQLite" as DB
}

actor "Usuario" as User
node "Navegador" as Browser

Dev --> Main : push
Main --> FB : branch
Main --> BB : branch

FB --> CDN : CI/CD
BB --> Docker : CI/CD

Docker --> DB

User --> Browser
Browser --> CDN : HTTPS
CDN --> Docker : REST API

@enduml
```

---

## Requisitos Previos

### Cuentas Necesarias
- ✅ Cuenta en [GitHub](https://github.com)
- ✅ Cuenta en [Vercel](https://vercel.com)
- ✅ Cuenta en [Render](https://render.com)

### Herramientas Locales
- Git instalado
- Docker y Docker Compose
- Node.js 18+
- Go 1.21+

---

## Paso 1: Preparar el Repositorio

### Estructura de Ramas

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Estructura de Ramas Git

object "main" as main {
  Código completo
  Documentación
}

object "frontend" as front {
  Solo React
  vercel.json
  package.json
}

object "backend" as back {
  Solo Go
  Dockerfile
  render.yaml
}

main --> front : crear rama
main --> back : crear rama

@enduml
```

### Crear Rama Frontend

```bash
# Desde main
git checkout main
git pull origin main

# Crear rama frontend
git checkout -b frontend

# Copiar archivos del frontend a la raíz
cp -r frontend/* .
rm -rf backend/

# Commit y push
git add .
git commit -m "feat: setup frontend for Vercel"
git push origin frontend
```

### Crear Rama Backend

```bash
# Volver a main
git checkout main

# Crear rama backend
git checkout -b backend

# Copiar archivos del backend a la raíz
cp -r backend/* .
rm -rf frontend/

# Commit y push
git add .
git commit -m "feat: setup backend for Render"
git push origin backend
```

---

## Paso 2: Desplegar Backend en Render

### Diagrama del Proceso

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Despliegue en Render

start
:Ir a Render Dashboard;
:Click "New +" > "Web Service";
:Conectar repositorio GitHub;
:Seleccionar rama "backend";
:Configurar servicio;

partition "Configuración" {
  :Name: payvue-api;
  :Region: Oregon;
  :Runtime: Docker;
  :Plan: Free;
}

partition "Variables de Entorno" {
  :PORT=8080;
  :DATABASE_PATH=/app/data/payvue.db;
  :ENVIRONMENT=production;
}

:Click "Create Web Service";
:Esperar build (~5 min);
:Verificar health check;
stop

@enduml
```

### Configuración del Servicio

| Campo | Valor |
|-------|-------|
| **Name** | payvue-api |
| **Region** | Oregon (US West) |
| **Branch** | backend |
| **Runtime** | Docker |
| **Plan** | Free |

### Variables de Entorno

```
PORT=8080
ENVIRONMENT=production
DATABASE_PATH=/app/data/payvue.db
```

### Archivo render.yaml

```yaml
services:
  - type: web
    name: payvue-api
    runtime: docker
    dockerfilePath: ./Dockerfile
    dockerContext: .
    region: oregon
    plan: free
    healthCheckPath: /health
    envVars:
      - key: PORT
        value: "8080"
      - key: ENVIRONMENT
        value: production
      - key: DATABASE_PATH
        value: /app/data/payvue.db
```

### Verificar Despliegue

```bash
# Health check
curl https://payvue-api.onrender.com/health
# Respuesta: OK - PayVue API Server
```

---

## Paso 3: Desplegar Frontend en Vercel

### Diagrama del Proceso

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Despliegue en Vercel

start
:Ir a Vercel Dashboard;
:Click "Add New..." > "Project";
:Importar repositorio GitHub;
:Seleccionar rama "frontend";

partition "Configuración" {
  :Framework: Create React App;
  :Root Directory: ./;
  :Build Command: npm run build;
  :Output Directory: build;
}

partition "Variables de Entorno" {
  :REACT_APP_API_URL=
  https://payvue-api.onrender.com;
}

:Click "Deploy";
:Esperar build (~2 min);
:Verificar aplicación;
stop

@enduml
```

### Configuración del Proyecto

| Campo | Valor |
|-------|-------|
| **Framework Preset** | Create React App |
| **Root Directory** | ./ |
| **Build Command** | npm run build |
| **Output Directory** | build |

### Variables de Entorno

```
REACT_APP_API_URL=https://payvue-api.onrender.com
```

### Archivo vercel.json

```json
{
  "rewrites": [
    { "source": "/(.*)", "destination": "/index.html" }
  ],
  "headers": [
    {
      "source": "/(.*)",
      "headers": [
        { "key": "X-Content-Type-Options", "value": "nosniff" },
        { "key": "X-Frame-Options", "value": "DENY" }
      ]
    }
  ]
}
```

---

## Flujo Completo de CI/CD

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Flujo CI/CD Completo

|Desarrollador|
start
:Escribir código;
:Probar localmente;
:git add .;
:git commit -m "feat: ...";
:git push origin [branch];

|GitHub|
:Recibir push;
fork
  :Webhook a Vercel;
fork again
  :Webhook a Render;
end fork

|Vercel|
if (branch == frontend?) then (yes)
  :Clonar repo;
  :npm ci;
  :npm run build;
  if (Build exitoso?) then (yes)
    :Deploy a CDN;
    :Invalidar cache;
  else (no)
    :Notificar error;
    stop
  endif
endif

|Render|
if (branch == backend?) then (yes)
  :Clonar repo;
  :docker build;
  if (Build exitoso?) then (yes)
    :Run container;
    :Health check;
    if (Health OK?) then (yes)
      :Swap containers;
      :Actualizar DNS;
    else (no)
      :Rollback;
      stop
    endif
  else (no)
    :Notificar error;
    stop
  endif
endif

|Producción|
:Servicios actualizados;
:URLs activas;

|Desarrollador|
:Recibir notificación;
:Verificar despliegue;
stop

@enduml
```

---

## Comandos Útiles

### Desarrollo Local

```bash
# Levantar todo con Docker
docker compose up -d --build

# Ver logs
docker compose logs -f

# Detener
docker compose down
```

### Verificación de Producción

```bash
# Frontend
curl -I https://payvue.vercel.app

# Backend health
curl https://payvue-api.onrender.com/health

# Test login
curl -X POST https://payvue-api.onrender.com/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"123456"}'
```

---

## Troubleshooting

### Backend tarda en responder

**Causa**: Render Free Tier pone el servicio en "sleep" después de inactividad.

**Solución**: Primera request tarda ~30s (cold start). El servicio se mantiene activo después.

### Frontend no conecta con backend

**Causa**: Variable de entorno no configurada.

**Solución**:
1. Verificar `REACT_APP_API_URL` en Vercel
2. Hacer redeploy

### Error de CORS

**Causa**: Backend no permite origen del frontend.

**Solución**: Verificar configuración CORS en `main.go`:
```go
router.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
}))
```

---

## Checklist de Despliegue

### Backend (Render)
- [ ] Rama `backend` actualizada
- [ ] Dockerfile funcional
- [ ] render.yaml configurado
- [ ] Variables de entorno
- [ ] Health check pasando

### Frontend (Vercel)
- [ ] Rama `frontend` actualizada
- [ ] vercel.json configurado
- [ ] REACT_APP_API_URL configurada
- [ ] Build exitoso

### Integración
- [ ] Login/Register funcional
- [ ] CRUD de datos funcional
- [ ] Separación de datos por usuario

---

## URLs de Producción

| Servicio | URL |
|----------|-----|
| **Frontend** | https://payvue.vercel.app |
| **Backend API** | https://payvue-api.onrender.com |
| **Health Check** | https://payvue-api.onrender.com/health |
| **GitHub** | https://github.com/juanmgg04/payvue_proyecto_software |
