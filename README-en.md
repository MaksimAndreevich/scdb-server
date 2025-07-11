# What is this project?

Read about us [here](https://scdb-landing-001e.twc1.net)😉

[Документация на Русском языке](./README.md)

## SCDB Server

The server-side of the School Control Database (SCDB) educational organization management system.

## Description

SCDB Server is a REST API server providing access to an educational organization database. The server is implemented in Go using the Gin framework and PostgreSQL as the database.

## Key Features

- Retrieval of educational organization information
- Organization search with filtering and pagination
- Database statistics retrieval
- Support for various types of educational institutions

## Tech Stack

- Go 1.24.3
- Gin Web Framework
- PostgreSQL 17.4
- Docker & Docker Compose

## Requirements

- Go 1.24.3 or higher
- Docker and Docker Compose
- PostgreSQL 17.4 (if running without Docker)

## Installation and Running

### Using Docker

1. Clone the repository:

```bash
git clone https://github.com/MaksimAndreevich/scdb-server
cd server
```

2. Launch the application using Docker Compose:

```bash
docker-compose up -d
```

The server will be available at: [http://localhost:8080](http://localhost:8080)

### Local Installation

1. Install dependencies:

```bash
go mod download
```

2. Create a .env file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=db
```

3. Run the server:

```bash
go run main.go
```

### Database Population

To import data, use the [UPDATER service](https://github.com/MaksimAndreevich/scdb-updater). Data can be downloaded from the [registry of organizations](https://obrnadzor.gov.ru/otkrytoe-pravitelstvo/opendata/7701537808-raoo/).

### Organizations

- `GET /api/organizations/{id}` - Retrieve an organization by ID
- `GET /api/organizations` - Retrieve a list of organizations with pagination and filtering

  - Parameters:

    - search - Search query
    - federal_district_id - Federal district ID
    - region_id - Region ID
    - city_id - City ID
    - education_type_key - Education type
    - page - Page number
    - per_page - Number of records per page

### Statistics

- `GET /api/stats` - Retrieve overall database statistics

## Development

### Project Structure

```
.
├── docs/                    # Documentation
├── internal/               # Internal application code
│   ├── config/            # Configuration
│   ├── database/          # Database operations
│   └── routers/           # Routing
├── main.go                # Entry point
├── Dockerfile             # Docker configuration
├── docker-compose.yml     # Docker Compose configuration
└── go.mod                 # Go dependencies
```

### Hot Reload

For development, Air is used for automatic reloading on code changes:

```bash
air
```

### Creating a DEMO Image

1. Create a demo image

```bash
docker-compose -f docker-compose.demo.yml up -d
```

2. Populate the DB with data using the [UPDATER service](https://github.com/MaksimAndreevich/scdb-updater)

3. Connect to the DB and manually perform data obfuscation using the SQL script database/obfuscate_database.sql

4. Create a server image with data

```bash
docker commit scdb-server-demo scdb-server-demo:v1.0
```

5. Save the server image to a file

```bash
docker save scdb-server-demo:v1.0 | gzip > ../demo/scdb-server-demo.tar.gz
```

6. Create a database dump and save it to ../demo/ (container must be running)

```bash
docker exec scdb-postgres-demo pg_dump -U postgres -d db > ../demo/demo-database-dump.sql
```

## Contributing

We welcome suggestions, improvements, and bug reports!

- If you find a bug or want to suggest an improvement, create an issue.
- Want to propose a feature from the roadmap or a new idea? Feel free to open an issue with the enhancement label.
- Pull requests are welcome! Please follow basic code formatting rules and describe your changes.

📌 We are especially interested in contributions in the following areas:

- Authorization and security (API keys, restrictions)
- Caching and performance
- Demo mode support
- Data import/export

## Roadmap

### 1. Security and Authorization (High Priority)

- [ ] Basic database protection (write restrictions except from allowed IPs)
- [ ] API key system

  - Unique key generation
  - Database storage
  - IP restrictions
  - Request limits

- [x] Demo mode

  - Limited data set
  - Restricted functionality
  - Temporary access

- [ ] Middleware for API key validation
- [ ] Access logging system
- [ ] Lockout on suspicious activity

### 2. Performance Optimization

- [ ] Caching frequently requested data
- [ ] Optimization of SQL queries
- [ ] Performance monitoring

### 3. API Functionality Expansion

- [ ] Data export (CSV, Excel, PDF)
- [ ] Data import (CLI for easy import or admin interface on frontend — whichever is more complex)
- [ ] Statistics and analytics (requests processed, etc.)

### 4. Scaling

- [ ] Containerization (optimize)
- [ ] CI/CD pipeline

### 5. Testing

- [ ] Unit tests
- [ ] Integration tests
- [ ] Load testing
- [ ] Security testing
