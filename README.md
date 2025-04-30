# UMS - User Management System

A comprehensive user management system for an election platform, supporting different user roles, verification hierarchies, and blockchain integration.

## Overview

UMS (User Management System) is a Go-based service that handles user management for an election platform. The system supports various roles in the election process including voters, KPU Kota (City Election Committee), KPU Provinsi (Provincial Election Committee), and KPU Pusat (Central Election Committee).

The service implements a hierarchical verification process where:
- KPU Pusat can verify KPU Provinsi
- KPU Provinsi can verify KPU Kota
- KPU Kota can verify Voters

Key features include user registration, authentication, role-based access control, and blockchain integration for secure and immutable record-keeping.

## System Requirements

- Go 1.23.6 or higher
- PostgreSQL
- Ethereum client (Ganache for development)
- Kafka (for event messaging)
- Docker & Docker Compose (for containerization)

## Setup and Installation

### Using Docker

1. Clone the repository:
   ```bash
   git clone https://github.com/nocturna-ta/ums.git
   cd ums
   ```

2. Create a configuration file:
   ```bash
   cp config/files/config.yaml.example config/files/config.yaml
   ```

3. Update the configuration with your environment settings

4. Build and run with Docker:
   ```bash
   docker build -t ums:latest .
   docker run -p 8900:8900 -p 35000:35000 ums:latest
   ```

### Manual Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/nocturna-ta/ums.git
   cd ums
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create and configure the database:
   ```bash
   createdb ums
   migrate -path db/migrations -database "postgres://[username]:[password]@localhost:5432/ums?sslmode=disable" up
   ```

4. Configure your environment:
   ```bash
   cp config/files/config.yaml.example config/files/config.yaml
   # Edit config/files/config.yaml with your settings
   ```

5. Run the API server:
   ```bash
   make run-api
   ```

6. Alternatively, run the gRPC server:
   ```bash
   make run-grpc
   ```

## Project Structure

The project follows clean architecture principles with a clear separation of concerns:

```
├── cmd/                 # Command-line interfaces
├── config/              # Configuration
├── db/                  # Database migrations
├── ethreum/             # Ethereum blockchain integration
├── internal/            # Application core
│   ├── domain/          # Domain models and repositories
│   ├── handler/         # HTTP and gRPC handlers
│   ├── infrastructures/ # Infrastructure components
│   ├── interfaces/      # Interface implementations
│   └── usecases/        # Business logic
├── pkg/                 # Shared packages
│   ├── constants/       # Constants
│   ├── roles/           # Role definitions
│   └── utils/           # Utilities
```

## Features

- **User Management**: Registration, login, and profile management
- **Role-Based Access Control**: Different permissions for various roles
- **Verification Workflow**: Hierarchical verification process
- **Blockchain Integration**: Secure and immutable data storage
- **Event-Driven Architecture**: Kafka-based event messaging
- **API Access**: Both HTTP REST and gRPC endpoints

## API Endpoints

The service exposes both HTTP and gRPC endpoints:

### HTTP API (Port 8900)

- **Authentication**
    - POST `/v1/user/register` - Register a new user
    - POST `/v1/user/login` - User login

- **User Management**
    - GET `/v1/user/me` - Get current user
    - PUT `/v1/user/update` - Update user info

- **KPU Management**
    - GET `/v1/kpu-provinsi` - Get all provincial committees
    - GET `/v1/kpu-kota` - Get all city committees

- **Voter Management**
    - POST `/v1/voter/register` - Register a voter
    - GET `/v1/voter/nik/{nik}` - Get voter by NIK

- **Verification**
    - GET `/v1/verifications/pending` - Get pending verifications
    - POST `/v1/verifications/approve` - Approve a verification
    - POST `/v1/verifications/reject` - Reject a verification

### gRPC Server (Port 35000)

The service also provides gRPC endpoints for authentication validation.

## Configuration

Configuration is managed through YAML files in the `config/files` directory. Key configuration sections:

- **Server**: HTTP server settings
- **API**: API endpoints and timeouts
- **Database**: PostgreSQL connection details
- **Blockchain**: Ethereum client and contract addresses
- **JWT**: Authentication tokens
- **Kafka**: Message broker settings

## Development

### Makefile Commands

- `make dependency` - Download dependencies
- `make swag-init` - Initialize Swagger documentation
- `make run-api` - Run the HTTP API server
- `make run-grpc` - Run the gRPC server
- `make migrate-up` - Run database migrations
- `make migrate-down` - Rollback database migrations
- `make remock` - Regenerate mock objects for testing

## License

This project is licensed under the MIT License - see the LICENSE file for details.