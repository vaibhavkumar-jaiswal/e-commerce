package constants

// .env keys

const LOCAL_ENV = "local"
const DEV_ENV = "dev"
const TEST_ENV = "test"

const DB_USER = "DB_USER"
const DB_PASSWORD = "DB_PASSWORD"
const APP_ENV = "APP_ENV"
const APP_PORT = "APP_PORT"
const REDIS_HOST = "REDIS_HOST"
const REDIS_PORT = "REDIS_PORT"
const LOG_LEVEL = "LOG_LEVEL"
const SMTP_USER = "EMAIL_USER"
const SMTP_PASSWORD = "EMAIL_PASSWORD"
const EMAIL_FROM = "EMAIL_FROM"
const COMPANY_NAME = "COMPANY_NAME"

// const DB_USER = "DB_USER"

const SECRETE_KEY = "JWT_SECRET"

const USER_JWT_CLAIM_KEY = "user_details"
const USER_DATA_CONTEXT_KEY = "logged_in_user_data"
const USER_DATA_OF_SESSION = "user_data_session"

// key for setting response in the gin context
const RESPONSE_DATA_KEY = "response_data"
const RESPONSE_STATUS_KEY = "response_status"

const RATE_LIMIT_PREFIX = "rate_limit_"

const OTP_VERIFICATION_EMAIL_SUBJECT = `%s | Email Verification OTP`

const OTP_VERIFICATION_EMAIL_FORMAT_HTML = `
<!DOCTYPE html>
<html>
<head>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9f9f9;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 20px auto;
            background-color: #ffffff;
            border: 1px solid #ddd;
            border-radius: 8px;
            padding: 20px;
            text-align: center;
        }
        .header {
            background-color: #007bff;
            color: #ffffff;
            padding: 10px 0;
            border-radius: 8px 8px 0 0;
        }
        .otp {
            font-size: 24px;
            font-weight: bold;
            color: #007bff;
            margin: 20px 0;
        }
        .footer {
            font-size: 12px;
            color: #777;
            margin-top: 20px;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            background-color: #007bff;
            color: #ffffff;
            text-decoration: none;
            border-radius: 4px;
            margin-top: 20px;
        }
        .button:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>%s</h1>
        </div>
        <p>Hi %s,</p>
        <p>Thank you for signing up with us! Please use the One-Time Password (OTP) below to verify your email address:</p>
        <div class="otp">%s</div>
        <p>This OTP is valid for the next 10 minutes. Please do not share this code with anyone.</p>
        <p>If you did not request this, please ignore this email or contact our support team at <a href="mailto:support@shopify.com">support@shopify.com</a>.</p>
        <div class="footer">
            © %d Shopify Pvt Ltd. All rights reserved.
        </div>
    </div>
</body>
</html>
`

const OTP_VERIFICATION_EMAIL_FORMAT_TXT = `
Dear %s,

Thank you for signing up with %s! To verify your email address, please use the One-Time Password (OTP) provided below:

### **Your OTP: %s**

This OTP is valid for the next 10 minutes. Please do not share this code with anyone.

To complete your verification, enter this OTP on the %s website or app.

If you did not request this, please ignore this email or contact our support team at support@shopify.com.

Thank you for choosing %s. We're excited to have you on board!

Best regards,  
%s

---

**Disclaimer:**  
This email and any attachments are confidential and intended solely for the recipient. If you are not the intended recipient, please notify us immediately and delete this email.
`

const SHARE_CREDENTIAL_EMAIL_SUBJECT = `Welcome to %s - Your Login Credentials`

const SHARE_CREDENTIAL_EMAIL_FORMAT_HTML = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <style>
      body {
        font-family: Arial, sans-serif;
        line-height: 1.6;
        color: #333;
      }
      .container {
        max-width: 600px;
        margin: 20px auto;
        border: 1px solid #ddd;
        padding: 20px;
        border-radius: 8px;
      }
      .header {
        text-align: center;
        margin-bottom: 20px;
      }
      .footer {
        text-align: center;
        font-size: 12px;
        color: #777;
        margin-top: 20px;
      }
      a {
        color: #007BFF;
        text-decoration: none;
      }
      .highlight {
        font-weight: bold;
        color: #333;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h2>Welcome to <span style="color: #007BFF;">%s</span></h2>
      </div>
      <p>Dear <span class="highlight">%s</span>,</p>
      <p>Congratulations! Your account has been successfully created.</p>
      <p>Here are your login details:</p>
      <ul>
        <li><strong>Login ID:</strong> <span class="highlight">%s</span></li>
        <li><strong>Password:</strong> <span class="highlight">%s</span></li>
      </ul>
      
      <p>For security reasons, we recommend changing your password after your first login.</p>
      <p>If you have any questions or need assistance, please contact us at <a href="mailto:support@shopify.com">support@shopify.com</a> or visit our Help Center.</p>
      <p>Thank you for choosing <span class="highlight">%s</span>!</p>
      <div class="footer">
        <p>&copy; %d Shopify Pvt Ltd. All rights reserved.</p>
      </div>
    </div>
  </body>
</html>
`

const SHARE_CREDENTIAL_EMAIL_FORMAT_TXT = `
Dear %s,

Congratulations! Your account has been successfully created on %s.

Here are your login details:

Login ID: %s  
Password: %s

You can now log in to your account using the credentials provided above.

For security purposes, we recommend changing your password after your first login.

If you have any questions or need further assistance, feel free to contact us at support@shopify.com or visit our Help Center.

Thank you for choosing %s!

Best regards,  
%s
`
