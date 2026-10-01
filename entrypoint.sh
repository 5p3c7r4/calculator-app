#!/bin/sh
set -e

# Start the backend API server in the background
echo "Starting backend API server on port 8080..."
./calculator-api
