# Fly.io Deployment Guide for FALTOR Backend

## Prerequisites
- Go 1.21+ installed
- Fly.io account (free tier available)
- Gmail app password ready

## Step 1: Install Fly CLI

### Linux/WSL:
```bash
curl -L https://fly.io/install.sh | sh
```

### Mac:
```bash
brew install flyctl
```

### Windows (PowerShell):
```powershell
iwr https://fly.io/install.ps1 -useb | iex
```

After installation, add to your PATH and restart terminal.

## Step 2: Login to Fly.io

```bash
flyctl auth login
```

This will open a browser to log in/sign up.

## Step 3: Initialize Your App

Navigate to the backend directory:
```bash
cd /home/jones/projects/faltordrafting/backend
```

Initialize the Fly app:
```bash
flyctl launch
```

When prompted:
- **App name**: Choose a name like `faltordrafting-backend` (or let it auto-generate)
- **Organization**: Select your personal org
- **Region**: Choose closest to you (e.g., `iad` for US East)
- **PostgreSQL**: **No** (we don't need a database)
- **Redis**: **No** (we don't need Redis)
- **Deploy now**: **No** (we need to set secrets first)

This creates a `fly.toml` configuration file.

## Step 4: Set Environment Secrets

Set your email configuration as secrets (NOT environment variables):

```bash
flyctl secrets set SMTP_HOST=smtp.gmail.com
flyctl secrets set SMTP_PORT=587
flyctl secrets set SMTP_USER=faltorco.ben@gmail.com
flyctl secrets set SMTP_PASSWORD="xudj azrh kuvh tsuz"
flyctl secrets set TO_EMAIL=faltorco.ben@gmail.com
```

**Note**: Secrets are encrypted and not visible in logs.

## Step 5: Deploy

```bash
flyctl deploy
```

This will:
- Build your Go app
- Create a Docker image
- Deploy to Fly.io
- Start your app

## Step 6: Get Your App URL

```bash
flyctl info
```

Look for the **Hostname** - it will be something like:
`faltordrafting-backend.fly.dev`

Your full API URL will be:
`https://faltordrafting-backend.fly.dev`

## Step 7: Test Your API

```bash
curl https://your-app.fly.dev/api/health
```

Should return: `{"status":"ok"}`

Test the contact endpoint:
```bash
curl -X POST https://your-app.fly.dev/api/contact \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "phone": "555-123-4567",
    "location": "Test City",
    "projectType": "Custom house plan",
    "timeline": "ASAP",
    "budget": "100000",
    "message": "This is a test message"
  }'
```

You should receive an email at `faltorco.ben@gmail.com`!

## Step 8: Update Vue Frontend

Once deployed, update the Vue frontend with your Fly.io URL.

In `vue-frontend/src/components/ContactSection.vue`, replace:
```javascript
const apiUrl = import.meta.env.VITE_API_URL || 'https://your-railway-app.railway.app'
```

With your actual Fly.io URL:
```javascript
const apiUrl = import.meta.env.VITE_API_URL || 'https://faltordrafting-backend.fly.dev'
```

## Useful Commands

### View logs:
```bash
flyctl logs
```

### Check app status:
```bash
flyctl status
```

### SSH into your app:
```bash
flyctl ssh console
```

### Scale your app (if needed):
```bash
flyctl scale count 1
```

### View secrets:
```bash
flyctl secrets list
```

### Update a secret:
```bash
flyctl secrets set SMTP_PASSWORD="new-password"
```

## Troubleshooting

### App not responding?
```bash
flyctl status
flyctl logs
```

### Email not sending?
Check logs for SMTP errors:
```bash
flyctl logs | grep -i smtp
```

### Need to redeploy?
```bash
flyctl deploy --force
```

## Free Tier Limits

- 3 shared-cpu VMs with 256MB RAM
- 160GB outbound data transfer
- Plenty for a contact form! You can handle thousands of submissions per month.

## Monitoring

Check your app dashboard:
```bash
flyctl dashboard
```

Or visit: https://fly.io/dashboard

---

That's it! Your backend is now live on Fly.io for free! 🚀
