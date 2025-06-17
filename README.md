# Ticketing System - Proactive Maintenance & Asset Management Platform

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue)](https://golang.org)
[![Node Version](https://img.shields.io/badge/Node-16%2B-green)](https://nodejs.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13%2B-blue)](https://www.postgresql.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

![Ticketing System Banner](https://placehold.co/1200x300/6c5ce7/ffffff?text=Ticketing+System&font=inter)

## 📋 Overview

**Ticketing System** is a full-stack application designed to revolutionize maintenance operations for a large chain of cafes. It shifts the paradigm from a reactive, costly repair model to a proactive, data-driven maintenance strategy. This platform provides a centralized system for tracking maintenance tickets, managing all physical assets via QR codes, scheduling preventive care, and generating insightful reports for management.

### 🎯 Key Benefits

- **Reduced Downtime**: Proactive maintenance scheduling prevents equipment failures
- **Cost Optimization**: Better inventory management and preventive maintenance reduce repair costs
- **Improved Efficiency**: Streamlined workflows and automated task routing
- **Data-Driven Decisions**: Comprehensive reporting and analytics
- **Enhanced Asset Life**: Regular maintenance extends equipment lifespan

---

## ✨ Core Features

* **Centralized Dashboard**: A real-time command center showing KPIs like open tickets, resolution times, and critical alerts.
* **QR-Powered Asset Management**: Every piece of equipment gets a unique QR code for instant identification and access to its full maintenance history.
* **Streamlined Ticketing System**: Branch staff can create detailed tickets in seconds, which are automatically routed to the appropriate technician.
* **Automated Preventive Maintenance**: The system automatically generates tickets for routine servicing based on predefined schedules, preventing failures before they happen.
* **Spare Parts Inventory**: Tracks parts usage and inventory levels, flagging items for reordering to prevent delays.
* **Robust Reporting**: Generates daily summary reports and allows for trend analysis over time to inform business decisions.
* **Role-Based Access Control**: Different views and permissions for Branch Staff, Technicians, and Managers.

---

## 🚀 Technology Stack

This project is built with a modern, scalable, and type-safe technology stack.

### Backend
* **Go (Golang)** v1.18+
  * `gqlgen` for GraphQL API generation
  * `gin` for HTTP routing
  * `gorm` for database ORM
  * `jwt-go` for authentication

### Frontend
* **SvelteKit** v1.0+
  * **Tailwind CSS** v3.0+ for styling
  * **Apollo Client** for GraphQL integration
  * **Chart.js** for data visualization
  * **QR Code Generator** for asset tagging

### Database
* **PostgreSQL** v13+
  * JSONB support for flexible data storage
  * Custom ENUM types for strict data validation
  * Full-text search capabilities
  * Robust indexing for performance

### Development Tools
* **Docker** for containerization
* **Air** for hot reloading
* **ESLint** & **Prettier** for code formatting
* **Jest** for testing

---

## 📂 Project Structure

```
ticket-system/
├── /backend-go/
│   ├── /graph/          # gqlgen generated files, schema, and resolvers
│   │   ├── /generated/  # Auto-generated GraphQL code
│   │   ├── /model/      # GraphQL models
│   │   └── /resolvers/  # GraphQL resolvers
│   ├── /db/             # Database connection and queries
│   │   ├── /migrations/ # Database migrations
│   │   └── /models/     # Database models
│   ├── /internal/       # Core business logic and services
│   │   ├── /auth/       # Authentication logic
│   │   ├── /middleware/ # HTTP middleware
│   │   └── /services/   # Business services
│   ├── /config/         # Configuration files
│   └── go.mod
├── /frontend-svelte/
│   ├── /src/
│   │   ├── /routes/     # SvelteKit file-based routing
│   │   │   ├── /tickets/ # The main tickets dashboard
│   │   │   └── /reports/ # The reporting page
│   │   ├── /lib/        # Reusable components and utilities
│   │   │   ├── /components/ # UI components
│   │   │   ├── /stores/    # Svelte stores
│   │   │   └── /utils/     # Helper functions
│   │   └── /styles/     # Global styles
│   ├── /static/         # Static assets
│   └── svelte.config.js
├── /migrations/
│   ├── /sql/            # Migration files (.up.sql and .down.sql)
│   ├── Makefile         # Migration commands and utilities
│   └── README.md        # Detailed migration documentation
├── /docs/               # Project documentation
├── /scripts/            # Utility scripts
├── .env.example         # Example environment variables
├── docker-compose.yml   # Docker configuration
└── README.md
```

---

## 🏁 Getting Started

To get the full Ticketing System application running locally, follow these steps.

### Prerequisites

* Go (version 1.18+)
* Node.js (version 16+)
* PostgreSQL (version 13+)
* Docker (optional, for containerized setup)

### 1. Database Setup

#### Prerequisites

First, install the `golang-migrate` CLI tool:

**macOS (Homebrew):**
```bash
brew install golang-migrate
```

**Linux/macOS (Binary):**
```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

#### Migration Setup

1. **Configure environment variables** by creating a `.env` file in the project root:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=ticketing_system
   DB_SSL_MODE=disable
   ```

2. **Navigate to the migrations directory**:
   ```bash
   cd migrations/
   ```

3. **Create the database** (if it doesn't exist):
   ```bash
   make create-db
   ```

4. **Run all migrations**:
   ```bash
   make up
   ```

5. **Verify the setup**:
   ```bash
   make status
   ```

For detailed migration management, see the [migrations README](./migrations/README.md).

### 2. Backend (Go & GraphQL)

1. Navigate to the `/backend-go` directory.
2. Create a `.env` file and configure your database connection string:
   ```
   DATABASE_URL="postgres://your_username:your_password@localhost:5432/ticketing_system?sslmode=disable"
   JWT_SECRET="your-secret-key"
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Generate GraphQL models and resolvers:
   ```bash
   go run github.com/99designs/gqlgen generate
   ```
5. Start the backend server:
   ```bash
   go run server.go
   ```
   The GraphQL API will be running at `http://localhost:8080`.

### 3. Frontend (SvelteKit)

1. Navigate to the `/frontend-svelte` directory.
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the SvelteKit development server:
   ```bash
   npm run dev
   ```
   The frontend application will be available at `http://localhost:5173`.

### 4. Docker Setup (Optional)

For a containerized setup:

```bash
# Build and start all services
docker-compose up --build

# Run in detached mode
docker-compose up -d

# Stop all services
docker-compose down
```

---

## 🔄 Development Workflow

1. **Branch Management**
   - `main` - Production-ready code
   - `develop` - Integration branch
   - `feature/*` - New features
   - `bugfix/*` - Bug fixes
   - `hotfix/*` - Urgent production fixes

2. **Code Style**
   - Follow Go best practices and style guide
   - Use Prettier for frontend code formatting
   - Write meaningful commit messages

3. **Testing**
   - Write unit tests for critical functionality
   - Run tests before committing
   - Maintain good test coverage

4. **Documentation**
   - Update documentation for new features
   - Keep API documentation current
   - Document breaking changes

---

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Issues**
   - Verify PostgreSQL is running
   - Check connection string in `.env`
   - Ensure database exists

2. **GraphQL Generation Errors**
   - Clear generated files
   - Update gqlgen
   - Check schema syntax

3. **Frontend Build Issues**
   - Clear node_modules
   - Update dependencies
   - Check for version conflicts

### Getting Help

- Check the [documentation](./docs)
- Search existing issues
- Create a new issue with detailed information

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests
5. Submit a pull request

See [CONTRIBUTING.md](./CONTRIBUTING.md) for detailed guidelines.

---

## 📖 API Documentation

The API is documented via the GraphQL schema itself. Once the backend server is running, you can explore the entire API interactively using the **GraphQL Playground** available at `http://localhost:8080/`. You can write queries, view the schema, and see real-time results.

For detailed API documentation, see [API.md](./docs/API.md).

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- Thanks to all contributors
- Inspired by modern maintenance management systems
- Built with amazing open-source tools
