# Product Catalog Backend API

A robust and secure RESTful API backend for a product catalog application built with Go, Gin, GORM, and PostgreSQL. It includes secure user authentication using JWT (JSON Web Tokens) and password hashing with bcrypt.

## 🚀 Tech Stack

Language: Go (Golang)

Web Framework: Gin

ORM: GORM

Database: PostgreSQL

Authentication: golang-jwt/v5 (Stateless JWT)

Security: bcrypt (Password Hashing)

Database Migration: GORM AutoMigrate

## 📁 Project Structure

```text
product-catalog/
├── config/         # Database and environment configurations
├── controllers/    # Request handlers (Auth, Product CRUD)
├── models/         # Database models (User, Product structs)
├── routes/         # API endpoint routing setup
├── middlewares/    # JWT Authentication middleware
├── .env            # Environment variables (local)
├── go.mod          # Go modules dependencies
├── go.sum          # Go modules checksums
└── main.go         # Application entry point
```

## ⚙️ Installation & Setup

Follow these steps to run the project locally on your machine.

### 1. Prerequisites

Make sure the following software is installed:

- Go
- PostgreSQL
- Git

### 2. Clone the Repository

```bash
git clone git@github.com:your-username/product-catalog.git
cd product-catalog
```

### 3. Install Dependencies

Download and install all required Go packages:

```bash
go mod tidy
```

### 4. Configure Environment Variables

Create a `.env` file in the root directory of the project:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_db_password
DB_NAME=product_db
DB_PORT=5432
JWT_SECRET=your_super_secret_key_here
```

Important: Never commit your `.env` file or expose your JWT_SECRET in a public repository.

### 5. Run the Application

Start the Go server:

```bash
go run main.go
```

The server will automatically run database migrations using GORM and start listening on:

```text
http://localhost:8080
```

## 🔌 API Endpoints

### 1. Authentication

#### Register a New User

Endpoint:

```http
POST /api/register
```

Request Body:

```json
{
  "username": "admin",
  "password": "yourpassword"
}
```

#### Login

Endpoint:

```http
POST /api/login
```

Request Body:

```json
{
  "username": "admin",
  "password": "yourpassword"
}
```

The login endpoint returns a JWT token that can be used to access protected endpoints.

### 2. Products

Product endpoints are protected and require a valid JWT token.

Add the following HTTP header to protected requests:

```http
Authorization: Bearer <token>
```

#### Get All Products

```http
GET /api/products
```

#### Create a Product

```http
POST /api/products
```

Authorization: Required

```http
Authorization: Bearer <token>
```

#### Get Product by ID

```http
GET /api/products/:id
```

#### Update Product

```http
PUT /api/products/:id
```

Authorization: Required

```http
Authorization: Bearer <token>
```

#### Delete Product

```http
DELETE /api/products/:id
```

Authorization: Required

```http
Authorization: Bearer <token>
```

## 🔐 Authentication Flow

The authentication flow works as follows:

1. User registers with a username and password.
2. The password is securely hashed using bcrypt before being stored.
3. User logs in with their credentials.
4. The server validates the username and password.
5. If valid, the server generates a JWT token.
6. The client sends the JWT token using the Authorization header.
7. JWT middleware validates the token before allowing access to protected endpoints.

Example:

```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...
```

## 🗄️ Database

This project uses PostgreSQL as the database and GORM as the ORM.

GORM automatically runs migrations when the application starts, ensuring the required database tables are created or updated based on the defined models.

Make sure the PostgreSQL database specified in `.env` already exists:

```text
product_db
```

## 🛡️ Security

This project implements several security practices:

- Passwords are hashed using bcrypt.
- Authentication uses stateless JWT tokens.
- Protected routes require a valid JWT.
- Database credentials are stored in environment variables.
- JWT secrets are stored in environment variables.
- `.env` should not be committed to version control.

Add `.env` to `.gitignore`:

```gitignore
.env
```

## 📌 Feature Status (Technical Assignment Requirements)

### 1. YANG SUDAH SELESAI (DONE)

- **Tema Data (Katalog Produk):** Menggunakan Tema B (Nama, Deskripsi, Harga, Stok).
- **Backend Framework:** REST API menggunakan Go/Gin.
- **Frontend Stack:** React + Vite + TypeScript + Tailwind CSS (consume API sendiri).
- **Auth (Login & Token):** Login berbasis JWT token sudah berjalan.
- **Protected Routes (Backend & Frontend):** Endpoint CRUD di backend dilindungi middleware JWT, dan di frontend rute /dashboard memvalidasi token.
- **CRUD Operations (Create, Read List, Update, Delete):** Berhasil diintegrasikan antara React dan Go.
- **Database & Seeder:** Memakai PostgreSQL/MySQL dengan GORM dan sudah membuat seed data awal untuk admin (database/seed.go).
- **State Management Frontend:** Menangani state loading, error handling, dan redirect ke halaman login jika token hilang/invalid.
- **CORS Handling:** Sudah mengizinkan method (OPTIONS, PUT, DELETE, dll.) dan header Authorization.

### 2. YANG BELUM / PERLU DILENGKAPI (TODO)

#### A. Fitur Teknis di Kode

**Fitur Register (Pendaftaran Akun Baru):**

- **Status:** Belum ada. Dokumen meminta Register & Login.
- **Solusi Sederhana:** Tambahkan endpoint POST /api/register di backend dan tombol/halaman/modal register sederhana di frontend.

**Read Detail per Item:**

- **Status:** Baru ada Read List (Daftar Produk).
- **Solusi Sederhana:** Tambahkan endpoint GET /api/products/:id di backend. Di frontend, bisa ditampilkan lewat Modal Detail saat kartu produk diklik.

**Validasi Input & Kategori Produk (Tema B):**

- **Status:** Sebagian field sudah ada, tapi pastikan validasi backend membalikkan status code & pesan error yang sesuai jika input kosong. Tambahkan field kategori pada skema produk jika belum ada.

## 🔑 Testing Credentials

Default seed account available for testing:

- **Username:** admin
- **Password:** passwordBaru123
