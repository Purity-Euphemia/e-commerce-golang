# Quick Setup Guide

## Prerequisites
- Go 1.25.4+
- Git

## Quick Start (5 minutes)

### 1. Install Dependencies
```bash
go mod download
go mod tidy
```

### 2. Configure Environment
```bash
# Create .env file with your settings (already created with defaults)
cat .env
```

### 3. Run the Server
```bash
go run main.go
```

You should see:
```
[GIN-debug] listening and serving HTTP on :8080
```

## Test the API

### 1. Register a User
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 2. Get the Token
Save the token from the response above.

### 3. Add a Product (Need to set role to admin first)
```bash
# First, register/login as admin (you'd need to manually update DB or create admin endpoint)
# For now, use the token you got from register

curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "name": "Laptop",
    "price": 999.99,
    "stock": 10
  }'
```

### 4. Get Products
```bash
curl http://localhost:8080/products
```

### 5. Add to Cart
```bash
curl -X POST http://localhost:8080/cart \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "product_id": 1,
    "quantity": 2
  }'
```

### 6. Checkout
```bash
curl -X POST http://localhost:8080/checkout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "total": 1999.98
  }'
```

## Frontend Integration

### CORS Support
✅ Enabled - Frontend can make requests from any origin

### Response Format
All endpoints return:
```json
{
  "success": true/false,
  "data": {...},
  "error": "error message (if any)"
}
```

### Authentication
Include JWT token in every protected request:
```
Authorization: Bearer <token>
```

## Troubleshooting

### Port 8080 Already in Use
```bash
go run main.go  # Will show error
# Change port in main.go (line: r.Run(":8080"))
```

### Database Issues
```bash
# Delete the database file to reset
rm ecommerce.db
go run main.go  # Will create new database
```

### Dependencies Not Found
```bash
go mod tidy
go mod download
```

## Key Files Modified/Created

1. ✅ **Fixed all imports** - Changed `config` to `database`
2. ✅ **Added password hashing** - Using bcrypt
3. ✅ **JWT environment variables** - Use .env file
4. ✅ **Added CORS** - CORS enabled for frontend
5. ✅ **Fixed model relationships** - Proper foreign keys
6. ✅ **Standardized responses** - Consistent format
7. ✅ **Added validation** - Input validation on all endpoints
8. ✅ **Created .env** - Environment configuration
9. ✅ **Created .gitignore** - For version control

## Next Steps

1. Create `.env` file with your JWT_SECRET
2. Run `go run main.go`
3. Test with curl or Postman
4. Integrate with your frontend
5. Deploy to production

## Support Files

- **API_DOCUMENTATION.md** - Complete API reference
- **.env** - Environment configuration
- **.gitignore** - Git ignore rules

Enjoy! 🚀
