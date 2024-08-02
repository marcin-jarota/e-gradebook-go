# School Management System

This project is a comprehensive School Management System implemented using Go for the backend and Vue.js for the frontend. It follows Domain-Driven Design (DDD) principles and Clean Architecture to create a scalable and maintainable application.

## Project Structure

```bash
├── cmd # Command-line interfaces
│ ├── api # Main API application
│ └── seed # Database seeding utility
├── db # Database-related files and migrations
├── internal # Internal packages
│ ├── adapters # Adapters for external services (DB, storage, transport)
│ ├── app # Application core (domain models and ports)
│ ├── auth # Authentication-related functionality
│ ├── class_group # Class group management
│ ├── lesson # Lesson management
│ ├── mark # Grading and marks
│ ├── middleware # HTTP middleware
│ ├── notification # Notification system
│ ├── school_year # School year management
│ ├── student # Student management
│ ├── subject # Subject management
│ ├── teacher # Teacher management
│ └── user # User management
├── uml # UML diagrams
└── web # Frontend Vue.js application
```

## Key Features

- User authentication and authorization
- Student, teacher, and class management
- Subject and lesson planning
- Grading system
- Notifications
- School year management

## Technology Stack

- Backend: Go
- Frontend: Vue.js
- Database: PostgreSQL (inferred from SQL files)
- Cache: Redis (inferred from `redis_storage.go`)
- API: RESTful (inferred from project structure)

## Getting Started

### Prerequisites

- Go 1.20
- Node.js and npm
- Docker and Docker Compose

## Architecture

This project follows Clean Architecture and Domain-Driven Design principles:

- The `internal/app/domain` directory contains the core domain models.
- The `internal/app/ports` directory defines the interfaces for the application.
- Each feature (e.g., `class_group`, `lesson`, `mark`) is organized into its own package with handlers, repositories, and services.

![architecture diagram](https://github.com/marcin-jarota/e-gradebook-go/blob/main/uml/architecture.png?raw=true)
