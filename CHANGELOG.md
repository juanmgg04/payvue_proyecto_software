# Changelog

Todos los cambios notables de este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/),
y este proyecto adhiere a [Semantic Versioning](https://semver.org/lang/es/).

## [Unreleased]

### Added
- Arquitectura CQRS con servicios Reader y Writer separados
- Servicio Reader (puerto 8080) para operaciones de consulta
- Servicio Writer (puerto 8081) para operaciones de escritura
- APIs REST completas:
  - Autenticación (registro, login, logout)
  - Gestión de ingresos (CRUD completo)
  - Gestión de deudas (CRUD completo)
  - Gestión de pagos (crear, listar, eliminar)
- Base de datos SQLite con persistencia
- Configuración con Docker y docker-compose
- Makefile con comandos de build y ejecución
- CI/CD con GitHub Actions:
  - Workflow de CI (formato, build, tests)
  - Workflow de Docker build
  - Auto-asignación de revisores en PRs
- Templates para Pull Requests e Issues
- Documentación completa:
  - Guía de CI/CD simplificada
  - Guía de configuración de GitHub
  - Guía de Git y conventional commits
  - Guía de troubleshooting
- Configuración de golangci-lint
- Configuración de .gitattributes para normalización de archivos
- CODEOWNERS para asignación automática de revisores
- Utilidad de carga de archivos (fileupload)
- Sistema de validación de requests con tags

### Changed
- Actualizado .gitignore para incluir artifacts de CI/CD
- Versión de Go establecida en 1.19 para compatibilidad

### Architecture
- Clean Architecture con separación por capas:
  - `pkg/domain/`: Lógica de negocio y entidades
  - `pkg/repository/`: Capa de acceso a datos
  - `pkg/rest/`: Capa HTTP y handlers
- Inyección de dependencias con contenedores
- Mappers entre capas para desacoplar DTOs de entidades de dominio
- Patrón Repository para acceso a datos
- Separación CQRS para escalabilidad

### Technical Details
- Go 1.19+
- Chi router v5 para routing HTTP
- SQLite como base de datos
- Docker multi-stage builds
- GitHub Actions para CI/CD
- Validación con go-playground/validator

---

## [0.1.0] - 2025-10-21

### Added
- Configuración inicial del proyecto
- Estructura de directorios base

[Unreleased]: https://github.com/tu-usuario/payvue-backend/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/tu-usuario/payvue-backend/releases/tag/v0.1.0

