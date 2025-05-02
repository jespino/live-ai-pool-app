# Pool App Frontend

This is the React TypeScript frontend for the Pool App, allowing users to create polls, vote, and view results.

## Features

- View active polls with a QR code for easy access
- Vote on polls from mobile devices
- View real-time poll results
- Authentication system to ensure one vote per user
- Responsive design for all device types

## Technology Stack

- React with TypeScript
- Vite for fast development and builds
- Chakra UI for responsive, accessible components
- React Router for navigation
- Axios for API communication
- QR Code generator for mobile access

## Getting Started

### Prerequisites

- Node.js 14+
- npm or yarn
- Backend API running (see main project README)

### Installation

1. Install dependencies:

```bash
npm install
```

2. Configure environment variables:

Create a `.env` file with the following content:

```
VITE_API_URL=http://localhost:8080/api
```

3. Start the development server:

```bash
npm run dev
```

The application will be available at http://localhost:5173.

## Configuration

You can customize the application by editing the `src/config.ts` file. This allows you to:

- Define polls to display
- Set the API URL
- Configure refresh intervals
- Set the default active poll

## Building for Production

```bash
npm run build
```

The built files will be in the `dist` directory.

## Development

### Code Structure

- `/src/components`: Reusable UI components
- `/src/contexts`: React context for state management (User and Pool contexts)
- `/src/pages`: Page components for different routes
- `/src/services`: API services and data fetching logic
- `/src/config.ts`: Application configuration