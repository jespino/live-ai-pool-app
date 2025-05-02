# Pool App

A full-stack application for creating and managing polls/surveys that users can vote on.

## Components

- **Backend**: Go API server with in-memory or PostgreSQL storage
- **Frontend**: React TypeScript application with QR code voting capabilities

## Features

- Create and manage polls with multiple choice options
- Vote on polls via QR code on mobile devices
- View poll results with real-time updates
- User authentication and vote tracking
- Responsive design for all device types

## Architecture

The application consists of two main components:

1. **Backend API (Go)**: Handles data storage, poll management, and vote processing
2. **Frontend (React)**: Provides the user interface for both poll administrators and voters

## Getting Started

### Development Container (Recommended)

The project includes a devcontainer configuration for Visual Studio Code that sets up a complete development environment:

1. Install [VS Code](https://code.visualstudio.com/), [Docker](https://www.docker.com/products/docker-desktop), and the [Remote Development Extension Pack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)
2. Open VS Code, press F1, and select "Dev Containers: Open Folder in Container..."
3. Select the pool-app directory
4. Wait for the container to build and start (this may take a few minutes the first time)

See the [devcontainer README](./.devcontainer/README.md) for more details.

### Prerequisites (without devcontainer)

- Go 1.16+ (for backend)
- Node.js 14+ (for frontend)
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL (optional, for production use)

### Running the Full Stack with Docker

The easiest way to run the complete application is using Docker Compose:

```bash
# With in-memory database (default)
docker-compose up --build

# With PostgreSQL
DB_TYPE=postgres docker-compose up --build
```

This will:
- Start the Go backend API on port 8081
- Start the React frontend on port 5173
- Start PostgreSQL on port 5432 (if using postgres mode)

Access the application at http://localhost:5173

### Running Components Individually

#### Backend

See [backend instructions](#backend) below.

#### Frontend

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm run dev
```

The frontend will be available at http://localhost:5173

## Backend

The Go backend provides the API for managing polls and votes.

### API Endpoints

#### Pools

- `GET /api/pools` - List all pools
- `POST /api/pools` - Create a new pool
- `GET /api/pools/{id}` - Get a specific pool and its results
- `POST /api/pools/{id}/vote` - Vote on a specific pool

#### Users

- `POST /api/users` - Create a new user

### Storage Options

The backend supports two database options:

1. **In-memory database**: Default option, data is lost when the server restarts
2. **PostgreSQL**: Persistent storage, recommended for production

### Environment Variables

You can configure the backend using environment variables or a `.env` file:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_TYPE` | Database type (`memory` or `postgres`) | `memory` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | PostgreSQL username | `postgres` |
| `DB_PASSWORD` | PostgreSQL password | `postgres` |
| `DB_NAME` | PostgreSQL database name | `poolapp` |
| `DB_SSLMODE` | PostgreSQL SSL mode | `disable` |
| `PORT` | Server port | `8081` |

### Running the Backend

1. Navigate to the project root
2. Configure environment variables (optional)
3. Run the server:

```bash
make run
```

Or manually:

```bash
go run cmd/server/main.go
```

## Frontend

The React TypeScript frontend provides the user interface for both poll administrators and voters.

### Configuration

You can customize the frontend by editing the `frontend/src/config.ts` file:

- Configure polls to display
- Set the API URL
- Set refresh intervals
- Configure default display options

### Running the Frontend

See the [frontend README](./frontend/README.md) for detailed instructions.

## Deployment

For production deployment, use Docker Compose with PostgreSQL:

```bash
DB_TYPE=postgres docker-compose up -d
```

## Example API Usage

### Creating a Pool

```bash
curl -X POST http://localhost:8081/api/pools \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Favorite Programming Language",
    "description": "What is your favorite programming language?",
    "options": ["Go", "JavaScript", "Python", "Java", "C#"],
    "expires_in": 24
  }'
```

### Voting on a Pool

```bash
curl -X POST http://localhost:8081/api/pools/{pool_id}/vote \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "{user_id}",
    "option_id": "{option_id}"
  }'
```