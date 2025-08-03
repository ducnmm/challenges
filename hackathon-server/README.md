# Hackathon Server

Go web server với JWT authentication và file upload.

## Features

- **JWT Authentication**: Register, login với HS256 tokens
- **File Upload**: Image upload với validation (max 8MB)
- **SQLite Database**: User data và file metadata
- **Clean Architecture**: Modular code structure

## Quick Start

```bash
# Install dependencies
go mod tidy

# Run server
go run main.go

# Open browser
open http://localhost:8080
```

## API Endpoints

### Authentication
```bash
# Register
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"user","password":"pass"}'

# Login (returns JWT token)
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"user","password":"pass"}'
```

### File Upload
```bash
# Upload image (requires JWT token)
curl -X POST http://localhost:8080/upload \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "data=@image.jpg"
```

## Environment Variables

```bash
JWT_SECRET=your-secret-key    # JWT signing key
PORT=8080                     # Server port  
DB_PATH=./hackathon.db        # SQLite file path
```

## Testing

```bash
# Run unit tests
go test ./... -v

# Run with coverage
go test -cover ./...
```

## Project Structure

```
├── main.go                 # Entry point
├── models/                 # Data structures
├── handlers/               # HTTP handlers  
├── middleware/             # JWT auth
├── database/               # SQLite setup
└── utils/                  # Helper functions
```

## Security

- bcrypt password hashing (cost 14)
- JWT tokens (24h expiry)
- Image-only file uploads
- SQL injection protection
- File size limits (8MB)