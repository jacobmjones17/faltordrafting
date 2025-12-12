#!/bin/bash
# Load environment variables and start the server
export SMTP_USERNAME="jacobmjones17@gmail.com"
export SMTP_PASSWORD="your-app-password-here"  # Replace with actual App Password

echo "🚀 Starting server with environment variables..."
go run main.go