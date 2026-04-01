# Complete E-Commerce API Documentation

**Version:** 2.0 (Complete Edition)  
**Status:** Production Ready ✅

## Table of Contents
1. [Overview](#overview)
2. [Authentication](#authentication)
3. [Products & Categories](#products--categories)
4. [Shopping Cart](#shopping-cart)    
5. [Orders & Checkout](#orders--checkout)
6. [Reviews & Ratings](#reviews--ratings)
7. [Wishlist](#wishlist)
8. [Coupons & Discounts](#coupons--discounts)
9. [User Profile](#user-profile)
10. [Admin Endpoints](#admin-endpoints)
11. [Error Handling](#error-handling)

---

## Overview

Complete and production-ready E-Commerce API built with Go, Gin, and GORM.

**Base URL:** `http://localhost:8080`

### Features
- ✅ User authentication with JWT
- ✅ Product catalog with categories
- ✅ Search and filtering
- ✅ Product reviews and ratings
- ✅ Shopping cart management
- ✅ Wishlist functionality
- ✅ Discount coupons
- ✅ Order management with tracking
- ✅ Admin dashboard
- ✅ Role-based access control

---

## Authentication

### Register
```
POST /register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "SecurePassword123"
}
```

**Response (201 Created):**
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

### Login
```
POST /login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "SecurePassword123"
}
```

**Response (200 OK):**
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

### Token Usage
Include in all protected requests:
```
Authorization: Bearer <token>
```

---

## Products & Categories

### Get All Categories
```
GET /categories
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Electronics",
      "slug": "electronics",
      "description": "Electronic devices and gadgets",
      "icon": "📱"
    }
  ]
}
```

### Get Category By Slug
```
GET /categories/:slug
```

Example: `GET /categories/electronics`

### Create Category (Admin)
```
POST /admin/categories
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Electronics",
  "description": "Electronic devices and gadgets",
  "icon": "📱"
}
```

### Get All Products
```
GET /products?page=1&page_size=10
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Laptop",
      "slug": "laptop-hp-15",
      "description": "High performance laptop",
      "price": 999.99,
      "discount_price": 799.99,
      "stock": 15,
      "category_id": 1,
      "image": "https://...",
      "sku": "LAPTOP001",
      "rating": 4.5
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 50
  }
}
```

### Get Product By ID
```
GET /products/:id
```

Example: `GET /products/1`

### Search Products
```
GET /products/search?q=laptop&page=1&page_size=10
```

### Get Products By Category
```
GET /products/category/:category_id?page=1&page_size=10
```

Example: `GET /products/category/1?page=1&page_size=10`

### Create Product (Admin)
```
POST /admin/products
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Laptop",
  "slug": "laptop-hp-15",
  "description": "High performance laptop",
  "price": 999.99,
  "discount_price": 799.99,
  "stock": 15,
  "category_id": 1,
  "image": "https://...",
  "sku": "LAPTOP001"
}
```

### Update Product (Admin)
```
PUT /admin/products/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Laptop Updated",
  "price": 949.99,
  "stock": 20,
  ...
}
```

### Delete Product (Admin)
```
DELETE /admin/products/:id
Authorization: Bearer <token>
```

---

## Shopping Cart

### Add to Cart
```
POST /cart
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_id": 1,
  "quantity": 2
}
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "message": "Product added to cart"
  }
}
```

### Update Cart Item
```
PUT /cart
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_id": 1,
  "quantity": 5
}
```

### Remove from Cart
```
DELETE /cart/:product_id
Authorization: Bearer <token>
```

---

## Orders & Checkout

### Checkout
```
POST /checkout
Authorization: Bearer <token>
Content-Type: application/json

{
  "total": 1999.98
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "order_number": "ORD-2026-001",
    "total_amount": 1999.98,
    "discount_amount": 100,
    "coupon_code": "SUMMER20",
    "status": "pending",
    "payment_status": "pending",
    "items": [...]
  }
}
```

### Get My Orders
```
GET /my-orders
Authorization: Bearer <token>
```

### Get All Orders (Admin)
```
GET /admin/orders?page=1&page_size=10
Authorization: Bearer <token>
```

### Update Order Status (Admin)
```
PUT /admin/orders/:order_id/status
Authorization: Bearer <token>
Content-Type: application/json

{
  "status": "shipped"
}
```

**Valid statuses:** `pending`, `confirmed`, `shipped`, `delivered`, `cancelled`

---

## Reviews & Ratings

### Add Review
```
POST /products/:product_id/reviews
Authorization: Bearer <token>
Content-Type: application/json

{
  "rating": 5,
  "title": "Excellent Product!",
  "comment": "Amazing quality and fast shipping"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "product_id": 1,
    "user_id": 1,
    "rating": 5,
    "title": "Excellent Product!",
    "comment": "Amazing quality and fast shipping",
    "helpful": 0
  }
}
```

### Get Product Reviews
```
GET /products/:product_id/reviews?page=1&page_size=10
```

### Delete Review
```
DELETE /reviews/:review_id
Authorization: Bearer <token>
```

---

## Wishlist

### Add to Wishlist (Toggle)
```
POST /wishlist
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_id": 1
}
```

### Get My Wishlist
```
GET /my-wishlist
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "product_id": 1,
      "product": {
        "id": 1,
        "name": "Laptop",
        "price": 999.99,
        ...
      }
    }
  ]
}
```

### Remove from Wishlist
```
DELETE /wishlist/:product_id
Authorization: Bearer <token>
```

---

## Coupons & Discounts

### Create Coupon (Admin)
```
POST /admin/coupons
Authorization: Bearer <token>
Content-Type: application/json

{
  "code": "SUMMER20",
  "discount_type": "percentage",
  "discount_value": 20,
  "min_amount": 50,
  "max_usage": 100
}
```

**Coupon Types:**
- `percentage`: Discount as percentage (e.g., 20%)
- `fixed`: Fixed amount discount (e.g., $20)

### Get All Coupons (Admin)
```
GET /admin/coupons
Authorization: Bearer <token>
```

### Delete Coupon (Admin)
```
DELETE /admin/coupons/:code
Authorization: Bearer <token>
```

### Apply Coupon at Checkout
Include `coupon_code` in checkout request:
```json
{
  "total": 1999.98,
  "coupon_code": "SUMMER20"
}
```

---

## User Profile

### Get Profile
```
GET /profile
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1-234-567-8900",
    "avatar": "https://...",
    "address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip_code": "10001"
  }
}
```

### Update Profile
```
PUT /profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "John Doe",
  "phone": "+1-234-567-8900",
  "avatar": "https://...",
  "address": "123 Main St",
  "city": "New York",
  "state": "NY",
  "zip_code": "10001"
}
```

---

## Admin Endpoints

### Dashboard Statistics
```
GET /admin/dashboard/stats
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "total_users": 150,
    "total_orders": 500,
    "total_products": 50,
    "total_revenue": 125000.50,
    "today_revenue": 3200.75,
    "pending_orders": 12
  }
}
```

### Manage Products
```
POST /admin/products              # Create
PUT /admin/products/:id           # Update
DELETE /admin/products/:id        # Delete
```

### Manage Categories
```
POST /admin/categories            # Create
```

### Manage Orders
```
GET /admin/orders                 # List all
PUT /admin/orders/:order_id/status # Update status
```

### Manage Coupons
```
POST /admin/coupons               # Create
GET /admin/coupons                # List all
DELETE /admin/coupons/:code       # Delete
```

---

## Error Handling

### Error Response Format
```json
{
  "success": false,
  "error": "Error message describing the issue"
}
```

### Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK - Request successful |
| 201 | Created - Resource created |
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Missing/invalid token |
| 403 | Forbidden - Insufficient permissions |
| 404 | Not Found - Resource doesn't exist |
| 500 | Internal Server Error |

### Common Errors

**Missing Token:**
```json
{
  "success": false,
  "error": "missing token"
}
```

**Invalid Token:**
```json
{
  "success": false,
  "error": "invalid token"
}
```

**Admin Access Only:**
```json
{
  "success": false,
  "error": "admin access only"
}
```

**Insufficient Stock:**
```json
{
  "success": false,
  "error": "not enough stock"
}
```

---

## Pagination

All list endpoints support pagination:

```
GET /products?page=1&page_size=10
```

**Parameters:**
- `page` (default: 1) - Page number (1-based)
- `page_size` (default: 10, max: 100) - Items per page

**Response includes:**
```json
{
  "pagination": {
    "page": 1,
    "page_size": 10,
    "total": 150
  }
}
```

---

## Best Practices

### Security
1. Always use HTTPS in production
2. Never expose JWT secrets
3. Implement rate limiting
4. Validate all inputs
5. Use strong passwords (min 6 chars)

### Performance
1. Use pagination for list endpoints
2. Filter and search appropriately
3. Cache frequently accessed data
4. Use appropriate indexes

### API Usage
1. Always include Content-Type header
2. Include Authorization token in protected requests
3. Handle errors gracefully
4. Implement retry logic for failed requests

---

## Example: Complete Checkout Flow

```bash
# 1. Register
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com","password":"pass123"}'

# Save token from response

# 2. Get products
curl http://localhost:8080/products

# 3. Add to cart
curl -X POST http://localhost:8080/cart \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"product_id":1,"quantity":2}'

# 4. Apply coupon (optional)
# Include "coupon_code" in checkout

# 5. Checkout
curl -X POST http://localhost:8080/checkout \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"total":1999.98,"coupon_code":"SUMMER20"}'

# 6. View order
curl http://localhost:8080/my-orders \
  -H "Authorization: Bearer TOKEN"
```

---

## Contact & Support
For issues or questions, please contact the development team.
