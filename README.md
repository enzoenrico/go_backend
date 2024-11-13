# go_backend

A REST API backend developed for the backend course @pucpr using Go and Echo framework
Made with blood, sweat & tears by [Enzo Enrico](https://github.com/enzoenrico) & [Theo Ravgalia](https://github.com/TheoRavaglia)

[Youtube video about the project](https://youtu.be/xOGrUIhQ9iw)

## Project Structure

>server.go - Main entry point and server configuration
>
> app/
> ├── config.go - Application configuration
> ├── config.json       - Config file for JWT and roles
> ├── database/        - Database operations
> ├── handlers/        - HTTP route handlers
> ├── logger/          - Logging configuration
> ├── permissions/     - Role-based permissions
> ├── posts/           - Post model
> └── users/           - User model and authentication

## Stack Used

- Echo - High performance web framework
- Air - Hot reload for development
- Zap - Structured logging
- JWT - Authentication
- SQLite3 - Database (pending)

## Features

### Implemented ✅

- [x] REST API Routing
  - User routes (GET/POST)
  - Post routes (GET/POST)
- [x] JWT Authentication
- [x] Role-based Authorization
- [x] Structured Logging with Zap
- [x] Configuration Management
- [x] In-memory Data Storage

## API Endpoints

### Users

- `GET /users` - Get all users (requires JWT)
- `GET /users/:id` - Get user by ID (requires JWT)
- `POST /users` - Create new user (requires JWT)

### Posts

- `GET /posts` - Get all posts
- `GET /posts/:id` - Get post by ID
- `POST /posts` - Create new post

### Authentication

- `POST /login` - Get JWT token

## Development

Run the server with hot reload:

```sh
    air -c ./air.toml
```

The API will be available at `http://localhost:5000`
