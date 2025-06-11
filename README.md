# Go Web API with Redis Cache and PostgreSQL Persistence

## Architecture Overview

- **Go Web Server** – Handles routing, API logic, and landing page rendering.
- **Redis** – Acts as a fast-access cache layer.
- **PostgreSQL** – Stores persistent data when a cache miss occurs.
- **Docker** – Containers for easy deployment and consistency.
- **Rancher** – Optional container orchestration for scalable deployments.

## Features
- **RESTful API** for storing and retrieving key-value pairs.
- **Redis** for fast in-memory caching.
- **PostgreSQL** for persistent storage.
- Automatic fallback: If a key is not found in Redis, it is fetched from PostgreSQL and cached in Redis.
- Graceful error handling and resource management.


---
## Project Description

This is a simple web application written in **Go (Golang)** that exposes two endpoints:

- `/` – **Landing page**
- `/?` – **API endpoint** that accepts parameters via URL, checks if the result exists in a **Redis cache**, and:
  - If **found**: returns the cached result.
  - If **not found**: saves the data to **PostgreSQL**, then returns the result.

The application is fully containerized using **Docker** and can be run locally, with Docker Compose, or deployed via **Rancher** for orchestration.

---

## Requirements

To run the project you need to install:
- Go 1.18+
- PostgreSQL
- Redis
- Docker Compose
---


## Running with Docker

You can run the entire application stack (Go app, PostgreSQL, and Redis) using Docker Compose.

### 1. Build and Start All Services

From the project root, run:

```sh
docker-compose up --build
```

This will:
- Build the Go application image.
- Start PostgreSQL and Redis containers.
- Start the Go app, connecting to both services using environment variables.

### 2. Access the Application

Once all containers are running, open [http://localhost:8080](http://localhost:8080) in your browser to access the app.

### 3. Stopping the Stack

To stop all containers, press `Ctrl+C` in the terminal where Docker Compose is running, then run:

```sh
docker-compose down
```

### 4. Notes

- The database and Redis data are persisted in Docker volumes.
- The Go app uses environment variables (see `docker-compose.yml`) to connect to PostgreSQL and Redis.
- On first run, the database schema is initialized automatically using the SQL in the `db` directory.

---

## Code Overview

- **`db/dbService.go`**: Implements `Set`, `Get`, and `GetAll` for PostgreSQL.
- **`dataGetter.go`**: Handles Redis caching and orchestrates cache/database logic.
- **`main.go`**: Initializes services and starts the HTTP server.

## Troubleshooting

- Ensure PostgreSQL and Redis are running and accessible.
- Check environment variables for correct connection strings.
- Review console output for error logs.

## To Do

- [ ] Add user authentication and authorization
- [ ] Implement more robust error handling and logging
- [ ] Add unit and integration tests
- [ ] Add metrics and health check endpoints
- [ ] Improve landing page UI/UX
- [ ] Add CI/CD pipeline for automated testing and deployment

- [ ] Add database connection pooling and tuning
- [ ] Add database backup and restore scripts
- [ ] Support for multiple database environments (development, staging, production)
- [ ] Add data validation and constraints at the database level
- [ ] Add indexes for frequently queried columns

## License

MIT License

