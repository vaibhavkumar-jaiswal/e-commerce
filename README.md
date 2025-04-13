# 🛍️ Go E-Commerce Backend

A full-featured e-commerce backend built with **Go**, using the Gin web framework and GORM ORM. Includes RESTful APIs, JWT authentication, Redis caching, PostgreSQL database, Swagger docs, input validation, logging, SMTP support, and more.

---

## 🚀 Tech Stack

- **Go**
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) - ORM for Golang
- PostgreSQL - Primary database
- Redis - For caching and sessions
- SMTP - Email functionality (order confirmation, etc.)
- JWT - Secure authentication
- [golangci-lint](https://golangci-lint.run/) - Linter for clean code
- [Swagger](https://swagger.io/) - Auto-generated API docs
- Gin Validator - Input validation

---

## 📂 Project Structure

```
.
├── config/             # Configuration loading logic
├── database/           # DB initialization, migrations
├── docs/               # Swagger docs and OpenAPI definitions
├── middleware/         # Custom middleware like auth, logging, etc.
├── modules/            # Feature modules (user, product, order, etc.)
├── services/           # Service integration like SMTP, SQS, etc
├── models/             # models (tables, struct)
├── utils/              # Helper functions (constants, config load data etc.)
├── bin/                # Compiled binary output
├── .gitignore          # Git ignored files
├── .golangci.yml       # GolangCI Lint configuration
├── config.env          # Main environment config
├── example-config.env  # Sample env file for development
├── go.mod/go.sum       # Go modules
├── main.go             # Main app file
├── nodemon.json        # Nodemon config for hot-reload
├── run.sh              # Shell script to build and run the app
└── README.md           # Project README
```

---

## 🔧 Setup & Run

### 1. Clone the repo

```bash
git clone https://github.com/vaibhavkumar-jaiswal/e-commerce.git
```

### 2. Create `.env` file

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
REDIS_ADDR=localhost:6379
JWT_SECRET=your_jwt_secret
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=your_email@example.com
SMTP_PASSWORD=your_password
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Start the app using run script

If you have [nodemon](https://github.com/remy/nodemon) installed, it will auto-reload on file changes.

Check if nodemon is installed:

```bash
nodemon --version
```

If it's not installed, you can install it globally:

```bash
npm install -g nodemon
```

Then start the project using:

```bash
nodemon
```

If you don't want to install `nodemon`, run the following command

```bash
bash run.sh
```

> ℹ️ `run.sh` handles generating Swagger docs, linting, building, and running the project.

---

## 🧪 API Docs

After starting the server, access Swagger UI:

```
http://localhost:8080/api-docs
```

---

## Linting

```bash
golangci-lint run ./...
```

---

## Features

- User registration, email verification & login with JWT
- Product CRUD
- Order & Cart management
- Email notifications via SMTP
- Route-level JWT auth middleware
- Validation using `binding:"required"`
- Swagger API docs auto-generated
- Clean code with `golangci-lint`

---

## Email Example

When a user places an order, an email confirmation is sent via configured SMTP settings.

---

## TODO (Optional)

- Payment gateway integration (Stripe/Razorpay/etc.)
- Admin panel for product management
- Docker support for full-stack dev

---

## 🙌 Contributing

Pull requests are welcome! For major changes, open an issue first to discuss what you’d like to change.

---

## 👨‍💻 Author

- **Vaibhavkumar Jaiswal** - [@vaibhavkumar-jaiswal](https://github.com/vaibhavkumar-jaiswal)