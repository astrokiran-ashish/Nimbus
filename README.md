# Nimbus Service

Nimbus is the monolith service for Astrokiran, designed to handle various functionalities within a single codebase. This document provides an overview of the folder structure and instructions on how to run the application.

## Folder Structure

- **cmd/api**: Contains the main application code for the API server.
  - `main.go`: Entry point for the application.
  - `health.go`: Handles health check endpoints.
  - `errors.go`: Contains error handling logic.
  - `middleware.go`: Defines middleware functions for the application.
  - `server.go`: Manages the HTTP server setup and lifecycle.

## Modules Overview
- **Consultation**: Handles consultation-related functionalities.
- **Consultant**: Manages consultant profiles and expertise.
- **User**: Manages user profiles and authentication.
- **Communication**: Handles communication channels.
- **Workflows**: Manages workflows for consultations and other processes.
- **Notification**: Manages notifications and alerts.

- **internal/common**: Contains shared utilities and configurations.
  - **configs**: Configuration management utilities.
  - **database**: Database connection and management utilities.

- **vendor**: Contains third-party dependencies and their respective documentation and licenses.

## How to Run the Application

1. **Prerequisites**: Ensure you have Go 1.20 or later installed on your machine.

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/astrokiran/nimbus.git
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

## Models and Data Structures
### SchemaMigrations
- **Fields**: Version, Dirty

### Role
- **Fields**: RoleID, RoleName, RoleDescription, Version, CreatedAt, UpdatedAt

### Permission
- **Fields**: PermissionID, PermissionName, PermissionDescription, RoleID, Version, CreatedAt, UpdatedAt

### User
- **Fields**: UserID, Name, Email, Dob, PhoneNumber, Version, CreatedAt, UpdatedAt

### Consultant
- **Fields**: ConsultantID, UserID, Expertise, State, Version, ChatChannel, CallChannel, LiveChannel, VideoCallChannel, CreatedAt, UpdatedAt

### Consultation
- **Fields**: ConsultationID, UserID, ConsultantID, SessionID, ConsultationTimeSecs, ConsultationType, ConsultationState, UserWaitTimeSecs, AgoraChannel, CreatedAt, UpdatedAt

### UserAuth
- **Fields**: ID, UserID, SessionID, JwtTokenHash, RefreshTokenHash, DeviceDetails, OTP, OTP_created_at, OTP_validity_secs, OTP_attempts, PhoneNumber, CreatedAt, ExpiresAt, UpdatedAt

### UserRoles
- **Fields**: UserID, RoleID, CreatedAt

## Database Schema
### Role Table
- Columns: role_id, role_name, role_description, version, created_at, updated_at

### Permission Table
- Columns: permission_id, permission_name, permission_description, role_id, version, created_at, updated_at

### User Table
- Columns: user_id, name, email, dob, phone_number, version, created_at, updated_at

### Consultant Table
- Columns: consultant_id, user_id, expertise, state, version, chat_channel, call_channel, live_channel, video_call_channel, created_at, updated_at

### Consultation Table
- Columns: Consultation_ID, User_ID, Consultant_ID, Session_ID, Consultation_Time_Secs, Consultation_Type, Consultation_State, User_Wait_Time_Secs, Agora_Channel, Created_At, Updated_At

### User Auth Table
- Columns: id, user_id, session_id, jwt_token_hash, refresh_token_hash, device_details, OTP, OTP_created_at, OTP_validity_secs, OTP_attempts, phone_number, created_at, expires_at, updated_at

