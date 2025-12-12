# FALTOR Consulting Backend

## Environment Variables

Create a `.env` file in the go-backend directory with your email configuration:

```
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
TO_EMAIL=faltorco.ben@gmail.com
```

## Gmail Setup

1. Enable 2-Factor Authentication on your Gmail account
2. Generate an App Password for this application
3. Use the App Password (not your regular password) in the SMTP_PASSWORD field

## Running the Backend

```bash
cd go-backend
go mod tidy
go run main.go
```

The server will start on http://localhost:8080

## API Endpoints

- `POST /api/contact` - Submit contact form
- `GET /api/health` - Health check

## Frontend Integration

The Vue frontend is configured to send form data to this backend API.