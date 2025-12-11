# Principios SOLID

## Descripción General

Los principios SOLID son cinco principios de diseño orientado a objetos que hacen el software más comprensible, flexible y mantenible.

---

## Los 5 Principios SOLID

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Principios SOLID

package "S - Single Responsibility" #LightBlue {
  note "Una clase, una razón\npara cambiar" as N1
}

package "O - Open/Closed" #LightGreen {
  note "Abierto a extensión\nCerrado a modificación" as N2
}

package "L - Liskov Substitution" #LightYellow {
  note "Subtipos sustituibles\npor tipos base" as N3
}

package "I - Interface Segregation" #LightPink {
  note "Interfaces pequeñas\ny específicas" as N4
}

package "D - Dependency Inversion" #LightCyan {
  note "Depender de\nabstracciones" as N5
}

@enduml
```

---

## S - Single Responsibility Principle (SRP)

### Principio
> "Una clase debe tener una, y solo una, razón para cambiar."

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title SRP - Separación de Responsabilidades

class "DebtHandler" as H {
  + HandleCreate(w, r)
  + HandleGet(w, r)
  + ParseRequest(r)
  + FormatResponse(w)
}
note right: Solo HTTP\nNo lógica de negocio

class "DebtService" as S {
  + CreateDebt(ctx, req)
  + ValidateDebt(debt)
  + CalculateInterest(debt)
}
note right: Solo lógica de negocio\nNo acceso a datos

class "DebtRepository" as R {
  + Create(ctx, debt)
  + FindAll(ctx)
  + FindByID(ctx, id)
  + Update(ctx, debt)
  + Delete(ctx, id)
}
note right: Solo acceso a datos\nNo HTTP

H --> S : usa
S --> R : usa

@enduml
```

### Implementación - Separación de Responsabilidades

```go
// Handler - SOLO manejo HTTP
func makeCreateDebtHandler(service debt.Service) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var request entities.CreateDebtRequest
        json.NewDecoder(r.Body).Decode(&request)
        debt, _ := service.CreateDebt(r.Context(), request.ToDomain())
        respondWithJSON(w, http.StatusCreated, debt)
    }
}

// Service - SOLO lógica de negocio
func (s *service) CreateDebt(ctx context.Context, req CreateDebtRequest) (*Debt, error) {
    if req.TotalAmount <= 0 {
        return nil, ErrInvalidAmount
    }
    debt := &Debt{Name: req.Name, TotalAmount: req.TotalAmount}
    return s.Repository.CreateDebt(ctx, debt)
}

// Repository - SOLO acceso a datos
func (r *repository) CreateDebt(ctx context.Context, d *debt.Debt) (*debt.Debt, error) {
    query := `INSERT INTO debts (...) VALUES (...)`
    result, _ := r.db.ExecContext(ctx, query, ...)
    return d, nil
}
```

---

## O - Open/Closed Principle (OCP)

### Principio
> "Las entidades deben estar abiertas para extensión, pero cerradas para modificación."

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title OCP - Extensión sin Modificación

interface "Repository" as IR {
  + Create(ctx, entity)
  + FindAll(ctx)
  + FindByID(ctx, id)
  + Update(ctx, entity)
  + Delete(ctx, id)
}

class "SQLiteRepository" as SQLite {
  - db: *sql.DB
  + Create(ctx, entity)
  + FindAll(ctx)
  + FindByID(ctx, id)
}

class "PostgresRepository" as Postgres {
  - db: *sql.DB
  + Create(ctx, entity)
  + FindAll(ctx)
  + FindByID(ctx, id)
}

class "MongoRepository" as Mongo {
  - client: *mongo.Client
  + Create(ctx, entity)
  + FindAll(ctx)
  + FindByID(ctx, id)
}

IR <|.. SQLite
IR <|.. Postgres
IR <|.. Mongo

note bottom of IR
  La interface está CERRADA
  No necesita modificarse
end note

note bottom of Mongo
  Nuevas implementaciones
  EXTIENDEN sin modificar
  código existente
end note

@enduml
```

### Implementación

```go
// Interface (cerrada a modificación)
type Repository interface {
    CreateDebt(ctx context.Context, debt *Debt) (*Debt, error)
    GetAllDebts(ctx context.Context) ([]Debt, error)
}

// Implementación SQLite existente
type sqliteRepository struct {
    db *sql.DB
}

// Nueva implementación PostgreSQL (extensión)
type postgresRepository struct {
    db *sql.DB
}

func (r *postgresRepository) CreateDebt(ctx context.Context, d *debt.Debt) (*debt.Debt, error) {
    // Implementación específica para PostgreSQL
    query := `INSERT INTO debts (...) VALUES (...) RETURNING id`
    return d, nil
}
```

---

## L - Liskov Substitution Principle (LSP)

### Principio
> "Los objetos deben ser reemplazables por instancias de sus subtipos sin alterar el funcionamiento."

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title LSP - Sustitución de Liskov

class "Service" as S {
  - repository: Repository
  + GetAllDebts(ctx, userID)
}

interface "Repository" as IR {
  + GetDebtsByUserID(ctx, userID)
}

class "SQLiteRepository" as SQLite {
  + GetDebtsByUserID(ctx, userID)
}

class "MockRepository" as Mock {
  + GetDebtsByUserID(ctx, userID)
}

S --> IR : usa interface
IR <|.. SQLite
IR <|.. Mock

note right of S
  El Service funciona
  igual con cualquier
  implementación
end note

note bottom of Mock
  Mock es sustituible
  por SQLite sin
  cambiar el Service
end note

@enduml
```

### Implementación

```go
// El Service puede usar cualquier implementación de Repository
type service struct {
    Repository Repository // Interface
}

func (s *service) GetAllDebts(ctx context.Context, userID int) ([]Debt, error) {
    // Este código funciona con SQLiteRepo, PostgresRepo, MockRepo, etc.
    return s.Repository.GetDebtsByUserID(ctx, userID)
}

// MockRepository cumple el mismo contrato
type MockRepository struct {
    debts []debt.Debt
}

func (m *MockRepository) GetDebtsByUserID(ctx context.Context, userID int) ([]debt.Debt, error) {
    var result []debt.Debt
    for _, d := range m.debts {
        if d.UserID == userID {
            result = append(result, d)
        }
    }
    return result, nil
}

// En tests - sustitución sin problemas
func TestGetAllDebts(t *testing.T) {
    mockRepo := &MockRepository{debts: testDebts}
    service := debt.NewService(&debt.Container{Repository: mockRepo})
    debts, err := service.GetAllDebts(context.Background(), 1)
    // Funciona exactamente igual
}
```

---

## I - Interface Segregation Principle (ISP)

### Principio
> "Muchas interfaces específicas son mejores que una interfaz de propósito general."

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title ISP - Segregación de Interfaces

interface "DebtReader" as R {
  + GetAllDebts(ctx)
  + GetDebtByID(ctx, id)
  + GetDebtsByUserID(ctx, userID)
}

interface "DebtWriter" as W {
  + CreateDebt(ctx, debt)
  + UpdateDebt(ctx, debt)
  + DeleteDebt(ctx, id)
}

interface "Repository" as Full {
}

R <|-- Full
W <|-- Full

class "ReadOnlyHandler" as ROH {
  - reader: DebtReader
}

class "WriteHandler" as WH {
  - writer: DebtWriter
}

class "FullRepository" as FR {
}

Full <|.. FR
ROH --> R : solo lectura
WH --> W : solo escritura

note right of R
  Clientes que solo leen
  no necesitan métodos
  de escritura
end note

note right of W
  Clientes que solo escriben
  no necesitan métodos
  de lectura
end note

@enduml
```

### Implementación

```go
// Interfaces segregadas por funcionalidad

// Interface solo para lectura
type DebtReader interface {
    GetAllDebts(ctx context.Context) ([]Debt, error)
    GetDebtsByUserID(ctx context.Context, userID int) ([]Debt, error)
    GetDebtByID(ctx context.Context, id int) (*Debt, error)
}

// Interface solo para escritura
type DebtWriter interface {
    CreateDebt(ctx context.Context, debt *Debt) (*Debt, error)
    UpdateDebt(ctx context.Context, debt *Debt) (*Debt, error)
    DeleteDebt(ctx context.Context, id int) error
}

// Interface completa (composición)
type Repository interface {
    DebtReader
    DebtWriter
}

// Handler que solo necesita lectura
type readOnlyHandler struct {
    reader DebtReader // No expone métodos de escritura
}

// Handler que solo necesita escritura
type writeHandler struct {
    writer DebtWriter // No expone métodos de lectura
}
```

---

## D - Dependency Inversion Principle (DIP)

### Principio
> "Depende de abstracciones, no de implementaciones concretas."

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title DIP - Inversión de Dependencias

package "Alto Nivel" #LightBlue {
  class "Service" as S {
    - repository: Repository
    + CreateDebt(ctx, req)
  }
}

package "Abstracción" #LightYellow {
  interface "Repository" as IR {
    + CreateDebt(ctx, debt)
    + GetDebtsByUserID(ctx, userID)
  }
}

package "Bajo Nivel" #LightGreen {
  class "SQLiteRepository" as SQLite {
    - db: *sql.DB
    + CreateDebt(ctx, debt)
    + GetDebtsByUserID(ctx, userID)
  }
}

S --> IR : depende de abstracción
SQLite ..|> IR : implementa

note right of S
  Service depende de
  la INTERFACE, no de
  SQLiteRepository
end note

note right of IR
  La abstracción es
  el punto de unión
end note

note right of SQLite
  Implementación concreta
  puede cambiar sin
  afectar al Service
end note

@enduml
```

### Implementación

```go
// El Service depende de la abstracción (interface)
type service struct {
    Repository Repository // Interface, NO *sqliteRepository
}

func New(container *Container) Service {
    return &service{
        Repository: container.Repository, // Se inyecta la dependencia
    }
}

func (s *service) CreateDebt(ctx context.Context, req CreateDebtRequest) (*Debt, error) {
    // El service no sabe qué implementación de Repository usa
    debt := &Debt{...}
    return s.Repository.CreateDebt(ctx, debt) // Usa la abstracción
}

// Inyección de Dependencias en Container
func New(cfg config.Config) *Container {
    db, _ := database.InitDB(cfg.DatabasePath)
    
    // Crear implementación concreta
    debtRepo := debtRepository.NewRepository(db) // Retorna interface
    
    // Inyectar dependencia abstracta
    debtContainer := &debt.Container{
        Repository: debtRepo, // Interface, no implementación
    }
    
    return &Container{
        DebtService: debt.New(debtContainer),
    }
}
```

---

## Diagrama de Aplicación Completa de SOLID

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Aplicación de SOLID en PayVue

package "S - Single Responsibility" {
  [Handler] as H1
  [Service] as S1
  [Repository] as R1
  H1 --> S1
  S1 --> R1
}

package "O - Open/Closed" {
  interface "IRepo" as IR
  [SQLite] as SQL
  [Postgres] as PG
  IR <|.. SQL
  IR <|.. PG
}

package "L - Liskov Substitution" {
  [Service] as S2
  interface "Repository" as IR2
  [Real] as Real
  [Mock] as Mock
  S2 --> IR2
  IR2 <|.. Real
  IR2 <|.. Mock
}

package "I - Interface Segregation" {
  interface "Reader" as Reader
  interface "Writer" as Writer
  interface "Full" as Full
  Reader <|-- Full
  Writer <|-- Full
}

package "D - Dependency Inversion" {
  [Alto Nivel] as Alto
  interface "Abstracción" as Abs
  [Bajo Nivel] as Bajo
  Alto --> Abs
  Bajo ..|> Abs
}

@enduml
```

---

## Resumen de Principios SOLID en PayVue

| Principio | Implementación | Beneficio |
|-----------|---------------|-----------|
| **S**ingle Responsibility | Handler, Service, Repository separados | Código fácil de mantener |
| **O**pen/Closed | Interfaces para extensión | Agregar features sin modificar |
| **L**iskov Substitution | Repositories intercambiables | Testing con mocks |
| **I**nterface Segregation | Interfaces pequeñas (Reader, Writer) | Componentes desacoplados |
| **D**ependency Inversion | Inyección de dependencias | Flexibilidad y testing |

---

## Conclusión

La aplicación consistente de SOLID en PayVue resulta en:

1. **Código mantenible** - Cambios localizados y predecibles
2. **Alta cohesión** - Componentes enfocados en una tarea
3. **Bajo acoplamiento** - Dependencias mediante abstracciones
4. **Facilidad de testing** - Mocks e inyección de dependencias
5. **Extensibilidad** - Nuevas features sin romper existentes
