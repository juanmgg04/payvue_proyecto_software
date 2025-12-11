# Arquitectura del Software

## Descripción General

PayVue implementa una **arquitectura de capas** (Layered Architecture) combinada con principios de **Clean Architecture** para el backend, y una arquitectura basada en **componentes** para el frontend.

---

## Diagrama de Arquitectura General

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Arquitectura General de PayVue

rectangle "Cliente" #LightBlue {
  rectangle "Navegador Web" as Browser
}

rectangle "Frontend - React" #LightGreen {
  rectangle "Componentes UI" as UI
  rectangle "Pages" as Pages
  rectangle "Config/API" as Config
}

rectangle "Backend - Go" #LightYellow {
  rectangle "HTTP Handlers" as Handlers
  rectangle "Services" as Services
  rectangle "Repositories" as Repos
}

rectangle "Base de Datos" #LightGray {
  storage "SQLite" as DB
}

Browser --> UI
UI --> Pages
Pages --> Config
Config --> Handlers : HTTP/REST
Handlers --> Services
Services --> Repos
Repos --> DB

@enduml
```

---

## Diagrama de Capas del Backend

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Arquitectura de Capas - Backend Go

rectangle "Capa de Presentación" #LightBlue {
  rectangle "REST Handlers" as Handlers {
  }
  rectangle "Entities/DTOs" as DTOs {
  }
}

rectangle "Capa de Dominio/Negocio" #LightGreen {
  rectangle "Service Interface" as IService {
  }
  rectangle "Domain Entities" as Domain {
  }
}

rectangle "Capa de Datos" #LightYellow {
  rectangle "Repository Interface" as IRepo {
  }
  rectangle "SQLite Repository" as SQLiteRepo {
  }
}

rectangle "Base de Datos" #LightGray {
  storage "SQLite" as DB
}

Handlers --> IService : usa
DTOs --> Domain : mapea
IService --> IRepo : usa
SQLiteRepo ..|> IRepo : implementa
SQLiteRepo --> DB : queries

note right of Handlers
  HandleCreate()
  HandleGet()
  HandleUpdate()
  HandleDelete()
end note

note right of IService
  Create()
  GetAll()
  GetByID()
  Update()
  Delete()
end note

note right of IRepo
  Create()
  FindAll()
  FindByID()
  Update()
  Delete()
end note

@enduml
```

---

## Diagrama de Componentes del Frontend

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE
skinparam componentStyle rectangle

title Arquitectura de Componentes - Frontend React

package "App.js - Router" {
  [Router Principal] as Router
}

package "Pages" {
  [Login] as Login
  [Register] as Register
  [Dashboard] as Dashboard
  [AddDebt] as AddDebt
  [AddIncome] as AddIncome
  [AddPayment] as AddPayment
  [History] as History
}

package "Components" {
  [Sidebar] as Sidebar
  [ProfilePanel] as Profile
  [Toast] as Toast
}

package "Config" {
  [api.js - Axios] as API
}

cloud "Backend API" as Backend

Router --> Login
Router --> Register
Router --> Dashboard
Router --> AddDebt
Router --> AddIncome
Router --> AddPayment
Router --> History

Dashboard --> Sidebar
Dashboard --> Profile
AddDebt --> Toast
AddIncome --> Toast
History --> Toast

Login --> API
Register --> API
Dashboard --> API
AddDebt --> API
History --> API

API --> Backend : HTTPS

@enduml
```

---

## Diagrama de Secuencia - Crear Deuda

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Secuencia: Crear Nueva Deuda

actor Usuario as U
participant "Frontend\nReact" as F
participant "Handler\nHTTP" as H
participant "Service\nNegocio" as S
participant "Repository\nDatos" as R
database "SQLite" as DB

U -> F: Llenar formulario
U -> F: Click "Guardar"
F -> F: Validar campos

F -> H: POST /finances/debt\n(JSON + X-User-ID)
activate H

H -> H: Validar request
H -> S: CreateDebt(request)
activate S

S -> S: Aplicar reglas de negocio
S -> R: CreateDebt(debt)
activate R

R -> DB: INSERT INTO debts...
DB --> R: OK (lastInsertId)
R --> S: debt con ID
deactivate R

S --> H: debt
deactivate S

H -> H: ToDebtResponse(debt)
H --> F: 201 Created (JSON)
deactivate H

F -> F: Mostrar Toast "Éxito"
F -> F: Redireccionar
F --> U: Ver Dashboard actualizado

@enduml
```

---

## Diagrama de Base de Datos (Entidad-Relación)

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Diagrama Entidad-Relación - PayVue

entity "USERS" as users {
  * id : INTEGER <<PK>>
  --
  * email : TEXT <<UNIQUE>>
  * password_hash : TEXT
  * created_at : DATETIME
  * updated_at : DATETIME
}

entity "DEBTS" as debts {
  * id : INTEGER <<PK>>
  --
  * user_id : INTEGER <<FK>>
  * name : TEXT
  * total_amount : REAL
  * remaining_amount : REAL
  * due_date : DATETIME
  * interest_rate : REAL
  * num_installments : INTEGER
  * installment_amount : REAL
  * payment_day : INTEGER
  * paid : BOOLEAN
  * created_at : DATETIME
  * updated_at : DATETIME
}

entity "INCOMES" as incomes {
  * id : INTEGER <<PK>>
  --
  * user_id : INTEGER <<FK>>
  * amount : REAL
  * source : TEXT
  * date : DATETIME
  * created_at : DATETIME
  * updated_at : DATETIME
}

entity "PAYMENTS" as payments {
  * id : INTEGER <<PK>>
  --
  * user_id : INTEGER <<FK>>
  * debt_id : INTEGER <<FK>>
  * amount : REAL
  * receipt_filename : TEXT
  * date : DATETIME
  * created_at : DATETIME
  * updated_at : DATETIME
}

users ||--o{ debts : "tiene"
users ||--o{ incomes : "tiene"
users ||--o{ payments : "realiza"
debts ||--o{ payments : "recibe"

@enduml
```

---

## Diagrama de Paquetes del Backend

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Estructura de Paquetes - Backend Go

package "cmd" {
  package "server" {
    [main.go] as main
  }
  package "app" {
    [config] as config
    [container] as container
  }
}

package "pkg" {
  package "domain" {
    package "debt" {
      [entities.go]
      [service.go]
      [mapper.go]
      [container.go]
    }
    package "income" {
      [entities.go] as ie
      [service.go] as is
    }
    package "payment" {
      [entities.go] as pe
      [service.go] as ps
    }
    package "user" {
      [entities.go] as ue
      [service.go] as us
    }
  }
  
  package "repository" {
    package "database" {
      [database.go]
    }
    package "debt" as debtRepo {
      [repository.go]
    }
  }
  
  package "rest" {
    package "entities" {
      [entities.go] as re
    }
  }
}

main --> container
container --> config
container --> domain
domain --> repository

@enduml
```

---

## Beneficios de la Arquitectura

### Separación de Responsabilidades
- Cada capa tiene una responsabilidad clara
- Facilita el testing unitario
- Permite cambios aislados

### Escalabilidad
- Fácil agregar nuevos endpoints
- Nuevos módulos siguen la misma estructura
- Posibilidad de separar en microservicios

### Mantenibilidad
- Código organizado y predecible
- Fácil localizar y corregir bugs
- Documentación implícita por estructura

### Testabilidad
- Interfaces permiten mocks
- Capas independientes testeables
- Facilita pruebas de integración

---

## Código de Ejemplo por Capa

### 1. Capa de Presentación (Handler)

```go
func makeCreateDebtHandler(debtService debt.Service) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := getUserIDFromHeader(r)
        
        var request entities.CreateDebtRequest
        if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
            respondWithError(w, http.StatusBadRequest, "error", err.Error())
            return
        }
        
        d, err := debtService.CreateDebt(r.Context(), userID, request.ToDomain())
        if err != nil {
            respondWithError(w, http.StatusInternalServerError, "error", err.Error())
            return
        }
        
        respondWithJSON(w, http.StatusCreated, debt.ToDebtResponse(d))
    }
}
```

### 2. Capa de Dominio (Service)

```go
func (s *service) CreateDebt(ctx context.Context, userID int, req CreateDebtRequest) (*Debt, error) {
    debt := &Debt{
        UserID:            userID,
        Name:              req.Name,
        TotalAmount:       req.TotalAmount,
        RemainingAmount:   req.RemainingAmount,
        InstallmentAmount: req.InstallmentAmount,
        CreatedAt:         time.Now(),
        UpdatedAt:         time.Now(),
    }
    return s.Repository.CreateDebt(ctx, debt)
}
```

### 3. Capa de Datos (Repository)

```go
func (r *repository) CreateDebt(ctx context.Context, d *debt.Debt) (*debt.Debt, error) {
    query := `INSERT INTO debts (...) VALUES (?, ?, ?, ...)`
    result, err := r.db.ExecContext(ctx, query, d.UserID, d.Name, ...)
    if err != nil {
        return nil, err
    }
    id, _ := result.LastInsertId()
    d.ID = int(id)
    return d, nil
}
```

---

## Decisiones de Arquitectura

| Decisión | Justificación |
|----------|---------------|
| Go para Backend | Alto rendimiento, tipado estático, concurrencia |
| SQLite | Simplicidad, sin servidor separado, ideal para MVP |
| React | Ecosistema maduro, componentes reutilizables |
| Clean Architecture | Separación clara, testabilidad, mantenibilidad |
| REST API | Estándar de la industria, stateless |
| Docker | Consistencia entre ambientes, fácil despliegue |
