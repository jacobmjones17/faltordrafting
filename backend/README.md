# Faltor Drafting Backend

Go backend server for handling contact form submissions and sending emails.

## Setup for Railway

1. Push this code to your GitHub repository
2. Go to [Railway](https://railway.app)
3. Click "New Project" → "Deploy from GitHub repo"
4. Select your repository
5. Railway will auto-detect the Go app and deploy it

## Environment Variables (Set in Railway)

In Railway, go to your service → Variables tab and add:

```
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your-email@gmail.com
SMTP_PASSWORD=your-gmail-app-password
TO_EMAIL=faltorco.ben@gmail.com
```

### Getting Gmail App Password:

1. Go to your Google Account → Security
2. Enable 2-Step Verification
3. Go to App Passwords
4. Generate a new app password for "Mail"
5. Use that password (not your regular Gmail password)

## API Endpoints

- `POST /api/contact` - Submit contact form
- `GET /api/health` - Health check

## Local Development

```bash
# Set environment variables
export SMTP_HOST=smtp.gmail.com
export SMTP_PORT=587
export SMTP_USER=your-email@gmail.com
export SMTP_PASSWORD=your-app-password
export TO_EMAIL=faltorco.ben@gmail.com

# Run the server
go run main.go
```

Server will run on `http://localhost:8080`
