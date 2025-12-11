# Patrones de Diseño

## Descripción General

PayVue implementa diversos patrones de diseño tanto en el backend (Go) como en el frontend (React) para garantizar código mantenible, escalable y de alta calidad.

---

## Resumen de Patrones Implementados

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrones de Diseño en PayVue

package "Patrones Creacionales" #LightBlue {
  [Factory] as F
  [Singleton] as S
  [Builder] as B
}

package "Patrones Estructurales" #LightGreen {
  [Repository] as R
  [Adapter/Mapper] as A
  [Facade] as Fa
}

package "Patrones de Comportamiento" #LightYellow {
  [Observer] as O
  [Interceptor] as I
  [Strategy] as St
}

note right of F : Creación de Container\ny dependencias
note right of S : Config y Database\ninstancia única
note right of R : Abstracción de\nacceso a datos
note right of A : Conversión de\nentidades a DTOs
note right of O : React hooks\nuseState/useEffect

@enduml
```

---

## 1. Patrón Repository

### Descripción
El patrón Repository actúa como una capa de abstracción entre la lógica de negocio y el acceso a datos.

### Diagrama de Clases

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrón Repository - PayVue

interface "Repository" as IRepo {
  + CreateDebt(ctx, debt): Debt
  + GetAllDebts(ctx): []Debt
  + GetDebtByID(ctx, id): Debt
  + UpdateDebt(ctx, debt): Debt
  + DeleteDebt(ctx, id): error
}

class "SQLiteRepository" as SQLite {
  - db: *sql.DB
  + CreateDebt(ctx, debt): Debt
  + GetAllDebts(ctx): []Debt
  + GetDebtByID(ctx, id): Debt
  + UpdateDebt(ctx, debt): Debt
  + DeleteDebt(ctx, id): error
}

class "Service" as Svc {
  - repository: Repository
  + CreateDebt(ctx, req): Debt
  + GetAllDebts(ctx): []Debt
}

IRepo <|.. SQLite : implements
Svc --> IRepo : uses

note right of IRepo
  Interface define el contrato
  Cualquier BD puede implementarla
end note

note right of SQLite
  Implementación concreta
  para SQLite
end note

@enduml
```

### Implementación en Go

```go
// Interface del Repository
type Repository interface {
    CreateDebt(ctx context.Context, debt *Debt) (*Debt, error)
    GetAllDebts(ctx context.Context) ([]Debt, error)
    GetDebtsByUserID(ctx context.Context, userID int) ([]Debt, error)
    GetDebtByID(ctx context.Context, id int) (*Debt, error)
    UpdateDebt(ctx context.Context, debt *Debt) (*Debt, error)
    DeleteDebt(ctx context.Context, id int) error
}

// Implementación concreta
type repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) debt.Repository {
    return &repository{db: db}
}

func (r *repository) CreateDebt(ctx context.Context, d *debt.Debt) (*debt.Debt, error) {
    query := `INSERT INTO debts (...) VALUES (...)`
    result, err := r.db.ExecContext(ctx, query, ...)
    return d, nil
}
```

### Beneficios
- ✅ Desacopla la lógica de negocio del acceso a datos
- ✅ Facilita el testing con mocks
- ✅ Permite cambiar la base de datos sin afectar servicios

---

## 2. Patrón Factory

### Descripción
El patrón Factory se utiliza para crear instancias de objetos complejos, encapsulando la lógica de creación.

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrón Factory - Container

class "Container" as C {
  + DebtService: debt.Service
  + IncomeService: income.Service
  + PaymentService: payment.Service
  + UserService: user.Service
  + Close(): void
}

class "ContainerFactory" as CF {
  + New(cfg: Config): Container
}

class "Config" as Cfg {
  + Port: string
  + DatabasePath: string
  + Environment: string
}

CF --> C : <<creates>>
CF --> Cfg : <<uses>>

note right of CF
  Factory Method
  Crea y configura
  todas las dependencias
end note

@enduml
```

### Implementación

```go
// Factory function - crea y configura todas las dependencias
func New(cfg config.Config) *Container {
    // Inicializar base de datos
    db, err := database.InitDB(cfg.DatabasePath)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Crear repositorios
    debtRepo := debtRepository.NewRepository(db)
    incomeRepo := incomeRepository.NewRepository(db)
    paymentRepo := paymentRepository.NewRepository(db)
    userRepo := userRepository.NewRepository(db)

    // Crear servicios
    return &Container{
        db:             db,
        DebtService:    debt.New(&debt.Container{Repository: debtRepo}),
        IncomeService:  income.New(&income.Container{Repository: incomeRepo}),
        PaymentService: payment.New(&payment.Container{Repository: paymentRepo}),
        UserService:    user.New(&user.Container{Repository: userRepo}),
    }
}
```

---

## 3. Patrón Singleton

### Descripción
Garantiza que una clase tenga solo una instancia y proporciona un punto de acceso global.

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrón Singleton - Configuración

class "Config" as C {
  - {static} instance: *Config
  - {static} once: sync.Once
  --
  + Port: string
  + DatabasePath: string
  + Environment: string
  --
  + {static} Get(): Config
}

note right of C
  sync.Once garantiza
  que solo se crea
  una instancia
end note

C --> C : instance

@enduml
```

### Implementación

```go
var configInstance *Config
var once sync.Once

func Get() Config {
    once.Do(func() {
        configInstance = &Config{
            Port:         getEnv("PORT", "8080"),
            DatabasePath: getEnv("DATABASE_PATH", "./data/payvue.db"),
            Environment:  getEnv("ENVIRONMENT", "development"),
        }
    })
    return *configInstance
}
```

---

## 4. Patrón Adapter/Mapper

### Descripción
Convierte la interfaz de una clase en otra que el cliente espera. Se usa para transformar entidades de dominio a DTOs.

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrón Adapter/Mapper

class "Debt" as D {
  + ID: int
  + UserID: int
  + Name: string
  + TotalAmount: float64
  + DueDate: time.Time
  + CreatedAt: time.Time
}

class "DebtResponse" as DR {
  + ID: int
  + Name: string
  + TotalAmount: float64
  + DueDate: string
  + RemainingPayments: int
}

class "Mapper" as M {
  + ToDebtResponse(debt): DebtResponse
  + ToDebtListResponse(debts): []DebtResponse
}

D --> M : input
M --> DR : output

note bottom of M
  Transforma entidades
  internas a DTOs
  de respuesta API
end note

@enduml
```

### Implementación

```go
func ToDebtResponse(d *Debt) DebtResponse {
    remainingPayments := 0
    if d.InstallmentAmount > 0 {
        remainingPayments = int(math.Ceil(d.RemainingAmount / d.InstallmentAmount))
    }

    return DebtResponse{
        ID:                d.ID,
        Name:              d.Name,
        TotalAmount:       d.TotalAmount,
        RemainingAmount:   d.RemainingAmount,
        DueDate:           d.DueDate.Format("2006-01-02"),
        RemainingPayments: remainingPayments,
        Paid:              d.Paid,
    }
}
```

---

## 5. Patrón Observer (Frontend)

### Descripción
En React, el patrón Observer se implementa mediante hooks que "observan" cambios en el estado.

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrón Observer - React Hooks

participant "Componente" as C
participant "useState" as S
participant "useEffect" as E
participant "API" as A

C -> S: Inicializar estado\n[debts, setDebts]
S --> C: Estado inicial []

C -> E: Registrar efecto\nuseEffect(fetchData, [])

E -> A: GET /finances/debt
A --> E: JSON data

E -> S: setDebts(data)
S --> C: Re-render con\nnuevos datos

note over E
  Observer pattern:
  useEffect "observa"
  cambios en dependencias
  y ejecuta callback
end note

@enduml
```

### Implementación

```jsx
function Dashboard() {
  // Estado observable
  const [debts, setDebts] = useState([]);

  // Observer - se ejecuta cuando cambian dependencias
  const fetchData = useCallback(async () => {
    try {
      const response = await api.get('/finances/debt');
      setDebts(response.data || []);
    } catch (error) {
      console.error('Error:', error);
    }
  }, []);

  // Suscripción a cambios
  useEffect(() => {
    fetchData();
    const interval = setInterval(fetchData, 5000);
    return () => clearInterval(interval);
  }, [fetchData]);

  return (
    <div>
      {debts.map(debt => <DebtCard key={debt.id} debt={debt} />)}
    </div>
  );
}
```

---

## 6. Patrón Interceptor (Frontend)

### Descripción
Los interceptores de Axios modifican requests/responses de forma centralizada.

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrón Interceptor - Axios

participant "Componente" as C
participant "Axios" as A
participant "Request\nInterceptor" as RI
participant "Response\nInterceptor" as RSI
participant "Backend" as B

C -> A: api.get('/finances/debt')
A -> RI: Interceptar request

RI -> RI: Agregar X-User-ID\ndesde localStorage

RI -> B: Request con headers

B --> RSI: Response

RSI -> RSI: Verificar errores\n(401 → logout)

RSI --> A: Response procesada
A --> C: Datos finales

note over RI
  Request Interceptor:
  - Agrega autenticación
  - Modifica headers
end note

note over RSI
  Response Interceptor:
  - Maneja errores globales
  - Redirige si 401
end note

@enduml
```

### Implementación

```javascript
const api = axios.create({
  baseURL: API_URL,
  headers: { 'Content-Type': 'application/json' }
});

// Request Interceptor
api.interceptors.request.use(
  (config) => {
    const user = JSON.parse(localStorage.getItem('user') || '{}');
    if (user.user_id) {
      config.headers['X-User-ID'] = user.user_id;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response Interceptor
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('user');
      window.location.href = '/';
    }
    return Promise.reject(error);
  }
);
```

---

## 7. Patrón Facade

### Descripción
Proporciona una interfaz simplificada a un conjunto de interfaces en un subsistema.

### Diagrama

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Patrón Facade - API Client

class "APIFacade" as F {
  + api: AxiosInstance
  + get(url): Promise
  + post(url, data): Promise
  + put(url, data): Promise
  + delete(url): Promise
  --
  + getCurrentUserId(): int
  + setUserData(data): void
  + clearUserData(): void
  + getUserData(): object
}

class "AxiosInstance" as A {
  + request(config): Promise
  + interceptors: object
}

class "LocalStorage" as LS {
  + getItem(key): string
  + setItem(key, value): void
  + removeItem(key): void
}

class "Interceptors" as I {
  + request: []
  + response: []
}

F --> A
F --> LS
A --> I

note right of F
  Facade simplifica
  el uso de múltiples
  subsistemas
end note

@enduml
```

---

## Diagrama de Relación de Patrones

```plantuml
@startuml
!theme plain
skinparam backgroundColor #FEFEFE

title Relación entre Patrones

package "Creación" #LightBlue {
  [Factory] as F
  [Singleton] as S
}

package "Estructura" #LightGreen {
  [Repository] as R
  [Adapter] as A
  [Facade] as Fa
}

package "Comportamiento" #LightYellow {
  [Observer] as O
  [Interceptor] as I
}

F --> R : crea
S --> F : configura
R --> A : usa
Fa --> I : contiene
O --> Fa : usa

note "Factory crea Repositories\ncon configuración Singleton" as N1
note "Observer usa Facade\npara llamadas API" as N2

@enduml
```

---

## Resumen de Patrones

| Patrón | Ubicación | Propósito |
|--------|-----------|-----------|
| **Repository** | `pkg/repository/` | Abstracción de acceso a datos |
| **Factory** | `cmd/app/container/` | Creación de dependencias |
| **Singleton** | `config/`, `database/` | Instancia única de config y DB |
| **Adapter/Mapper** | `pkg/domain/*/mapper.go` | Transformación de entidades |
| **Observer** | React Hooks | Reactividad de UI |
| **Interceptor** | `api.js` | Modificación de requests |
| **Facade** | `api.js` | Interfaz simplificada |

---

## Conclusión

La combinación de estos patrones proporciona:

1. **Código mantenible** - Cada patrón tiene una responsabilidad clara
2. **Testabilidad** - Las interfaces permiten mocks
3. **Escalabilidad** - Fácil agregar nuevas funcionalidades
4. **Reutilización** - Componentes y lógica reutilizable
5. **Desacoplamiento** - Capas independientes entre sí
