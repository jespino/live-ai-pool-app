# Gitpod Configuration for Pool App

This directory contains configuration files for developing Pool App in [Gitpod](https://www.gitpod.io/), a cloud-based development environment.

## Features

- Preconfigured development environment
- Automated dependencies installation
- One-click commands through automations
- Hot-reloading for Go and React
- VS Code extensions pre-installed
- Organized automation groups by purpose

## Using Gitpod

1. Click the Gitpod button in your repository or visit gitpod.io/#<your-repository-url>
2. Wait for the workspace to initialize
3. Access the Automations panel by clicking the icon in the left sidebar or pressing `⌘+Shift+A` (Mac) or `Ctrl+Shift+A` (Windows/Linux)

## Available Automations

The automations are grouped by purpose for easy access:

### Development

| Command | Description |
|---------|-------------|
| Start Development Mode | Run both backend and frontend in development mode |
| Hot Reload Backend | Run the backend with air for hot reloading |
| Start Frontend Only | Run just the frontend development server |
| Run Backend Server | Run only the backend server |

### Build

| Command | Description |
|---------|-------------|
| Build App | Build both backend and frontend for production |
| Clean & Build | Clean artifacts and rebuild everything |

### Testing

| Command | Description |
|---------|-------------|
| Run Tests | Run the test suite |

### Dependencies

| Command | Description |
|---------|-------------|
| Update Go Dependencies | Update and tidy Go modules |
| Update Frontend Dependencies | Install frontend npm packages |

### Utilities

| Command | Description |
|---------|-------------|
| Show Help | Display all available make targets |

## Port Forwarding

| Port | Description |
|------|-------------|
| 8081 | Backend API server |
| 5173 | Frontend development server |

## Configuration Files

- `.gitpod.yml` - Main Gitpod configuration file
- `.gitpod/automations.yaml` - Automation commands definition
- `.gitpod/tasks.sh` - Setup script for tools installation

## Customizing

You can modify the `.gitpod.yml` and `.gitpod/automations.yaml` files to add or change automations based on your workflow needs.

For more information about the automations format, see the [Gitpod Automations Documentation](https://www.gitpod.io/docs/flex/configuration/automations/overview).