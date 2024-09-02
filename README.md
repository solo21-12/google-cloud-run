---

# Go REST API for User, Group, and Role Management

![Go](https://img.shields.io/badge/Go-1.22-blue.svg) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg) ![Docker](https://img.shields.io/badge/Docker-20.10-blue.svg) ![Makefile](https://img.shields.io/badge/Makefile-Enabled-brightgreen.svg)

## Overview

This project is a Golang-based REST API designed to manage users, groups, and roles within a PostgreSQL database. It provides complete CRUD operations for each entity, along with specific endpoints to handle the relationships between users, groups, and roles. The API is optimized to run on Google Cloud Run.

## Features

- **Users**: Create, update, delete, and manage user details. Query users by UID or search by name.
- **Groups**: Manage groups, including adding/removing users from groups.
- **Roles**: Assign roles to users, manage role details including rights stored as JSON.
- **JWT-Based Authentication**: Secure endpoints using JWT tokens.
- **Dynamic Database Selection**: Middleware dynamically selects the database based on the authorization header.

## Endpoints

### Users
- `GET /users`: Retrieve all users.
- `GET /users?search=searchterm&limit=max_results&orderby=column`: Search users by name.
- `GET /users/{uid}`: Retrieve user details by UID, including assigned role.
- `GET /users/{uid}/groups`: Retrieve all groups associated with the user.
- `POST /users`: Create a new user.
- `PUT /users/{uid}`: Update user details.
- `DELETE /users/{uid}`: Delete a user.
- `Patch  /users/{uid}/groups`: Add user to groups

### Groups
- `GET /groups`: Retrieve all groups.
- `GET /groups/{uid}`: Retrieve group details by UID.
- `GET /groups/{uid}/users`: Retrieve all users within a specific group.
- `POST /groups`: Create a new group.
- `PUT /groups/{uid}`: Update group details.
- `DELETE /groups/{uid}`: Delete a group.

### Roles
- `GET /roles`: Retrieve all roles.
- `GET /roles/{uid}`: Retrieve role details by UID.
- `GET /roles/{uid}/users`: Retrieve all users assigned to a specific role.
- `POST /roles`: Create a new role.
- `PUT /roles/{uid}`: Update role details.
- `DELETE /roles/{uid}`: Delete a role.

## Project Structure

```plaintext
├── cmd
│   └── main.go
├── config
│   ├── env.go
│   └── postgres.go
├── Delivery
│   ├── Controllers
│   ├── Middlewares
│   └── Routers
├── Dockerfile
├── Domain
│   ├── Dtos
│   ├── Interfaces
│   └── Models
├── Infrastructure
├── Repository
├── Tests
├── tmp
└── Usecases
```

## Local Setup

To run this project locally, you need to set up the following environment variables in the docker-compose file:

```plaintext
JWT_SECRET="your-jwt-secret"
DB_USER="your-db-user"
DB_PASS="your-db-password"
DB_HOST="localhost"
DB_PORT=5432
```

### Running the Application

- **Using Docker**:  
  Build and run the Docker container:
  ```bash
  make run
  ```

**Note**: Make sure to update the `Docker-compose` with your environment variable information.

### Running Tests

To execute the tests, run:
```bash
make test
```

### To stop the server
```bash
make down
```

### To see logs
```bash
make logs
```

## API Documentation

For detailed API documentation, please refer to the [Postman Collection](https://documenter.getpostman.com/view/22911710/2sAXjM2qbR).

## Deployment on Google Cloud Run

This project is designed to be deployed on Google Cloud Run. Ensure all necessary configurations are set before deploying.

---
