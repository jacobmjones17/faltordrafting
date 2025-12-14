# Faltor Drafting Backend

Go backend server for handling contact form submissions and sending emails.

## Setup for Fly.io

See `FLY_DEPLOYMENT.md` for detailed deployment instructions.

## Environment Variables

Required environment variables (set in Fly.io secrets):

```
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=faltorco.ben@gmail.com
SMTP_PASSWORD=your-gmail-app-password
TO_EMAIL=faltorco.ben@gmail.com
```

## API Endpoints

- `POST /api/contact` - Submit contact form
- `GET /api/health` - Health check

## Local Development

```bash
# Set environment variables
export SMTP_HOST=smtp.gmail.com
export SMTP_PORT=587
export SMTP_USER=faltorco.ben@gmail.com
export SMTP_PASSWORD=your-app-password
export TO_EMAIL=faltorco.ben@gmail.com

# Run the server
go run main.go
```

Server will run on `http://localhost:8080`
