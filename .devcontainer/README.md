# Pool App Development Container

This development container provides a complete development environment for the Pool App project including:

- Go 1.20 with standard tools and linting
- Node.js LTS with NPM
- Docker-in-Docker support
- VS Code extensions and settings

## Getting Started

### Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- [Remote Development Extension Pack](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)

### Opening the Dev Container

1. Open VS Code
2. Press F1 and select "Dev Containers: Open Folder in Container..."
3. Select the pool-app directory
4. Wait for the container to build and start

VS Code will automatically build the development container and connect to it. This process may take several minutes the first time.

## Components

The development container includes:

- Go development environment with Go 1.20
- Node.js for frontend development
- Hot-reloading for development
- Docker-in-Docker support for container testing
- Code formatting and linting tools

## Services and Ports

| Service        | Port | Description                     |
|----------------|------|---------------------------------|
| Go Backend     | 8081 | Pool API Server                 |
| React Frontend | 5173 | Pool UI (Vite dev server)       |

## Development Workflow

### Backend Development

The Go server supports hot reloading via Air:

```bash
cd /workspace
air
```

### Frontend Development

The React frontend uses Vite for development:

```bash
cd /workspace/frontend
npm run dev
```

## Notes

- Node modules are stored in a volume to improve performance
- Go modules are cached in a volume