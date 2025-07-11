# External API Client

Aplikasi Go yang mengintegrasikan dengan 3 endpoint eksternal untuk mengambil data negara, pelabuhan, dan barang.

## Fitur

- **GET /api/v1/negaras** - Mengambil semua data negara
- **GET /api/v1/pelabuhans?id_negara=2** - Mengambil data pelabuhan berdasarkan ID negara
- **GET /api/v1/barangs?id_pelabuhan=1** - Mengambil data barang berdasarkan ID pelabuhan
- **GET /health** - Health check endpoint

## Instalasi

1. **Clone atau download project ini**

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup environment variables**
   ```bash
   cp .env.example .env
   ```

4. **Jalankan aplikasi**
   ```bash
   go run .
   ```

Aplikasi akan berjalan di `http://localhost:8080`

## API Endpoints

### 1. Get All Countries
```bash
GET /api/v1/negaras
```

**Response Success:**
```json
{
  "status": "success",
  "message": "Countries data retrieved successfully",
  "data": [
    {
      "id": 1,
      "nama": "Indonesia",
      "kode": "ID",
      "created": "2024-01-01T00:00:00Z",
      "modified": "2024-01-01T00:00:00Z"
    }
  ]
}
```

**Response Empty:**
```json
{
  "status": "success",
  "message": "No countries found",
  "data": []
}
```

### 2. Get Ports by Country ID
```bash
GET /api/v1/pelabuhans?id_negara=2
```

**Response Success:**
```json
{
  "status": "success",
  "message": "Ports data retrieved successfully",
  "data": [
    {
      "id": 1,
      "nama": "Pelabuhan Jakarta",
      "kode": "JKT",
      "id_negara": 2,
      "created": "2024-01-01T00:00:00Z",
      "modified": "2024-01-01T00:00:00Z"
    }
  ]
}
```

**Response Empty:**
```json
{
  "status": "success",
  "message": "No ports found for the specified country",
  "data": []
}
```

**Response Error (Missing Parameter):**
```json
{
  "status": "error",
  "message": "Parameter id_negara is required"
}
```

### 3. Get Goods by Port ID
```bash
GET /api/v1/barangs?id_pelabuhan=1
```

**Response Success:**
```json
{
  "status": "success",
  "message": "Goods data retrieved successfully",
  "data": [
    {
      "id": 1,
      "nama": "Kopi",
      "kode": "KPI",
      "id_pelabuhan": 1,
      "harga": 50000,
      "stok": 100,
      "created": "2024-01-01T00:00:00Z",
      "modified": "2024-01-01T00:00:00Z"
    }
  ]
}
```

**Response Empty:**
```json
{
  "status": "success",
  "message": "No goods found for the specified port",
  "data": []
}
```

**Response Error (Missing Parameter):**
```json
{
  "status": "error",
  "message": "Parameter id_pelabuhan is required"
}
```

### 4. Health Check
```bash
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "message": "Service is running"
}
```

## Error Handling

Aplikasi menangani berbagai kondisi error:

- **Parameter tidak valid**: Mengembalikan error 400 dengan pesan yang jelas
- **API eksternal tidak tersedia**: Mengembalikan error 500 dengan detail error
- **Data kosong**: Mengembalikan response sukses dengan array kosong
- **Route tidak ditemukan**: Mengembalikan error 404

## Testing

Contoh testing menggunakan curl:

```bash
# Test get countries
curl http://localhost:8080/api/v1/negaras

# Test get ports by country ID
curl "http://localhost:8080/api/v1/pelabuhans?id_negara=2"

# Test get goods by port ID
curl "http://localhost:8080/api/v1/barangs?id_pelabuhan=1"

# Test health check
curl http://localhost:8080/health
```

## Struktur Project

```
.
├── main.go          # Entry point aplikasi
├── models.go        # Data structures dan models
├── service.go       # Business logic untuk API calls
├── controller.go    # HTTP handlers
├── routes.go        # Route definitions
├── go.mod          # Go modules
├── .env.example    # Environment variables template
└── README.md       # Dokumentasi
```

## Dependencies

- **Gin**: Web framework untuk Go
- **godotenv**: Untuk loading environment variables

## Environment Variables

- `PORT`: Port server (default: 8080)
- `EXTERNAL_API_BASE_URL`: Base URL untuk API eksternal
- `EXTERNAL_API_TIMEOUT`: Timeout untuk API calls
- `GIN_MODE`: Mode Gin (debug/release)