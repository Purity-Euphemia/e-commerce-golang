# E-Commerce API

A robust Go-based e-commerce backend built with Gin and GORM.

## Features

- ✅ User authentication with JWT tokens
- ✅ Password hashing with bcrypt
- ✅ Product management
- ✅ Shopping cart functionality
- ✅ Order processing
- ✅ CORS enabled for frontend integration
- ✅ Role-based access control (Admin/Customer)
- ✅ Input validation
- ✅ Consistent API responses

## Setup Instructions

### Prerequisites

- Go 1.25.4 or higher
- SQLite3 (or modify config for another database)

### Installation

1. **Clone and navigate to project**
```bash
cd ecommerce-go
```

2. **Install dependencies**
```bash
go mod download
go mod tidy
```

3. **Configure environment**
Create a `.env` file in the project root:
```
JWT_SECRET=your-secret-key-change-this-in-production
DATABASE_URL=ecommerce.db
GIN_MODE=debug
```

4. **Run the server**
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Documentation

### Base URL
```
http://localhost:8080
```

### Authentication
Include JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

---

## Endpoints

### User Management

#### Register User
```
POST /register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "customer",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### Login
```
POST /login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "customer",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

### Products

#### Get All Products
```
GET /products
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Laptop",
      "price": 999.99,
      "stock": 10
    }
  ]
}
```

#### Create Product (Admin Only)
```
POST /products
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Laptop",
  "price": 999.99,
  "stock": 10
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Laptop",
    "price": 999.99,
    "stock": 10
  }
}
```

---

### Shopping Cart

#### Add to Cart
```
POST /cart
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_id": 1,
  "quantity": 2
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "message": "Product added to cart"
  }
}
```

#### Update Cart Item
```
PUT /cart
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_id": 1,
  "quantity": 5
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "message": "cart updated"
  }
}
```

#### Remove from Cart
```
DELETE /cart/:product_id
Authorization: Bearer <token>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "message": "item removed"
  }
}
```

---

### Orders

#### Checkout
```
POST /checkout
Authorization: Bearer <token>
Content-Type: application/json

{
  "total": 1999.98
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "total_amount": 1999.98,
    "status": "pending"
  }
}
```

#### Get My Orders
```
GET /my-orders
Authorization: Bearer <token>
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "total_amount": 1999.98,
      "status": "pending"
    }
  ]
}
```

---

## Error Handling

All error responses follow this format:

```json
{
  "success": false,
  "error": "Error message"
}
```

### Common Status Codes
- `200` - OK
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden (Admin access required)
- `500` - Internal Server Error

---

## Database Schema

### Users Table
- `id` (PK)
- `name`
- `email` (UNIQUE)
- `password` (hashed)
- `role` (customer/admin)
- `created_at`
- `updated_at`

### Products Table
- `id` (PK)
- `name`
- `price`
- `stock`
- `created_at`
- `updated_at`

### Carts Table
- `id` (PK)
- `user_id` (FK)
- `created_at`
- `updated_at`

### CartItems Table
- `id` (PK)
- `cart_id` (FK)
- `product_id` (FK)
- `quantity`
- `created_at`
- `updated_at`

### Orders Table
- `id` (PK)
- `user_id` (FK)
- `total_amount`
- `status`
- `created_at`
- `updated_at`

### OrderItems Table
- `id` (PK)
- `order_id` (FK)
- `product_id` (FK)
- `quantity`
- `price_at_purchase`
- `created_at`
- `updated_at`

---

## Security Notes

1. **JWT Secret**: Change the `JWT_SECRET` in `.env` file in production
2. **Passwords**: All passwords are hashed using bcrypt with DefaultCost
3. **CORS**: Currently allows all origins. Configure in `main.go` for production
4. **Admin Role**: Only users with `admin` role can create products

---

## Development

### Project Structure
```
├── main.go                 # Entry point
├── database/              # Database configuration
├── models/                # Data models
├── repositories/          # Data access layer
├── services/              # Business logic
├── controllers/           # HTTP handlers
├── middleware/            # Custom middleware
├── routes/                # Route definitions
├── utils/                 # Utility functions
├── exceptions/            # Error definitions
└── .env                   # Environment variables
```

---

## Future Enhancements

- [ ] Email notifications
- [ ] Payment gateway integration (Stripe)
- [ ] Product reviews and ratings
- [ ] Wishlist functionality
- [ ] Order tracking
- [ ] Admin dashboard
- [ ] Search and filtering
- [ ] Pagination

---

## License

MIT
