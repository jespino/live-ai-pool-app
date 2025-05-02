# Gitpod Configuration for Pool App

This directory contains configuration files for developing Pool App in [Gitpod](https://www.gitpod.io/), a cloud-based development environment.

## Features

- Preconfigured development environment
- Automated dependencies installation
- One-click commands through automations
- Hot-reloading for Go and React
- VS Code extensions pre-installed

## Using Gitpod

1. Click the Gitpod button in your repository or visit gitpod.io/#<your-repository-url>
2. Wait for the workspace to initialize
3. Access the Automations panel by clicking the icon in the left sidebar or pressing `⌘+Shift+A` (Mac) or `Ctrl+Shift+A` (Windows/Linux)

## Available Automations

The following commands are available through the Automations panel:

| Command | Description |
|---------|-------------|
| Start Development Mode | Run both backend and frontend in development mode |
| Hot Reload Backend | Run the backend with air for hot reloading |
| Start Frontend Only | Run just the frontend development server |
| Build App | Build both backend and frontend for production |
| Run Tests | Run the test suite |
| Clean & Build | Clean artifacts and rebuild everything |
| Run Backend Server | Run only the backend server |
| Update Go Dependencies | Update and tidy Go modules |
| Update Frontend Dependencies | Install frontend npm packages |
| Show Help | Display all available make targets |

## Port Forwarding

| Port | Description |
|------|-------------|
| 8081 | Backend API server |
| 5173 | Frontend development server |

## Customizing

You can modify the `.gitpod.yml` and `.gitpod/automations.yaml` files to add or change automations based on your workflow needs.