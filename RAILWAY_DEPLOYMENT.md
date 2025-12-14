# Railway Deployment Guide

## Step 1: Set up Gmail App Password

1. Go to your Google Account: https://myaccount.google.com/
2. Click on **Security** in the left sidebar
3. Scroll down and enable **2-Step Verification** (if not already enabled)
4. After enabling 2FA, go back to Security
5. Under "2-Step Verification", click on **App passwords**
6. Select **Mail** as the app and **Other** as the device
7. Name it "FALTOR Backend" 
8. Click **Generate**
9. **Copy the 16-character password** - you'll need this for Railway

## Step 2: Deploy to Railway

1. Go to https://railway.app and sign up/login with GitHub
2. Click **"New Project"**
3. Select **"Deploy from GitHub repo"**
4. Connect your GitHub account if needed
5. Select the **faltordrafting** repository
6. Railway will detect the Go app automatically
7. Click **"Deploy Now"**

## Step 3: Configure Environment Variables

1. Once deployed, click on your service
2. Go to the **Variables** tab
3. Click **"+ New Variable"** and add each of these:

```
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=faltorco.ben@gmail.com
SMTP_PASSWORD=<paste-the-16-char-app-password>
TO_EMAIL=faltorco.ben@gmail.com
```

4. Click **Save** and Railway will redeploy

## Step 4: Get Your Railway URL

1. Go to the **Settings** tab
2. Scroll to **Domains**
3. Click **"Generate Domain"**
4. Copy the URL (e.g., `https://faltordrafting-backend-production.up.railway.app`)

## Step 5: Update Vue Frontend

1. Create `.env.production` in `vue-frontend/` folder:
```
VITE_API_URL=https://your-railway-url.railway.app
```

2. Or update the code in `ContactSection.vue` line with your Railway URL:
```javascript
const apiUrl = import.meta.env.VITE_API_URL || 'https://your-railway-url.railway.app'
```

## Step 6: Test

1. Rebuild and deploy your Vue frontend to Netlify
2. Fill out the contact form on your site
3. Check if email arrives at `faltorco.ben@gmail.com`

## Troubleshooting

### Check Railway Logs
1. In Railway, click on your service
2. Go to **Deployments** tab
3. Click on the latest deployment
4. View logs to see any errors

### Test the API directly
```bash
curl -X POST https://your-railway-url.railway.app/api/contact \
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

You should get: `{"status":"success","message":"Your message has been sent successfully!"}`
