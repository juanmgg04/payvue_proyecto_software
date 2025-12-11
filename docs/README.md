# PayVue - Sistema de Gestión Financiera Personal

<p align="center">
  <strong>Documentación Técnica del Proyecto</strong>
</p>

---

## 📋 Descripción del Proyecto

**PayVue** es una aplicación web de gestión financiera personal que permite a los usuarios:

- 📊 Registrar y visualizar ingresos
- 💳 Administrar deudas y cuotas
- 💰 Realizar seguimiento de pagos
- 📈 Ver estadísticas y gráficos financieros
- 📎 Adjuntar recibos de pago

## 🏗️ Stack Tecnológico

### Backend
| Tecnología | Versión | Propósito |
|------------|---------|-----------|
| Go (Golang) | 1.21+ | Lenguaje de programación |
| Chi Router | v5 | Enrutamiento HTTP |
| SQLite | 3 | Base de datos |
| Docker | 24+ | Contenedorización |

### Frontend
| Tecnología | Versión | Propósito |
|------------|---------|-----------|
| React | 18+ | Framework UI |
| Axios | - | Cliente HTTP |
| Chart.js | - | Gráficos |
| CSS3 | - | Estilos |

### Despliegue
| Servicio | Propósito |
|----------|-----------|
| Render | Backend API |
| Vercel | Frontend |
| GitHub | Control de versiones |

## 📁 Estructura del Proyecto

```
payvue_proyecto_software/
├── backend/                    # Código del servidor
│   ├── cmd/
│   │   └── server/            # Punto de entrada
│   ├── pkg/
│   │   ├── domain/            # Lógica de negocio
│   │   ├── repository/        # Acceso a datos
│   │   └── rest/              # Handlers HTTP
│   └── Dockerfile
├── frontend/                   # Código del cliente
│   ├── src/
│   │   ├── components/        # Componentes React
│   │   ├── pages/             # Páginas
│   │   └── config/            # Configuración
│   └── Dockerfile
├── docs/                       # Documentación
└── docker-compose.yml          # Orquestación
```

## 🎯 Características Principales

### Autenticación
- Registro de usuarios
- Inicio de sesión
- Separación de datos por usuario

### Gestión de Deudas
- Crear, editar y eliminar deudas
- Seguimiento de cuotas
- Cálculo de días hasta vencimiento

### Gestión de Ingresos
- Registrar fuentes de ingreso
- Historial por fecha
- Cálculo de totales

### Gestión de Pagos
- Registrar pagos a deudas
- Subir recibos/facturas
- Actualización automática de saldo

### Dashboard
- Estadísticas en tiempo real
- Gráficos de ingresos
- Lista de deudas próximas

## 👥 Equipo de Desarrollo

| Rol | Nombre |
|-----|--------|
| Líder de Proyecto | Juan Miguel Valencia Atehortua |
| Desarrollador | Juan Andres Forero Guauque |

## 📅 Información del Proyecto

- **Fecha de inicio:** Noviembre 2024
- **Fecha de entrega:** Diciembre 2024
- **Versión actual:** 1.0.0
- **Estado:** ✅ Completado

## 🔗 Enlaces

- [Repositorio GitHub](https://github.com/juanmgg04/payvue_proyecto_software)
- [Demo Frontend](https://payvue.vercel.app)
- [API Backend](https://payvue-api.onrender.com)

---

<p align="center">
  <em>Proyecto de Software - Universidad 2024</em>
</p>

