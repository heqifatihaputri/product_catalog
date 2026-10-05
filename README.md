Product Catalog Backend API

A robust and secure RESTful API backend for a product catalog application built with Go, Gin, GORM, and PostgreSQL. It includes secure user authentication using JWT (JSON Web Tokens) and password hashing with bcrypt.

🚀 Tech Stack

Language: Go (Golang)

Web Framework: Gin

ORM: GORM

Database: PostgreSQL

Authentication: golang-jwt/v5 (Stateless JWT)

Security: bcrypt (Password Hashing)

Database Migration: GORM AutoMigrate

📁 Project Structure
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

⚙️ Installation & Setup

Follow these steps to run the project locally on your machine.

1. Prerequisites

Make sure the following software is installed:

Go

PostgreSQL

Git

2. Clone the Repository
git clone git@github.com:your-username/product-catalog.git
cd product-catalog

3. Install Dependencies

Download and install all required Go packages:

go mod tidy

4. Configure Environment Variables

Create a .env file in the root directory of the project:

DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_db_password
DB_NAME=product_db
DB_PORT=5432
JWT_SECRET=your_super_secret_key_here


Important: Never commit your .env file or expose your JWT_SECRET in a public repository.

5. Run the Application

Start the Go server:

go run main.go


The server will automatically run database migrations using GORM and start listening on:

http://localhost:8080

🔌 API Endpoints
1. Authentication
Register a New User

Endpoint:

POST /api/register


Request Body:

{
  "username": "admin",
  "password": "yourpassword"
}

Login

Endpoint:

POST /api/login


Request Body:

{
  "username": "admin",
  "password": "yourpassword"
}


The login endpoint returns a JWT token that can be used to access protected endpoints.

2. Products

Product endpoints are protected and require a valid JWT token.

Add the following HTTP header to protected requests:

Authorization: Bearer <token>

Get All Products
GET /api/products

Create a Product
POST /api/products


Authorization: Required

Authorization: Bearer <token>

Get Product by ID
GET /api/products/:id

Update Product
PUT /api/products/:id


Authorization: Required

Authorization: Bearer <token>

Delete Product
DELETE /api/products/:id


Authorization: Required

Authorization: Bearer <token>

🔐 Authentication Flow

The authentication flow works as follows:

User registers with a username and password.

The password is securely hashed using bcrypt before being stored.

User logs in with their credentials.

The server validates the username and password.

If valid, the server generates a JWT token.

The client sends the JWT token using the Authorization header.

JWT middleware validates the token before allowing access to protected endpoints.

Example:

Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...

🗄️ Database

This project uses PostgreSQL as the database and GORM as the ORM.

GORM automatically runs migrations when the application starts, ensuring the required database tables are created or updated based on the defined models.

Make sure the PostgreSQL database specified in .env already exists:

product_db

🛡️ Security

This project implements several security practices:

Passwords are hashed using bcrypt.

Authentication uses stateless JWT tokens.

Protected routes require a valid JWT.

Database credentials are stored in environment variables.

JWT secrets are stored in environment variables.

.env should not be committed to version control.

Add .env to .gitignore:

.env
