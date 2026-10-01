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