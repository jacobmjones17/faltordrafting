const nodemailer = require('nodemailer');
const multipart = require('parse-multipart-data');

exports.handler = async (event, context) => {
  // Only allow POST
  if (event.httpMethod !== 'POST') {
    return {
      statusCode: 405,
      body: JSON.stringify({ error: 'Method not allowed' })
    };
  }

  let data = {};
  let files = [];

  // Check if this is multipart form data (with files) or JSON
  const contentType = event.headers['content-type'] || event.headers['Content-Type'] || '';
  
  if (contentType.includes('multipart/form-data')) {
    // Parse multipart form data
    try {
      const boundary = contentType.split('boundary=')[1];
      const parts = multipart.parse(Buffer.from(event.body, 'base64'), boundary);
      
      // Extract form fields and files
      parts.forEach(part => {
        if (part.filename) {
          // This is a file
          files.push({
            filename: part.filename,
            content: part.data,
            contentType: part.type
          });
        } else {
          // This is a form field
          data[part.name] = part.data.toString();
        }
      });
    } catch (error) {
      console.error('Error parsing multipart data:', error);
      return {
        statusCode: 400,
        body: JSON.stringify({ error: 'Invalid form data' })
      };
    }
  } else {
    // Parse JSON body
    try {
      data = JSON.parse(event.body);
    } catch (error) {
      return {
        statusCode: 400,
        body: JSON.stringify({ error: 'Invalid request body' })
      };
    }
  }

  // Validate required fields
  const { name, email, message } = data;
  if (!name || !email || !message) {
    return {
      statusCode: 400,
      body: JSON.stringify({ error: 'Name, email, and message are required' })
    };
  }

  // Get SMTP credentials from environment variables
  const smtpConfig = {
    host: process.env.SMTP_HOST,
    port: parseInt(process.env.SMTP_PORT),
    secure: false, // true for 465, false for other ports
    auth: {
      user: process.env.SMTP_USER,
      pass: process.env.SMTP_PASSWORD
    }
  };

  // Create transporter
  const transporter = nodemailer.createTransport(smtpConfig);

  // Email content
  const mailOptions = {
    from: process.env.SMTP_USER,
    to: process.env.TO_EMAIL,
    subject: `🏗️ New Contact Request from ${name}`,
    text: `
New contact form submission from FALTOR Consulting website:

Name: ${data.name}
Email: ${data.email}
Phone: ${data.phone || 'Not provided'}
Location: ${data.location || 'Not provided'}
Project Type: ${data.projectType || 'Not provided'}
Timeline: ${data.timeline || 'Not provided'}
Budget: ${data.budget || 'Not provided'}

Message:
${data.message}

---
This email was sent from the FALTOR Consulting contact form.
    `,
    html: `
      <!DOCTYPE html>
      <html>
      <head>
        <style>
          body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f4f4f4;
          }
          .email-container {
            background-color: #ffffff;
            border-radius: 8px;
            padding: 30px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
          }
          .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 20px;
            border-radius: 8px 8px 0 0;
            margin: -30px -30px 30px -30px;
          }
          .header h1 {
            margin: 0;
            font-size: 24px;
            font-weight: 600;
          }
          .header p {
            margin: 5px 0 0 0;
            opacity: 0.9;
            font-size: 14px;
          }
          .info-grid {
            display: table;
            width: 100%;
            margin: 20px 0;
          }
          .info-row {
            display: table-row;
          }
          .info-label {
            display: table-cell;
            padding: 12px 20px 12px 0;
            font-weight: 600;
            color: #667eea;
            width: 140px;
            vertical-align: top;
          }
          .info-value {
            display: table-cell;
            padding: 12px 0;
            color: #333;
            vertical-align: top;
          }
          .message-section {
            background-color: #f8f9fa;
            border-left: 4px solid #667eea;
            padding: 20px;
            margin: 20px 0;
            border-radius: 4px;
          }
          .message-section h3 {
            margin-top: 0;
            color: #667eea;
            font-size: 18px;
          }
          .message-content {
            color: #555;
            white-space: pre-wrap;
            word-wrap: break-word;
          }
          .footer {
            margin-top: 30px;
            padding-top: 20px;
            border-top: 2px solid #e9ecef;
            text-align: center;
            color: #6c757d;
            font-size: 13px;
          }
          .badge {
            display: inline-block;
            padding: 4px 12px;
            background-color: #e7f0ff;
            color: #667eea;
            border-radius: 12px;
            font-size: 12px;
            font-weight: 600;
            margin-top: 5px;
          }
        </style>
      </head>
      <body>
        <div class="email-container">
          <div class="header">
            <h1>🏗️ New Contact Request</h1>
            <p>FALTOR Consulting Website</p>
          </div>
          
          <div class="info-grid">
            <div class="info-row">
              <div class="info-label">👤 Name:</div>
              <div class="info-value">${data.name}</div>
            </div>
            <div class="info-row">
              <div class="info-label">📧 Email:</div>
              <div class="info-value"><a href="mailto:${data.email}" style="color: #667eea; text-decoration: none;">${data.email}</a></div>
            </div>
            ${data.phone ? `
            <div class="info-row">
              <div class="info-label">📱 Phone:</div>
              <div class="info-value"><a href="tel:${data.phone}" style="color: #667eea; text-decoration: none;">${data.phone}</a></div>
            </div>
            ` : ''}
            ${data.location ? `
            <div class="info-row">
              <div class="info-label">📍 Location:</div>
              <div class="info-value">${data.location}</div>
            </div>
            ` : ''}
            ${data.projectType ? `
            <div class="info-row">
              <div class="info-label">🏠 Project Type:</div>
              <div class="info-value"><span class="badge">${data.projectType}</span></div>
            </div>
            ` : ''}
            ${data.timeline ? `
            <div class="info-row">
              <div class="info-label">⏰ Timeline:</div>
              <div class="info-value">${data.timeline}</div>
            </div>
            ` : ''}
            ${data.budget ? `
            <div class="info-row">
              <div class="info-label">💰 Budget:</div>
              <div class="info-value">$${data.budget}</div>
            </div>
            ` : ''}
          </div>

          <div class="message-section">
            <h3>💬 Message</h3>
            <div class="message-content">${data.message}</div>
          </div>

          ${files.length > 0 ? `
          <div style="margin: 20px 0; padding: 15px; background-color: #f0f8ff; border-radius: 4px; border-left: 4px solid #667eea;">
            <h3 style="margin-top: 0; color: #667eea; font-size: 16px;">📎 Attachments (${files.length})</h3>
            <ul style="margin: 10px 0; padding-left: 20px;">
              ${files.map(f => `<li style="color: #555; margin: 5px 0;">${f.filename}</li>`).join('')}
            </ul>
          </div>
          ` : ''}

          <div class="footer">
            <p>✉️ This email was sent from the FALTOR Consulting contact form</p>
            <p style="margin-top: 10px; font-size: 12px;">Received on ${new Date().toLocaleString('en-US', { 
              weekday: 'long', 
              year: 'numeric', 
              month: 'long', 
              day: 'numeric',
              hour: '2-digit',
              minute: '2-digit'
            })}</p>
          </div>
        </div>
      </body>
      </html>
    `
  };

  // Add file attachments if any
  if (files.length > 0) {
    mailOptions.attachments = files.map(file => ({
      filename: file.filename,
      content: file.content,
      contentType: file.contentType
    }));
  }

  // Send email
  try {
    // Verify SMTP configuration exists
    if (!smtpConfig.host || !smtpConfig.auth.user || !smtpConfig.auth.pass) {
      console.error('SMTP configuration missing:', {
        hasHost: !!smtpConfig.host,
        hasUser: !!smtpConfig.auth.user,
        hasPass: !!smtpConfig.auth.pass
      });
      throw new Error('Email configuration is incomplete. Please check environment variables.');
    }

    await transporter.sendMail(mailOptions);
    
    return {
      statusCode: 200,
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Headers': 'Content-Type',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        status: 'success',
        message: 'Your message has been sent successfully!'
      })
    };
  } catch (error) {
    console.error('Email error:', error);
    console.error('Error stack:', error.stack);
    
    return {
      statusCode: 500,
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        error: 'Failed to send email',
        details: error.message,
        hint: 'Check Netlify function logs for more details'
      })
    };
  }
};
