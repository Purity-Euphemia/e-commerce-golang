# Deployment & Deployment Guide

## Local Development

### Quick Start

```bash
# 1. Clone/Navigate to project
cd ecommerce-go

# 2. Install dependencies
go mod download
go mod tidy

# 3. Setup environment
cp .env.example .env

# 4. Update .env with your config
nano .env

# 5. Run the server
go run main.go
```

Server will start on `http://localhost:8080`

---

## Environment Configuration

### Development (.env)
```
JWT_SECRET=dev-secret-key-change-in-production
DATABASE_URL=ecommerce.db
GIN_MODE=debug
PORT=8080
```

### Production (.env)
```
JWT_SECRET=very-long-random-secret-key-with-special-chars-min-32-chars
DATABASE_URL=/var/lib/ecommerce/data.db
GIN_MODE=release
PORT=:8080
```

---

## Database Setup

### SQLite (Default)
Works out of the box. Database file created automatically.

### PostgreSQL Migration

1. **Update go.mod** - Change driver:
```bash
go get gorm.io/driver/postgres
```

2. **Update database/config.go**:
```go
import "gorm.io/driver/postgres"

dsn := "host=localhost user=ecommerce password=pass dbname=ecommerce port=5432 sslmode=disable"
database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

3. **Create database**:
```bash
createdb ecommerce
```

### MySQL Migration

1. **Update go.mod**:
```bash
go get gorm.io/driver/mysql
```

2. **Update database/config.go**:
```go
import "gorm.io/driver/mysql"

dsn := "ecommerce:password@tcp(localhost:3306)/ecommerce?charset=utf8mb4&parseTime=True"
database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

---

## Build for Production

### Single Binary
```bash
# Build standalone binary
go build -o ecommerce-api

# Run
./ecommerce-api
```

### Docker Deployment

**Dockerfile:**
```dockerfile
FROM golang:1.25.4-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ecommerce-api

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/ecommerce-api .
COPY .env .

EXPOSE 8080

CMD ["./ecommerce-api"]
```

**Build & Run:**
```bash
docker build -t ecommerce-api .
docker run -p 8080:8080 \
  -e JWT_SECRET=your-secret \
  -e DATABASE_URL=/data/ecommerce.db \
  -v ecommerce-data:/data \
  ecommerce-api
```

---

## Production Deployment Checklist

- [ ] Update JWT_SECRET to strong random string
- [ ] Switch to PostgreSQL/MySQL for database
- [ ] Enable HTTPS/SSL
- [ ] Setup backup strategy
- [ ] Configure logging
- [ ] Setup monitoring
- [ ] Configure rate limiting
- [ ] Setup error tracking (Sentry)
- [ ] Configure CDN for static files
- [ ] Setup domain name
- [ ] Configure payment gateway (Stripe)
- [ ] Setup email service
- [ ] Enable database backups
- [ ] Configure firewall rules
- [ ] Setup load balancer
- [ ] Monitor performance metrics

---

## Deployment Platforms

### Heroku
```bash
# Install Heroku CLI
heroku login

# Create app
heroku create ecommerce-api

# Set environment
heroku config:set JWT_SECRET=your-secret

# Deploy
git push heroku main
```

### AWS EC2
```bash
# SSH into instance
ssh -i key.pem ec2-user@your-instance

# Install Go
sudo yum install golang

# Clone repository
git clone your-repo
cd ecommerce-go

# Build and run
go build
./ecommerce-api
```

### Railway
```bash
# Install Railway CLI
railway login

# Deploy
railway up
```

### Render
```bash
# Connect GitHub repository
# Select Go project
# Set build command: go build
# Set start command: ./ecommerce-api
```

### Digital Ocean App Platform
```bash
# Create app.yaml
# Connect GitHub
# Deploy automatically on push
```

---

## Performance Optimization

### Caching
Add Redis for caching:
```bash
go get github.com/redis/go-redis/v9
```

### Database Optimization
- Add indexes on frequently searched fields
- Implement query pagination
- Use connection pooling

### Load Balancing
```nginx
upstream ecommerce_backend {
    server localhost:8080;
    server localhost:8081;
    server localhost:8082;
}

server {
    listen 80;
    server_name yourdomain.com;

    location / {
        proxy_pass http://ecommerce_backend;
    }
}
```

---

## Monitoring & Logging

### Setup Logging
```go
import "log"

log.SetFlags(log.LstdFlags | log.Lshortfile)
log.Println("Server started")
```

### Setup Monitoring
- Add Prometheus for metrics
- Add Grafana for dashboards
- Monitor response times
- Track error rates
- Monitor database performance

### Health Check Endpoint
Add to your routes:
```go
r.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
})
```

---

## Security in Production

### HTTPS
```bash
# Generate SSL certificate
# Using Let's Encrypt
certbot certonly -d yourdomain.com

# Update Gin
r.RunTLS(":443", "/path/to/cert.pem", "/path/to/key.pem")
```

### Rate Limiting
```bash
go get github.com/gin-contrib/ratelimit
```

### CORS Configuration
```go
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://yourdomain.com"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
}))
```

---

## Backup Strategy

### Database Backups
```bash
# SQLite
cp ecommerce.db backup-$(date +%Y%m%d).db

# PostgreSQL
pg_dump ecommerce > backup-$(date +%Y%m%d).sql

# MySQL
mysqldump ecommerce > backup-$(date +%Y%m%d).sql
```

### Automated Backups
Use cron jobs:
```bash
0 2 * * * /path/to/backup.sh
```

---

## Monitoring Checklists

### Daily
- [ ] Check error logs
- [ ] Verify database size
- [ ] Monitor server disk space

### Weekly
- [ ] Review API metrics
- [ ] Check backup integrity
- [ ] Analyze user activity

### Monthly
- [ ] Performance audit
- [ ] Security audit
- [ ] Data integrity check

---

## Troubleshooting

### Port Already in Use
```bash
# Find process using port
lsof -i :8080

# Kill process
kill -9 PID
```

### Database Connection Issues
```go
// Add connection string validation
// Test connection before running server
if err := database.DB.Exec("SELECT 1").Error; err != nil {
    log.Fatal("Database connection failed:", err)
}
```

### Memory Leaks
Monitor with pprof:
```go
import _ "net/http/pprof"
go http.ListenAndServe("localhost:6060", nil)
```

Access at `http://localhost:6060/debug/pprof/`

---

## Scaling

### Horizontal Scaling
- Use load balancer
- Multiple API instances
- Shared database

### Vertical Scaling
- Increase RAM
- Upgrade CPU
- Better storage

### Database Scaling
- Read replicas
- Connection pooling
- Query optimization

---

## Support & Documentation

- API Docs: [COMPLETE_API_DOCS.md](COMPLETE_API_DOCS.md)
- Features: [FEATURES.md](FEATURES.md)
- Setup: [SETUP.md](SETUP.md)

---

**Your ecommerce API is ready to go live!** 🚀
