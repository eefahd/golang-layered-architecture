# Go Layered Architecture Example

A **minimal example** demonstrating the **Layered Architecture pattern** in Go.
It shows how to structure an application with a clear **separation of concerns** between presentation, business logic, data access, and domain layers, improving maintainability, scalability, and testability.

> ⚠️ **Note:** This project is **for learning purposes only**.
> It intentionally omits many production aspects (comprehensive error handling, logging, validation, security, testing, etc.) to keep the focus on **core architectural principles**.

---

## Overview

This example implements a simple **Contact Manager** application that demonstrates:

* **Clean separation of concerns** across architectural layers
* **Swappable data stores** (SQLite, PostgreSQL, and File-based storage)
* **Multiple presentation interfaces** (REST API and CLI)
* **Factory pattern** for store and database creation
* **Dependency injection** for loose coupling

Each layer interacts only with the one directly below it, enforcing clean boundaries and easier testability.

---

## Layer Details

| Layer                          | Responsibility                    | Example                          |
| ------------------------------ | --------------------------------- | -------------------------------- |
| **Presentation**               | Entry points for user interaction | Chi-based REST API, CLI          |
| **Services (Business Logic)**  | Core logic orchestration          | `ContactService`                 |
| **Store (Data Access)**        | Persistence abstraction           | SQLite, PostgreSQL, File stores  |
| **Models (Domain Entities)**   | Data definition & validation      | Go structs                       |

---

## ⚙️ Design Choice: A Minimal Service Layer

The **service layer** here is intentionally left without abstractions to show that, depending on project needs, such layers may not require them.
Some engineers prefer to abstract every layer for consistency; others apply abstraction only when variation or testing needs justify it. This example simply illustrates a practical approach, not a statement on which is universally better.

---

## 📂 Project Structure

```bash
.
├── cmd/                        # Application entry points
│   ├── http/                   # HTTP API server
│   └── cli/                    # Command-line interface
├── internal/                   # Private application code
│   ├── models/                 # Domain models (structs)
│   ├── service/                # Business logic layer
│   ├── store/                  # Data access layer
│   │   ├── interfaces/         # Store interfaces
│   │   ├── sqlite/             # SQLite implementation
│   │   ├── postgres/           # PostgreSQL implementation
│   │   ├── filestore/          # File-based storage
│   │   └── factory.go          # Factory for store creation
│   ├── database/               # Database connection management
│   │   └── factory.go          # Factory for database creation
│   ├── server/                 # Presentation layer
│   │   ├── http/               # HTTP server (Chi router)
│   │   └── cli/                # CLI interface
│   ├── config/                 # Configuration management
│   └── utils/                  # Utilities (e.g., email messaging)
├── db/                         # Database files and migrations
│   └── migrations/             # SQL migration scripts
├── config.json                 # Default configuration
├── Dockerfile
├── docker-compose.yml
├── Makefile                    # Simplified development commands
├── go.mod
└── go.sum
```

---

## 🚀 Getting Started

### 🧱 Prerequisites

* **Docker & Docker Compose** (recommended)
* **Go 1.24+** for local development

---

### ⚡ Quick Start

The project includes a **Makefile** to simplify running and testing.
List available commands:

```bash
make help
```

---

### ▶️ Run the Application

**Start HTTP API with SQLite:**

```bash
make api-sqlite
# Access the API at http://localhost:8081
# Health check: http://localhost:8081/health
```

**Interactive CLI with SQLite:**

```bash
make cli-sqlite
```

**Start HTTP API with PostgreSQL:**

```bash
make api-postgres
# Access the API at http://localhost:8080
# Health check: http://localhost:8080/health
```

**Interactive CLI with PostgreSQL:**

```bash
make cli-postgres
```

---

## 🌐 API Endpoints

| Method   | Endpoint           | Description                        |
| -------- | ------------------ | ---------------------------------- |
| `GET`    | `/health`          | Health check endpoint              |
| `GET`    | `/contacts`        | List all contacts                  |
| `GET`    | `/contacts/{id}`   | Get a specific contact by ID       |
| `POST`   | `/contacts`        | Create a new contact               |
| `PUT`    | `/contacts/{id}`   | Update an existing contact         |
| `DELETE` | `/contacts/{id}`   | Delete a contact                   |

---

## ⚙️ Configuration

The application uses JSON configuration files to manage different database backends.

---

## 📘 Technical Terminology

Different sources may use varying terms for similar architectural concepts.
Here's how they map in this project:

| Term             | Also Known As                | Description                    |
| ---------------- | ---------------------------- | ------------------------------ |
| **Store**        | Repository, Persistence Layer, DAO | Handle database operations |
| **Services**     | Application Layer, Use Cases | Contain business logic         |
| **Presentation** | Controllers, Handlers, UI    | User-facing entry points       |
| **Models**       | Entities, Domain Models, DTO | Data structures and validation |

> There's no single "correct" terminology, adapt these concepts to your project's needs.

---

## 🛠️ Development

### Build Locally

```bash
# Build HTTP API
go build -o bin/api ./cmd/http

# Build CLI
go build -o bin/cli ./cmd/cli

# Run HTTP API
./bin/api

# Run CLI
./bin/cli
```

---

## 📚 Learn More

For a detailed explanation of the layered architecture pattern and the design decisions behind this example, see:

👉 **[Medium Article: How Layered Architecture Just Makes Sense. A Natural Way to Understand It](https://medium.com/@fahd.hus/how-layered-architecture-just-makes-sense-a-natural-way-to-understand-it-d85dce8ce914)**

---

## 🪪 License

This project is provided for **educational purposes** under the [MIT License](LICENSE).
Use it freely as a reference, and remember to add **comprehensive error handling, security, testing, and production best practices** before using it in real applications.
