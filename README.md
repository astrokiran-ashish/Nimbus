# Nimbus Service

Nimbus is the monolith service for Astrokiran, designed to handle various functionalities within a single codebase. This document provides an overview of the folder structure and instructions on how to run the application.

## Folder Structure

- **cmd/api**: Contains the main application code for the API server.
  - `main.go`: Entry point for the application.
  - `health.go`: Handles health check endpoints.
  - `errors.go`: Contains error handling logic.
  - `middleware.go`: Defines middleware functions for the application.
  - `server.go`: Manages the HTTP server setup and lifecycle.

- **internal/common**: Contains shared utilities and configurations.
  - **configs**: Configuration management utilities.
  - **database**: Database connection and management utilities.

- **vendor**: Contains third-party dependencies and their respective documentation and licenses.

## How to Run the Application

1. **Prerequisites**: Ensure you have Go 1.20 or later installed on your machine.

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/astrokiranashish/nimbus.git
   cd nimbus
   ```

3. **Install Dependencies**:
   Run the following command to install the necessary Go modules:
   ```bash
   go mod tidy
   ```

4. **Build the Application**:
   Use the Makefile to build the application:
   ```bash
   make build
   ```

5. **Run the Application**:
   Start the application using the following command:
   ```bash
   make run
   ```

   Alternatively, you can run the application with live reloading on file changes:
   ```bash
   make run/live
   ```

6. **Access the Application**:
   The application will be accessible at `http://localhost:4444` by default. You can change the port by setting the `HTTP_PORT` environment variable.

## Configuration

Configuration values are managed through environment variables. The following are some of the key configurations:

- `BASE_URL`: The base URL for the application.
- `HTTP_PORT`: The port on which the application will run.
- `DB_DSN`: The Data Source Name for the database connection.
- `DB_AUTOMIGRATE`: Boolean to enable or disable automatic database migrations.
