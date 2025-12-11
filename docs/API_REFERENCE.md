# API Reference

## Descripción General

PayVue expone una API REST para gestionar datos financieros personales.

**Base URL:** `https://payvue-api.onrender.com`

**Content-Type:** `application/json`

---

## Autenticación

Todas las rutas (excepto `/auth/*` y `/health`) requieren el header `X-User-ID`.

```
X-User-ID: 1
```

---

## Endpoints

### Health Check

```http
GET /health
```

**Response 200:**
```
OK - PayVue API Server
```

---

## Autenticación

### Registrar Usuario

```http
POST /auth/register
```

**Request Body:**
```json
{
  "email": "usuario@email.com",
  "password": "123456"
}
```

**Response 201:**
```json
{
  "message": "Usuario registrado exitosamente"
}
```

**Errores:**
| Código | Error | Descripción |
|--------|-------|-------------|
| 400 | validation_error | Email o password inválidos |
| 400 | email_already_exists | Email ya registrado |

---

### Iniciar Sesión

```http
POST /auth/login
```

**Request Body:**
```json
{
  "email": "usuario@email.com",
  "password": "123456"
}
```

**Response 200:**
```json
{
  "message": "Inicio de sesión exitoso",
  "user_id": 1
}
```

**Errores:**
| Código | Error | Descripción |
|--------|-------|-------------|
| 401 | invalid_credentials | Email o password incorrectos |

---

### Cerrar Sesión

```http
POST /auth/logout
```

**Response 200:**
```json
{
  "message": "Sesión cerrada exitosamente"
}
```

---

## Ingresos

### Obtener Todos los Ingresos

```http
GET /finances/income
```

**Headers:**
```
X-User-ID: 1
```

**Response 200:**
```json
[
  {
    "id": 1,
    "amount": 5000000,
    "source": "Salario",
    "date": "2024-12-01"
  },
  {
    "id": 2,
    "amount": 500000,
    "source": "Freelance",
    "date": "2024-12-05"
  }
]
```

---

### Obtener Ingreso por ID

```http
GET /finances/income/{id}
```

**Response 200:**
```json
{
  "id": 1,
  "amount": 5000000,
  "source": "Salario",
  "date": "2024-12-01"
}
```

**Errores:**
| Código | Error | Descripción |
|--------|-------|-------------|
| 404 | income_not_found | Ingreso no encontrado |

---

### Crear Ingreso

```http
POST /finances/income
```

**Request Body:**
```json
{
  "amount": 5000000,
  "source": "Salario",
  "date": "2024-12-01"
}
```

**Response 201:**
```json
{
  "id": 1,
  "amount": 5000000,
  "source": "Salario",
  "date": "2024-12-01"
}
```

---

### Actualizar Ingreso

```http
PUT /finances/income/{id}
```

**Request Body:**
```json
{
  "amount": 5500000,
  "source": "Salario + Bono",
  "date": "2024-12-01"
}
```

**Response 200:**
```json
{
  "id": 1,
  "amount": 5500000,
  "source": "Salario + Bono",
  "date": "2024-12-01"
}
```

---

### Eliminar Ingreso

```http
DELETE /finances/income/{id}
```

**Response 200:**
```json
{
  "message": "Ingreso eliminado exitosamente"
}
```

---

## Deudas

### Obtener Todas las Deudas

```http
GET /finances/debt
```

**Response 200:**
```json
[
  {
    "id": 1,
    "name": "Crédito Bancario",
    "total_amount": 10000000,
    "remaining_amount": 8000000,
    "due_date": "2025-12-01",
    "interest_rate": 1.5,
    "num_installments": 24,
    "installment_amount": 450000,
    "payment_day": 15,
    "remaining_payments": 18,
    "paid": false
  }
]
```

---

### Obtener Deuda por ID

```http
GET /finances/debt/{id}
```

**Response 200:**
```json
{
  "id": 1,
  "name": "Crédito Bancario",
  "total_amount": 10000000,
  "remaining_amount": 8000000,
  "due_date": "2025-12-01",
  "interest_rate": 1.5,
  "num_installments": 24,
  "installment_amount": 450000,
  "payment_day": 15,
  "remaining_payments": 18,
  "paid": false
}
```

---

### Crear Deuda

```http
POST /finances/debt
```

**Request Body:**
```json
{
  "name": "Crédito Bancario",
  "total_amount": 10000000,
  "remaining_amount": 10000000,
  "due_date": "2025-12-01",
  "interest_rate": 1.5,
  "num_installments": 24,
  "installment_amount": 450000,
  "payment_day": 15
}
```

**Response 201:**
```json
{
  "id": 1,
  "name": "Crédito Bancario",
  "total_amount": 10000000,
  "remaining_amount": 10000000,
  "due_date": "2025-12-01",
  "interest_rate": 1.5,
  "num_installments": 24,
  "installment_amount": 450000,
  "payment_day": 15,
  "remaining_payments": 24,
  "paid": false
}
```

---

### Actualizar Deuda

```http
PUT /finances/debt/{id}
```

**Request Body:**
```json
{
  "name": "Crédito Bancario",
  "total_amount": 10000000,
  "remaining_amount": 7500000,
  "due_date": "2025-12-01",
  "interest_rate": 1.5,
  "num_installments": 24,
  "installment_amount": 450000,
  "payment_day": 15,
  "paid": false
}
```

**Response 200:**
```json
{
  "id": 1,
  "name": "Crédito Bancario",
  "remaining_amount": 7500000,
  "remaining_payments": 17,
  "paid": false
}
```

---

### Eliminar Deuda

```http
DELETE /finances/debt/{id}
```

**Response 200:**
```json
{
  "message": "Deuda eliminada exitosamente"
}
```

> ⚠️ **Nota:** Eliminar una deuda también elimina todos los pagos asociados.

---

## Pagos

### Obtener Todos los Pagos

```http
GET /finances/payment
```

**Response 200:**
```json
[
  {
    "id": 1,
    "amount": 450000,
    "date": "2024-12-15",
    "debt_id": 1,
    "debt_name": "Crédito Bancario",
    "remaining_installments": 17,
    "remaining_amount": 7650000,
    "receipt_url": "/finances/payment/receipt/abc123.jpg"
  }
]
```

---

### Crear Pago

```http
POST /finances/payment
Content-Type: multipart/form-data
```

**Form Fields:**
| Campo | Tipo | Requerido | Descripción |
|-------|------|-----------|-------------|
| amount | number | ✅ | Monto del pago |
| debt_id | number | ✅ | ID de la deuda |
| date | string | ❌ | Fecha (YYYY-MM-DD) |
| receipt | file | ❌ | Imagen del recibo |

**Response 201:**
```json
{
  "id": 1,
  "amount": 450000,
  "date": "2024-12-15",
  "debt_name": "Crédito Bancario",
  "remaining_installments": 17,
  "remaining_amount": 7650000,
  "receipt_url": "/finances/payment/receipt/abc123.jpg"
}
```

---

### Ver Recibo

```http
GET /finances/payment/receipt/{filename}
```

**Response:** Archivo de imagen (JPEG, PNG)

---

### Eliminar Pago

```http
DELETE /finances/payment/{id}
```

**Response 200:**
```json
{
  "message": "Pago eliminado exitosamente"
}
```

---

## Códigos de Estado HTTP

| Código | Significado |
|--------|-------------|
| 200 | OK - Operación exitosa |
| 201 | Created - Recurso creado |
| 400 | Bad Request - Error de validación |
| 401 | Unauthorized - Sin autenticación |
| 404 | Not Found - Recurso no encontrado |
| 500 | Internal Server Error - Error del servidor |

---

## Ejemplos con cURL

### Registro

```bash
curl -X POST https://payvue-api.onrender.com/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"123456"}'
```

### Login

```bash
curl -X POST https://payvue-api.onrender.com/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@test.com","password":"123456"}'
```

### Crear Ingreso

```bash
curl -X POST https://payvue-api.onrender.com/finances/income \
  -H "Content-Type: application/json" \
  -H "X-User-ID: 1" \
  -d '{"amount":5000000,"source":"Salario","date":"2024-12-01"}'
```

### Crear Pago con Recibo

```bash
curl -X POST https://payvue-api.onrender.com/finances/payment \
  -H "X-User-ID: 1" \
  -F "amount=450000" \
  -F "debt_id=1" \
  -F "date=2024-12-15" \
  -F "receipt=@/path/to/receipt.jpg"
```

---

## Rate Limiting

Actualmente no hay rate limiting implementado (Free Tier). En producción se recomienda:
- 100 requests/minuto por IP
- 1000 requests/hora por usuario

