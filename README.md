```
go-social/
│── cmd/               
│   ├── main.go
│── config/               # Config (Viper)
│   ├── config.yaml
│── internal/    
│   ├── auth/             # Authentication logic
│   ├── post/             # Post logic
│   ├── user/             # User management
│   ├── notification/     # Real-time notifications
│   ├── friendship/       # Friend system
│── pkg/                  # Reusable packages
│── infrastructure/       # External systems (DB, Redis, RabbitMQ)
│── docs/                 # API documentation (Swagger)
│── scripts/              # Deployment & CI/CD scripts
│── .env                  # Environment variables
│── go.mod
│── go.sum
│── README.md
```
# Register a new user
```
curl -X POST -H "Content-Type: application/json" -d '{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}' http://localhost:8888/api/auth/register
```
# Login
```
curl -X POST -H "Content-Type: application/json" -d '{
  "email": "test@example.com",
  "password": "password123"
}' http://localhost:8888/api/auth/login
```
# Get current user (use the token from login)
```
curl -X GET -H "Authorization: Bearer <YOUR_TOKEN>" http://localhost:8000/api/me
```

# API Endpoints Design
## Authentication

    POST /api/auth/register - User registration

    POST /api/auth/login - User login (JWT generation)

    POST /api/auth/refresh - Refresh JWT token

    GET /api/auth/me - Get current user info

## User Management

    GET /api/users - List users (with pagination)

    GET /api/users/:id - Get user details

    PUT /api/users/:id - Update user profile

## Post Management

    POST /api/posts - Create new post

    GET /api/posts/:id - Get post details

    GET /api/posts/user/:userId - Get user's posts

    PUT /api/posts/:id - Update post

    DELETE /api/posts/:id - Delete post

## Friendship System

    POST /api/friends/request - Send friend request

    GET /api/friends/requests - Get pending requests

    PUT /api/friends/accept/:id - Accept friend request

    DELETE /api/friends/reject/:id - Reject/unfriend

## Interactions

    POST /api/posts/:id/comments - Add comment to post

    POST /api/posts/:id/like - Like a post

    DELETE /api/posts/:id/like - Remove like

## Notifications

    GET /api/notifications - Get user notifications

    PUT /api/notifications/:id/read - Mark as read

## News Feed

    GET /api/feed - Get personalized news feed