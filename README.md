# Calculator App

A full-stack calculator application with React frontend and Go backend.

## Setup Instructions

### Prerequisites
- Go (v1.26 or higher)
- Node.js (v22 or higher)
- npm (v11 or higher)

### Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd calculator-api
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Start the server:
   ```bash
   go run cmd/api/main.go
   ```

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd calculator-frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

### Docker Setup
Alternatively, you can build and run the application using Docker:

1. Build the Docker image:
   ```bash
   docker build -t calculator-app .
   ```

2. Run the container:
   ```bash
   docker run -p 8000:8000 calculator-app
   ```

The application will be available at http://localhost:8000

## API Usage

The calculator exposes the following API endpoint:

### POST /api/v1/calculator
Performs mathematical operations.

#### Request Body
```json
{
  "operation": "add|subtract|multiply|divide|power|sqrt|percentage",
  "a": number,
  "b": number (optional)
}
```

#### Response
```json
{
  "result": number
}
```

#### Error Response
```json
{
  "error": "error message"
}
```

## Design Rationale

### Backend
- Built with Go and Gin framework for performance and type safety
- RESTful API design with clear request/response structure
- Modular code organization with separate handler, service, and model layers
- Proper error handling with specific error types for different failure cases
- Support for basic and advanced mathematical operations
- Input validation and sanitization

### Frontend
- Built with React and TypeScript for type safety
- Responsive design using CSS
- Clean component structure with Calculator.tsx as the main component
- Error handling for API calls that displays "Error" in the calculator display
- Integration with backend API

### Architecture
- Separation of concerns between frontend and backend
- Clear API contract between client and server
- Error handling at both frontend and backend levels
- Graceful shutdown handling for the server
- Proper logging implementation

## Useful Prompts

### Architecture
- Act as a senior software architect. Design a clean monorepo structure for a React + TypeScript frontend and Go + Gin backend. Keep frontend and backend independent, testable, and easy to Dockerize. Avoid over-engineering.
### Backend
- Design an idiomatic Go + Gin project structure for a calculator REST API. Separate handlers, business logic, validation, and routing. Keep it simple and testable.
- Create a Go testing strategy for a Gin calculator API using table-driven tests. Cover business logic, handlers, validation, and edge cases.
### API
- Design a REST API contract for a calculator supporting add, subtract, multiply, divide, percentage, power, and square root. Define endpoints, request/response JSON, validation, and error handling.
- Show how to mock fetch with Vitest for a React app consuming a REST API. Include successful responses, HTTP errors, and network failures. Do not use MSW.
### Frontend
- Design the React state and component logic for a calculator supporting basic and advanced operations, decimals, clear, sign toggle, and API-based calculations. Keep the implementation simple.
### Unit test
- Create a testing strategy for a React + TypeScript calculator using Vitest and React Testing Library. Cover user interactions, API success/error responses, validation, and asynchronous state updates.
- Set up frontend coverage with Vitest and backend coverage with Go. Generate HTML reports and document the commands in the README.
### Docker
- Design a multi-stage Dockerfile for a React/Vite frontend and Go/Gin backend in the same repository. Build both and serve the React static files from Go in a minimal production image.
