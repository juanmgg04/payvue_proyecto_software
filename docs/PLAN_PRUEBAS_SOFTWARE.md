# Plan de Pruebas de Software

## PayVue - Sistema de Gestión Financiera Personal

**Fecha:** 10/12/2024

---

## Historial de Versiones

| Fecha | Versión | Autor | Organización | Descripción |
|-------|---------|-------|--------------|-------------|
| 10/12/2024 | 1.0 | Juan Miguel Valencia A. / Juan Andres Forero G. | Universidad | Versión inicial del plan de pruebas |

---

## Información del Proyecto

| Campo | Valor |
|-------|-------|
| **Empresa / Organización** | Universidad - Proyecto de Software |
| **Proyecto** | PayVue - Sistema de Gestión Financiera Personal |
| **Fecha de preparación** | 10/12/2024 |
| **Cliente** | Universidad |
| **Patrocinador principal** | Docente de Proyecto de Software |
| **Gerente / Líder de Proyecto** | Juan Miguel Valencia Atehortua |
| **Gerente / Líder de Pruebas de Software** | Juan Andres Forero Guauque |

---

## Aprobaciones

| Nombre y Apellido | Cargo | Departamento u Organización | Fecha | Firma |
|-------------------|-------|----------------------------|-------|-------|
| Juan Miguel Valencia Atehortua | Líder de Proyecto | Ingeniería de Software | 10/12/2024 | _______ |
| Juan Andres Forero Guauque | Desarrollador | Ingeniería de Software | 10/12/2024 | _______ |
| Docente | Supervisor | Universidad | 10/12/2024 | _______ |

---

## Resumen Ejecutivo

El presente documento describe el Plan de Pruebas de Software para el proyecto **PayVue**, una aplicación web de gestión financiera personal que permite a los usuarios registrar ingresos, administrar deudas y realizar seguimiento de pagos.

### Propósito
Este plan de pruebas tiene como objetivo garantizar la calidad del software mediante la ejecución sistemática de pruebas en diferentes niveles (unitarias, integración, sistema y aceptación) y técnicas (caja blanca, caja negra y caja gris).

### Alcance
- **Backend**: API REST desarrollada en Go (Golang) con base de datos SQLite
- **Frontend**: Aplicación React con interfaz de usuario moderna
- **Integración**: Comunicación entre frontend y backend mediante HTTP/REST

### Restricciones
- Tiempo limitado de desarrollo (proyecto académico)
- Recursos de hardware limitados (desarrollo local y hosting gratuito)
- Equipo de desarrollo reducido

---

## Alcance de las Pruebas

### Elementos de Pruebas

#### Backend (Go)
| Módulo | Componentes |
|--------|-------------|
| Autenticación | Registro, Login, Logout |
| Gestión de Deudas | CRUD de deudas, filtrado por usuario |
| Gestión de Ingresos | CRUD de ingresos, filtrado por usuario |
| Gestión de Pagos | Crear pago, subir recibo, eliminar pago |
| Base de Datos | Conexión SQLite, migraciones, persistencia |

#### Frontend (React)
| Módulo | Componentes |
|--------|-------------|
| Autenticación | Login, Register, ForgotPassword |
| Dashboard | Estadísticas, gráficos, resumen financiero |
| Formularios | AddDebt, AddIncome, AddPayment |
| Historial | Listado, filtros, edición, eliminación |
| Navegación | Sidebar, ProfilePanel |

### Nuevas Funcionalidades a Probar

1. **Autenticación de usuarios**
   - Registro con email y contraseña
   - Inicio de sesión
   - Cierre de sesión con limpieza de datos

2. **Gestión de deudas**
   - Crear nueva deuda con todos los campos requeridos
   - Editar deuda existente
   - Eliminar deuda
   - Marcar deuda como pagada

3. **Gestión de ingresos**
   - Registrar nuevo ingreso
   - Editar ingreso
   - Eliminar ingreso
   - Filtrar por fecha

4. **Gestión de pagos**
   - Registrar pago asociado a una deuda
   - Subir recibo/factura
   - Ver recibo
   - Eliminar pago

5. **Dashboard**
   - Visualización de estadísticas
   - Gráficos de ingresos y gastos
   - Lista de deudas próximas a vencer

6. **Separación de datos por usuario**
   - Cada usuario ve solo sus propios datos
   - Los datos no se mezclan entre usuarios

### Funcionalidades a No Probar

| Funcionalidad | Razón | Riesgo Asumido |
|---------------|-------|----------------|
| Autenticación con Google | No implementada en esta versión | Bajo - funcionalidad opcional |
| Recuperación de contraseña real | Requiere servicio de email | Bajo - se simula en frontend |
| Notificaciones push | Fuera del alcance del proyecto | Bajo - mejora futura |
| Exportación de reportes PDF | No implementada | Medio - funcionalidad deseable |

---

## Enfoque de Pruebas (Estrategia)

### Tipos de Pruebas a Realizar

#### 1. Pruebas Unitarias (Caja Blanca)

**Backend (Go)**
- Pruebas de funciones de servicio
- Pruebas de validación de datos
- Pruebas de mappers y conversiones

**Frontend (React)**
- Pruebas de componentes individuales
- Pruebas de funciones helper
- Pruebas de hooks personalizados

#### 2. Pruebas de Integración (Caja Gris)

- Integración Backend-Base de Datos
- Integración Frontend-Backend (API)
- Integración de módulos internos

#### 3. Pruebas de Sistema (Caja Negra)

- Flujos completos de usuario
- Escenarios de negocio end-to-end
- Pruebas de interfaz de usuario

#### 4. Pruebas de Aceptación

- Validación de requisitos funcionales
- Pruebas con usuarios finales
- Verificación de criterios de aceptación

#### 5. Pruebas No Funcionales

**Rendimiento**
- Tiempo de respuesta de API < 500ms
- Carga de dashboard < 3 segundos

**Seguridad**
- Validación de inputs
- Protección contra inyección SQL
- Separación de datos por usuario

**Usabilidad**
- Interfaz intuitiva
- Mensajes de error claros
- Responsive design

---

## Criterios de Aceptación o Rechazo

### Criterios de Aceptación

| Criterio | Métrica |
|----------|---------|
| Pruebas unitarias completadas | 100% de casos ejecutados |
| Casos de prueba exitosos | ≥ 90% |
| Defectos críticos | 0 sin resolver |
| Defectos mayores | ≤ 2 sin resolver |
| Cobertura de funcionalidades | 100% de features principales |
| Pruebas de regresión | 100% ejecutadas |

### Criterios de Suspensión

- Más del 50% de casos de prueba fallidos en un ciclo
- Defectos críticos que impidan continuar las pruebas
- Ambiente de pruebas no disponible
- Cambios significativos en requerimientos

### Criterios de Reanudación

- Corrección de defectos críticos verificada
- Ambiente de pruebas restaurado
- Aprobación del líder de pruebas
- Nueva build desplegada en ambiente de pruebas

---

## Entregables

1. ✅ Documento de Plan de Pruebas (este documento)
2. ✅ Casos de Prueba documentados
3. ✅ Resultados de ejecución de pruebas
4. ✅ Log de defectos encontrados
5. ✅ Evidencias de pruebas (capturas de pantalla)
6. ✅ Reporte final de pruebas

---

## Recursos

### Requerimientos de Entornos – Hardware

| Recurso | Especificación |
|---------|----------------|
| PC Desarrollo/Pruebas | Intel i5+, 8GB RAM, 256GB SSD |
| Servidor Backend | Render Free Tier (512MB RAM) |
| Servidor Frontend | Vercel Free Tier |
| Base de Datos | SQLite (archivo local/persistente) |

### Requerimientos de Entornos – Software

| Software | Versión | Propósito |
|----------|---------|-----------|
| Go | 1.21+ | Backend |
| Node.js | 18+ | Frontend |
| Docker | 24+ | Contenedores |
| Git | 2.40+ | Control de versiones |
| Chrome/Firefox | Última | Navegador de pruebas |
| Postman | Última | Pruebas de API |

### Herramientas de Pruebas Requeridas

| Herramienta | Uso |
|-------------|-----|
| Postman | Pruebas manuales de API REST |
| Chrome DevTools | Debugging frontend |
| Go testing | Pruebas unitarias backend |
| Jest/React Testing Library | Pruebas unitarias frontend |
| curl | Pruebas rápidas de endpoints |

### Personal

| Rol | Nombre | Responsabilidades |
|-----|--------|-------------------|
| Líder de Proyecto / Pruebas | Juan Miguel Valencia Atehortua | Planificación, supervisión, desarrollo |
| Desarrollador / Tester | Juan Andres Forero Guauque | Desarrollo, ejecución de pruebas |

### Entrenamiento

- Capacitación en uso del sistema PayVue
- Conocimiento de la API REST
- Manejo de herramientas de pruebas

---

## Definición de Casos de Prueba

### CP-001: Registro de Usuario

| Campo | Valor |
|-------|-------|
| **Código** | CP-001 |
| **Nombre** | Registro de Usuario Exitoso |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que un usuario nuevo puede registrarse correctamente en el sistema |
| **Prerrequisitos** | - Sistema desplegado y accesible<br>- Email no registrado previamente |
| **Pasos** | 1. Acceder a la página de registro<br>2. Ingresar email válido<br>3. Ingresar contraseña (mín. 6 caracteres)<br>4. Click en "Guardar" |
| **Resultado esperado** | - Mensaje "Usuario registrado exitosamente"<br>- Redirección a página de login |
| **Resultado obtenido** | ✅ EXITOSO - Usuario registrado correctamente, redirección funciona |

---

### CP-002: Registro con Email Inválido

| Campo | Valor |
|-------|-------|
| **Código** | CP-002 |
| **Nombre** | Registro con Email Inválido |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que el sistema rechaza emails con formato inválido |
| **Prerrequisitos** | - Sistema desplegado y accesible |
| **Pasos** | 1. Acceder a la página de registro<br>2. Ingresar email inválido (ej: "test")<br>3. Ingresar contraseña válida<br>4. Click en "Guardar" |
| **Resultado esperado** | - Mensaje de error de validación<br>- No se crea el usuario |
| **Resultado obtenido** | ✅ EXITOSO - Sistema muestra error de validación |

---

### CP-003: Registro con Contraseña Corta

| Campo | Valor |
|-------|-------|
| **Código** | CP-003 |
| **Nombre** | Registro con Contraseña Corta |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que el sistema rechaza contraseñas menores a 6 caracteres |
| **Prerrequisitos** | - Sistema desplegado y accesible |
| **Pasos** | 1. Acceder a la página de registro<br>2. Ingresar email válido<br>3. Ingresar contraseña corta (ej: "123")<br>4. Click en "Guardar" |
| **Resultado esperado** | - Mensaje "La contraseña debe tener al menos 6 caracteres" |
| **Resultado obtenido** | ✅ EXITOSO - Validación funciona correctamente |

---

### CP-004: Login Exitoso

| Campo | Valor |
|-------|-------|
| **Código** | CP-004 |
| **Nombre** | Inicio de Sesión Exitoso |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que un usuario registrado puede iniciar sesión |
| **Prerrequisitos** | - Usuario previamente registrado<br>- Sistema desplegado |
| **Pasos** | 1. Acceder a la página de login<br>2. Ingresar email registrado<br>3. Ingresar contraseña correcta<br>4. Click en "Iniciar Sesión" |
| **Resultado esperado** | - Redirección al Dashboard<br>- Usuario autenticado<br>- Datos del usuario en localStorage |
| **Resultado obtenido** | ✅ EXITOSO - Login funciona, redirección correcta |

---

### CP-005: Login con Credenciales Inválidas

| Campo | Valor |
|-------|-------|
| **Código** | CP-005 |
| **Nombre** | Login con Credenciales Inválidas |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que el sistema rechaza credenciales incorrectas |
| **Prerrequisitos** | - Sistema desplegado |
| **Pasos** | 1. Acceder a la página de login<br>2. Ingresar email no registrado o contraseña incorrecta<br>3. Click en "Iniciar Sesión" |
| **Resultado esperado** | - Mensaje "Credenciales inválidas"<br>- No hay redirección |
| **Resultado obtenido** | ✅ EXITOSO - Error mostrado correctamente |

---

### CP-006: Crear Deuda

| Campo | Valor |
|-------|-------|
| **Código** | CP-006 |
| **Nombre** | Crear Nueva Deuda |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que un usuario puede crear una nueva deuda |
| **Prerrequisitos** | - Usuario autenticado<br>- Sesión activa |
| **Pasos** | 1. Ir a "Agregar Deuda"<br>2. Completar todos los campos obligatorios<br>3. Click en "Guardar" |
| **Resultado esperado** | - Mensaje "¡Deuda guardada con éxito!"<br>- Deuda visible en Dashboard y Historial |
| **Resultado obtenido** | ✅ EXITOSO - Deuda creada y visible |

---

### CP-007: Crear Deuda sin Campos Requeridos

| Campo | Valor |
|-------|-------|
| **Código** | CP-007 |
| **Nombre** | Crear Deuda sin Campos Obligatorios |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar validación de campos obligatorios |
| **Prerrequisitos** | - Usuario autenticado |
| **Pasos** | 1. Ir a "Agregar Deuda"<br>2. Dejar campos vacíos<br>3. Click en "Guardar" |
| **Resultado esperado** | - Mensaje de error de validación<br>- Formulario no se envía |
| **Resultado obtenido** | ✅ EXITOSO - Validación HTML5 y backend funcionan |

---

### CP-008: Crear Ingreso

| Campo | Valor |
|-------|-------|
| **Código** | CP-008 |
| **Nombre** | Crear Nuevo Ingreso |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que un usuario puede registrar un ingreso |
| **Prerrequisitos** | - Usuario autenticado |
| **Pasos** | 1. Ir a "Agregar Ingresos"<br>2. Ingresar monto, fuente y fecha<br>3. Click en "Guardar" |
| **Resultado esperado** | - Mensaje "¡Ingreso guardado con éxito!"<br>- Ingreso visible en Dashboard |
| **Resultado obtenido** | ✅ EXITOSO - Ingreso creado correctamente |

---

### CP-009: Crear Pago con Recibo

| Campo | Valor |
|-------|-------|
| **Código** | CP-009 |
| **Nombre** | Crear Pago con Archivo de Recibo |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que un usuario puede registrar un pago con recibo adjunto |
| **Prerrequisitos** | - Usuario autenticado<br>- Al menos una deuda existente<br>- Archivo de imagen disponible |
| **Pasos** | 1. Ir a "Agregar Pagos"<br>2. Ingresar monto<br>3. Seleccionar deuda<br>4. Subir recibo<br>5. Click en "Guardar" |
| **Resultado esperado** | - Pago registrado<br>- Recibo almacenado<br>- Botón "Ver Recibo" en historial |
| **Resultado obtenido** | ✅ EXITOSO - Pago y recibo guardados |

---

### CP-010: Ver Recibo de Pago

| Campo | Valor |
|-------|-------|
| **Código** | CP-010 |
| **Nombre** | Visualizar Recibo de Pago |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que se puede ver el recibo adjunto de un pago |
| **Prerrequisitos** | - Pago con recibo existente |
| **Pasos** | 1. Ir a Historial > Pagos<br>2. Click en "Ver Recibo" |
| **Resultado esperado** | - Modal con imagen del recibo<br>- Imagen carga correctamente |
| **Resultado obtenido** | ✅ EXITOSO - Recibo se visualiza correctamente |

---

### CP-011: Editar Deuda

| Campo | Valor |
|-------|-------|
| **Código** | CP-011 |
| **Nombre** | Editar Deuda Existente |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que un usuario puede editar una deuda |
| **Prerrequisitos** | - Deuda existente<br>- Usuario autenticado |
| **Pasos** | 1. Ir a Historial > Deudas<br>2. Click en icono de editar<br>3. Modificar campos<br>4. Click en "Guardar" |
| **Resultado esperado** | - Mensaje "¡Actualizado con éxito!"<br>- Cambios reflejados |
| **Resultado obtenido** | ✅ EXITOSO - Deuda actualizada |

---

### CP-012: Eliminar Ingreso

| Campo | Valor |
|-------|-------|
| **Código** | CP-012 |
| **Nombre** | Eliminar Ingreso |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que un usuario puede eliminar un ingreso |
| **Prerrequisitos** | - Ingreso existente |
| **Pasos** | 1. Ir a Historial > Ingresos<br>2. Click en icono de eliminar<br>3. Confirmar eliminación |
| **Resultado esperado** | - Mensaje "¡Eliminado con éxito!"<br>- Ingreso removido de la lista |
| **Resultado obtenido** | ✅ EXITOSO - Ingreso eliminado |

---

### CP-013: Filtrar por Fecha en Historial

| Campo | Valor |
|-------|-------|
| **Código** | CP-013 |
| **Nombre** | Filtrar Registros por Fecha |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que los filtros de fecha funcionan |
| **Prerrequisitos** | - Registros con diferentes fechas |
| **Pasos** | 1. Ir a Historial<br>2. Seleccionar fecha "Desde"<br>3. Seleccionar fecha "Hasta"<br>4. Verificar resultados |
| **Resultado esperado** | - Solo se muestran registros en el rango de fechas |
| **Resultado obtenido** | ✅ EXITOSO - Filtro funciona correctamente |

---

### CP-014: Separación de Datos por Usuario

| Campo | Valor |
|-------|-------|
| **Código** | CP-014 |
| **Nombre** | Verificar Separación de Datos por Usuario |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que cada usuario solo ve sus propios datos |
| **Prerrequisitos** | - Dos usuarios registrados<br>- Cada uno con datos propios |
| **Pasos** | 1. Login con Usuario A<br>2. Verificar datos de A<br>3. Logout<br>4. Login con Usuario B<br>5. Verificar que no ve datos de A |
| **Resultado esperado** | - Usuario B no ve datos de Usuario A<br>- Cada usuario ve solo sus datos |
| **Resultado obtenido** | ✅ EXITOSO - Datos separados correctamente |

---

### CP-015: Cerrar Sesión

| Campo | Valor |
|-------|-------|
| **Código** | CP-015 |
| **Nombre** | Cerrar Sesión y Limpiar Datos |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que al cerrar sesión se limpian los datos del localStorage |
| **Prerrequisitos** | - Usuario autenticado |
| **Pasos** | 1. Click en "Cerrar Sesión" en el sidebar<br>2. Verificar redirección a login<br>3. Verificar localStorage vacío |
| **Resultado esperado** | - Redirección a login<br>- localStorage limpio<br>- No se puede acceder al dashboard |
| **Resultado obtenido** | ✅ EXITOSO - Logout funciona correctamente |

---

### CP-016: Prueba de API - Health Check

| Campo | Valor |
|-------|-------|
| **Código** | CP-016 |
| **Nombre** | Verificar Health Check del Backend |
| **¿Prueba de despliegue?** | Sí |
| **Descripción** | Verificar que el endpoint de health responde correctamente |
| **Prerrequisitos** | - Backend desplegado |
| **Pasos** | 1. Ejecutar: `curl https://payvue-api.onrender.com/health` |
| **Resultado esperado** | - Status 200 OK<br>- Respuesta: "OK - PayVue API Server" |
| **Resultado obtenido** | ✅ EXITOSO - Health check responde |

---

### CP-017: Prueba de Rendimiento - Tiempo de Respuesta

| Campo | Valor |
|-------|-------|
| **Código** | CP-017 |
| **Nombre** | Verificar Tiempo de Respuesta de API |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que las APIs responden en menos de 500ms |
| **Prerrequisitos** | - Backend desplegado y activo |
| **Pasos** | 1. Realizar peticiones GET a endpoints principales<br>2. Medir tiempo de respuesta |
| **Resultado esperado** | - Tiempo de respuesta < 500ms |
| **Resultado obtenido** | ✅ EXITOSO - Promedio 200ms (después de warmup) |

---

### CP-018: Prueba de Seguridad - Inyección SQL

| Campo | Valor |
|-------|-------|
| **Código** | CP-018 |
| **Nombre** | Verificar Protección contra Inyección SQL |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que el sistema está protegido contra inyección SQL |
| **Prerrequisitos** | - Sistema desplegado |
| **Pasos** | 1. Intentar login con: `email: "'; DROP TABLE users;--"`<br>2. Verificar comportamiento |
| **Resultado esperado** | - Sistema rechaza el input<br>- No hay error de base de datos<br>- Tablas intactas |
| **Resultado obtenido** | ✅ EXITOSO - Queries parametrizadas protegen el sistema |

---

### CP-019: Prueba de Responsividad

| Campo | Valor |
|-------|-------|
| **Código** | CP-019 |
| **Nombre** | Verificar Diseño Responsive |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que la interfaz se adapta a diferentes tamaños de pantalla |
| **Prerrequisitos** | - Frontend desplegado |
| **Pasos** | 1. Abrir en desktop (1920x1080)<br>2. Abrir en tablet (768x1024)<br>3. Abrir en móvil (375x667) |
| **Resultado esperado** | - Interfaz usable en todos los tamaños<br>- Sidebar se oculta en móvil<br>- Elementos se reorganizan |
| **Resultado obtenido** | ✅ EXITOSO - Diseño responsive funciona |

---

### CP-020: Prueba de Dashboard Vacío

| Campo | Valor |
|-------|-------|
| **Código** | CP-020 |
| **Nombre** | Verificar Dashboard sin Datos |
| **¿Prueba de despliegue?** | No |
| **Descripción** | Verificar que el dashboard muestra mensaje apropiado cuando no hay datos |
| **Prerrequisitos** | - Usuario nuevo sin datos |
| **Pasos** | 1. Registrar nuevo usuario<br>2. Ir al Dashboard |
| **Resultado esperado** | - Mensaje "No hay datos todavía"<br>- Estadísticas en $0<br>- No hay errores |
| **Resultado obtenido** | ✅ EXITOSO - Mensaje mostrado correctamente |

---

## Estrategia de Ejecución de Pruebas

### Ciclos de Prueba

| Ciclo | Descripción | Casos de Prueba |
|-------|-------------|-----------------|
| Ciclo 1 - Smoke Test | Pruebas básicas de despliegue | CP-001, CP-004, CP-006, CP-008, CP-016 |
| Ciclo 2 - Funcionalidad Core | Pruebas de funcionalidades principales | CP-001 a CP-015 |
| Ciclo 3 - Validaciones | Pruebas de validación y errores | CP-002, CP-003, CP-005, CP-007 |
| Ciclo 4 - No Funcionales | Pruebas de rendimiento y seguridad | CP-017, CP-018, CP-019 |
| Ciclo 5 - Regresión | Pruebas completas antes de release | CP-001 a CP-020 |

### Matriz de Ejecución

| Caso de Prueba | Ciclo 1 | Ciclo 2 | Ciclo 3 | Ciclo 4 | Ciclo 5 |
|----------------|---------|---------|---------|---------|---------|
| CP-001 | ✅ | ✅ | | | ✅ |
| CP-002 | | | ✅ | | ✅ |
| CP-003 | | | ✅ | | ✅ |
| CP-004 | ✅ | ✅ | | | ✅ |
| CP-005 | | | ✅ | | ✅ |
| CP-006 | ✅ | ✅ | | | ✅ |
| CP-007 | | | ✅ | | ✅ |
| CP-008 | ✅ | ✅ | | | ✅ |
| CP-009 | | ✅ | | | ✅ |
| CP-010 | | ✅ | | | ✅ |
| CP-011 | | ✅ | | | ✅ |
| CP-012 | | ✅ | | | ✅ |
| CP-013 | | ✅ | | | ✅ |
| CP-014 | | ✅ | | | ✅ |
| CP-015 | | ✅ | | | ✅ |
| CP-016 | ✅ | | | | ✅ |
| CP-017 | | | | ✅ | ✅ |
| CP-018 | | | | ✅ | ✅ |
| CP-019 | | | | ✅ | ✅ |
| CP-020 | | ✅ | | | ✅ |

---

## Dependencias y Riesgos

### Riesgos Identificados

| ID | Riesgo | Probabilidad | Impacto | Plan de Mitigación |
|----|--------|--------------|---------|-------------------|
| R-001 | Servidor de Render en "sleep" | Alta | Medio | Realizar warmup antes de pruebas |
| R-002 | Pérdida de datos por falta de disco persistente | Alta | Alto | Documentar limitación, usar BD externa para producción |
| R-003 | Cambios en API durante pruebas | Media | Alto | Congelar desarrollo durante ciclo de pruebas |
| R-004 | Falta de disponibilidad del ambiente | Baja | Alto | Tener ambiente local como backup |
| R-005 | Tiempo insuficiente para pruebas | Media | Medio | Priorizar pruebas críticas |

### Dependencias

- Despliegue exitoso en Render (backend)
- Despliegue exitoso en Vercel (frontend)
- Conectividad a internet estable
- Variables de entorno configuradas correctamente

---

## Resumen de Resultados

### Estadísticas Generales

| Métrica | Valor |
|---------|-------|
| Total de Casos de Prueba | 20 |
| Casos Ejecutados | 20 |
| Casos Exitosos | 20 |
| Casos Fallidos | 0 |
| Porcentaje de Éxito | 100% |

### Defectos Encontrados y Corregidos

| ID | Descripción | Severidad | Estado |
|----|-------------|-----------|--------|
| DEF-001 | Variable no usada en Dashboard.js | Baja | ✅ Corregido |
| DEF-002 | Filtro de fecha no funcionaba en pagos | Media | ✅ Corregido |
| DEF-003 | Deuda mostraba "Desconocido" en pagos | Media | ✅ Corregido |

---

## Conclusiones

El sistema PayVue ha sido sometido a un proceso exhaustivo de pruebas de software que incluyó:

1. **Pruebas Unitarias**: Validación de componentes individuales tanto en backend como frontend.

2. **Pruebas de Integración**: Verificación de la comunicación entre frontend y backend, y entre los diferentes módulos del sistema.

3. **Pruebas de Sistema**: Validación de flujos completos de usuario end-to-end.

4. **Pruebas de Aceptación**: Confirmación de que el sistema cumple con los requisitos establecidos.

5. **Pruebas No Funcionales**: Verificación de rendimiento, seguridad y usabilidad.

### Resultados

- **100%** de los casos de prueba ejecutados exitosamente
- **0** defectos críticos pendientes
- Sistema listo para despliegue en producción

### Recomendaciones

1. Implementar base de datos externa para persistencia en producción
2. Agregar pruebas automatizadas en pipeline de CI/CD
3. Considerar implementación de autenticación con tokens JWT
4. Realizar pruebas de carga con mayor volumen de datos

---

## Referencias

1. IEEE 829 - Standard for Software Test Documentation
2. ISTQB - International Software Testing Qualifications Board
3. Documentación de Go Testing: https://golang.org/pkg/testing/
4. React Testing Library: https://testing-library.com/docs/react-testing-library/intro/
5. Postman Documentation: https://learning.postman.com/docs/

---

**Documento preparado por:** Juan Miguel Valencia Atehortua & Juan Andres Forero Guauque  
**Fecha:** 10/12/2024  
**Versión:** 1.0

