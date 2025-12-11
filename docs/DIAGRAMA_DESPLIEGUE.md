# Diagrama de Despliegue

## Diagrama UML de Despliegue

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Diagrama de Despliegue UML - PayVue

node "Cliente" {
  node "Navegador Web" as Browser {
    artifact "React SPA" as SPA
  }
}

cloud "Internet" as Net

node "Vercel CDN" {
  node "Edge Server" as Edge {
    artifact "index.html" as HTML
    artifact "main.js" as JS
    artifact "main.css" as CSS
  }
}

node "Render Cloud" {
  node "Docker Container" as Container {
    artifact "Go Binary" as Binary
    artifact "SQLite DB" as DB
    folder "uploads" as Uploads
  }
}

Browser --> Net : HTTPS:443
Net --> Edge : HTTPS:443
Edge --> Container : HTTPS REST API
Binary --> DB : SQL
Binary --> Uploads : File I/O

note right of Edge
  CDN Global
  Archivos estáticos
  Cache distribuido
end note

note right of Container
  Container Docker
  Alpine Linux
  Port 8080
end note

@enduml
```

---

## Diagrama de Arquitectura de Red

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Arquitectura de Red - PayVue

actor "Usuario" as U
participant "Navegador" as B
participant "Vercel CDN\n(Edge Network)" as V
participant "Render\n(Load Balancer)" as R
participant "Container\n:8080" as C
database "SQLite" as DB

U -> B: Abrir aplicación
B -> V: GET https://payvue.vercel.app
V --> B: HTML + JS + CSS

B -> B: Render React App

B -> R: POST /auth/login\nHTTPS:443
R -> C: Forward :8080
C -> DB: SELECT user
DB --> C: User data
C --> R: JSON Response
R --> B: 200 OK + user_id

B -> R: GET /finances/debt\n+ X-User-ID header
R -> C: Forward :8080
C -> DB: SELECT debts WHERE user_id=?
DB --> C: Debt records
C --> R: JSON Array
R --> B: 200 OK

@enduml
```

---

## Diagrama de Componentes de Despliegue

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Componentes de Despliegue

package "GitHub Repository" {
  [main branch] as Main
  [frontend branch] as FBranch
  [backend branch] as BBranch
}

package "CI/CD Pipeline" {
  [Vercel Build] as VBuild
  [Render Build] as RBuild
}

package "Vercel Production" {
  [Static Files\nReact Build] as Static
  [CDN\nEdge Network] as CDN
}

package "Render Production" {
  [Docker Image\nGo Binary] as Docker
  [Web Service\n:8080] as Service
}

Main --> FBranch
Main --> BBranch

FBranch --> VBuild : Webhook
BBranch --> RBuild : Webhook

VBuild --> Static : npm run build
Static --> CDN : Deploy

RBuild --> Docker : docker build
Docker --> Service : Deploy

note right of CDN
  Global Edge
  ~200 locations
  Auto SSL
end note

note right of Service
  Oregon Region
  512MB RAM
  0.1 vCPU
end note

@enduml
```

---

## Diagrama de Flujo de CI/CD

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Flujo de CI/CD - PayVue

|Desarrollador|
start
:Escribir código;
:git commit;
:git push;

|GitHub|
:Recibir push;
:Trigger webhooks;

|Vercel|
if (branch == frontend?) then (yes)
  :npm install;
  :npm run build;
  :Deploy to CDN;
  :Generar URL preview;
endif

|Render|
if (branch == backend?) then (yes)
  :docker build;
  :Run health check;
  :Deploy container;
  :Update DNS;
endif

|Producción|
:Servicio actualizado;
:Notificar éxito;

stop

@enduml
```

---

## Diagrama de Infraestructura

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Infraestructura de Producción

rectangle "Usuario Final" #LightBlue {
  actor User
  rectangle "Browser" as Browser
}

cloud "Vercel Global Network" #LightGreen {
  rectangle "Edge US East" as E1
  rectangle "Edge Europe" as E2
  rectangle "Edge Asia" as EN
  artifact "React Build" as Build
}

cloud "Render Cloud" #LightYellow {
  rectangle "Oregon Region" as Oregon {
    rectangle "Load Balancer" as LB
    rectangle "Docker Container" as DC {
      rectangle "Go API :8080" as API
      storage "SQLite DB" as DB
      folder "Uploads" as UL
    }
  }
}

User --> Browser
Browser --> E1
Browser --> E2
E1 --> Build
E2 --> Build

Browser --> LB : API calls
LB --> DC
API --> DB
API --> UL

@enduml
```

---

## Especificaciones de Nodos

### Nodo: Vercel (Frontend)

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Especificación - Nodo Vercel

rectangle "Vercel Production" {
  rectangle "build/" {
    file "index.html" as index
    rectangle "static/js/" {
      file "main.[hash].js" as mainjs
      file "chunk.[hash].js" as chunk
    }
    rectangle "static/css/" {
      file "main.[hash].css" as maincss
    }
  }
}

note right of index
  Entry point
  de la aplicación
end note

note bottom of "Vercel Production"
  Build: npm run build
  Framework: Create React App
  Node: 18.x
  SSL: Let's Encrypt (auto)
end note

@enduml
```

### Nodo: Render (Backend)

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Especificación - Nodo Render

rectangle "Docker Container" {
  rectangle "Go Binary" as Go
  rectangle "/app/data/" {
    storage "payvue.db" as DB
  }
  rectangle "/app/uploads/" {
    file "receipts/*" as receipts
  }
}

Go --> DB : SQL queries
Go --> receipts : File I/O

note right of Go
  HTTP Server
  Port 8080
end note

note bottom of "Docker Container"
  Base Image: alpine:latest
  Runtime: Docker
  Region: Oregon (US West)
  RAM: 512 MB
  CPU: 0.1 vCPU (shared)
  Plan: Free Tier
end note

@enduml
```

---

## Diagrama de Secuencia de Despliegue

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Secuencia de Despliegue

actor "Dev" as D
participant "GitHub" as G
participant "Vercel" as V
participant "Render" as R

== Despliegue Frontend ==

D -> G: git push frontend
G -> V: Webhook trigger
activate V
V -> V: npm ci
V -> V: npm run build
V -> V: Deploy to CDN
V --> G: Status: Success
deactivate V
G --> D: ✅ Frontend deployed

== Despliegue Backend ==

D -> G: git push backend
G -> R: Webhook trigger
activate R
R -> R: docker build
R -> R: Health check
R -> R: Replace container
R --> G: Status: Success
deactivate R
G --> D: ✅ Backend deployed

@enduml
```

---

## Diagrama de Seguridad

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Capas de Seguridad

package "Transporte" #LightBlue {
  [TLS 1.3\nCifrado en tránsito] as TLS
}

package "API Gateway" #LightGreen {
  [CORS\nControl de origen] as CORS
  [Rate Limiting\n(futuro)] as Rate
}

package "Autenticación" #LightYellow {
  [X-User-ID Header] as Auth
  [Session Management] as Session
}

package "Validación" #LightPink {
  [Input Validation\ngo-playground/validator] as Valid
  [Sanitization] as Sanit
}

package "Base de Datos" #LightCyan {
  [Prepared Statements\nSQL Injection Prevention] as SQL
  [User Data Isolation] as Isolation
}

TLS --> CORS
CORS --> Auth
Auth --> Valid
Valid --> SQL

@enduml
```

---

## URLs de Producción

| Servicio | URL | Descripción |
|----------|-----|-------------|
| Frontend | https://payvue.vercel.app | Aplicación React |
| Backend | https://payvue-api.onrender.com | API REST |
| Health Check | https://payvue-api.onrender.com/health | Estado del servicio |
| GitHub | https://github.com/juanmgg04/payvue_proyecto_software | Código fuente |

---

## Resumen de Especificaciones

| Componente | Tecnología | Especificación |
|------------|------------|----------------|
| **Frontend** | React 18 | Vercel CDN, SSL automático |
| **Backend** | Go 1.21 | Docker Alpine, 512MB RAM |
| **Base de Datos** | SQLite | Embebida en container |
| **CI/CD** | GitHub Webhooks | Despliegue automático |
| **SSL** | Let's Encrypt | Certificados automáticos |
| **CDN** | Vercel Edge | ~200 ubicaciones globales |
