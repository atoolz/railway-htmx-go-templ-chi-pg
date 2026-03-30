<p align="center">
  <img src="assets/banner.png" alt="HTMX + Go + PostgreSQL Railway Template" width="100%">
</p>

<p align="center">
  <strong>Production-ready HTMX starter with Go, Templ, Chi, and PostgreSQL. One-click deploy to Railway.</strong>
</p>

<p align="center">
  <a href="https://railway.com/deploy/htmx-go-chi-postgres">
    <img src="https://railway.com/button.svg" alt="Deploy on Railway">
  </a>
</p>

<p align="center">
  <a href="https://github.com/atoolz/railway-htmx-go-templ-chi-pg/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/atoolz/railway-htmx-go-templ-chi-pg?style=flat-square&color=00c9a7" alt="License">
  </a>
  <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square" alt="Go 1.25">
  <img src="https://img.shields.io/badge/HTMX-2.0.7-3366CC?style=flat-square" alt="HTMX 2.0.7">
  <img src="https://img.shields.io/badge/Chi-v5.2-00c9a7?style=flat-square" alt="Chi v5.2">
  <img src="https://img.shields.io/badge/PostgreSQL-16-336791?style=flat-square" alt="PostgreSQL">
</p>

<br>

## Deploy and Host HTMX + Go + PostgreSQL Starter on Railway

HTMX + Go + PostgreSQL Starter is a production-ready template for building hypermedia-driven web applications. It combines HTMX for dynamic interactions without JavaScript, Go with Chi and Templ for type-safe server-side rendering, and PostgreSQL for persistent storage. Deploy in one click and start building.

### About Hosting HTMX + Go + PostgreSQL Starter

The template deploys as a single Go binary built via multi-stage Dockerfile (~15MB final image, ~31s build time). It connects to a PostgreSQL instance over Railway's private network using the pgx driver with connection pooling. Database migrations run automatically on startup, creating the required tables without manual SQL. The application reads `DATABASE_URL` from environment variables and listens on the port assigned by Railway. A health check endpoint at `/health` pings the database and returns status, ensuring Railway can verify deployments before routing traffic. Tailwind CSS and HTMX load via CDN, so there is no frontend build step.

### Common Use Cases

- Rapid prototyping of server-rendered web apps that need dynamic interactions without the complexity of React, Vue, or other JavaScript frameworks
- Building internal tools, admin panels, and CRUD dashboards where Go's performance and type safety matter but a full SPA is overkill
- Learning HTMX patterns (hx-get, hx-post, hx-swap, hx-target) with a working reference application backed by a real database

### Dependencies for Hosting

- A Railway PostgreSQL database instance (added via the Railway dashboard or CLI)
- The `DATABASE_URL` environment variable set to `${{Postgres.DATABASE_URL}}` on the web service

#### Deployment Dependencies

- [HTMX documentation](https://htmx.org/docs/)
- [Go Chi router](https://go-chi.io/)
- [Templ templating](https://templ.guide)
- [pgx PostgreSQL driver](https://github.com/jackc/pgx)

### Why Deploy on Railway?

Railway is a singular platform to deploy your infrastructure stack. Railway will host your infrastructure so you don't have to deal with configuration, while allowing you to vertically and horizontally scale it.

By deploying HTMX + Go + PostgreSQL Starter on Railway, you are one step closer to supporting a complete full-stack application with minimal burden. Host your servers, databases, AI agents, and more on Railway.

<br>

## What's Inside

A complete, working Todo application that demonstrates the full HTMX + Go stack. No JavaScript frameworks, no build steps for the frontend, no client-side state management. Just HTML over the wire.

| Layer | Technology | Role |
|-------|-----------|------|
| **Frontend** | HTMX 2.0.7 + Tailwind CSS | Hypermedia-driven interactions via CDN |
| **Templating** | Templ v0.3 | Type-safe Go templates with IDE support |
| **Routing** | Chi v5.2 | Lightweight, idiomatic HTTP router for Go |
| **Database** | PostgreSQL + pgx v5.9 | Connection pooling, prepared statements |
| **Deploy** | Dockerfile (multi-stage) | 31s build, ~15MB final image |

<br>

## Why This Stack

**HTMX** lets you build dynamic UIs by returning HTML fragments from the server instead of JSON. No React, no Vue, no bundle. Your Go backend renders HTML, HTMX swaps it into the DOM. The browser does what browsers do best.

**Templ** gives you type-safe, composable HTML components written in Go. Compile-time checking catches broken templates before they reach production. IDE autocompletion works because components are just Go functions.

**Chi** is a lightweight, idiomatic HTTP router for Go built on top of `net/http`. It uses standard `http.Handler` and `http.HandlerFunc` signatures, so your handlers work with any Go HTTP middleware. No framework lock-in, just composable routing with middleware support.

**PostgreSQL** needs no introduction. The pgx driver gives you connection pooling, prepared statements, and direct access to PostgreSQL-specific features without an ORM in the way.

<br>

## Project Structure

```
.
├── cmd/server/
│   └── main.go              # Server setup, routes, middleware
├── internal/
│   ├── database/
│   │   └── database.go      # pgx connection pool + auto-migration
│   ├── handlers/
│   │   └── handlers.go      # HTTP handlers (CRUD + health check)
│   └── models/
│       └── models.go        # Data structures
├── templates/
│   ├── layout.templ          # Base HTML with HTMX + Tailwind CDN
│   ├── home.templ            # Home page composition
│   └── components.templ      # HTMX-powered interactive components
├── Dockerfile                # Multi-stage build (golang -> alpine)
└── go.mod
```

<br>

## HTMX Patterns Demonstrated

The Todo app showcases core HTMX patterns you'll use in real projects:

- **`hx-post` + `hx-target` + `hx-swap="afterbegin"`** - Create items and prepend to list without page reload
- **`hx-patch` + `hx-target` (self)** - Toggle completion with in-place swap
- **`hx-delete` + `hx-swap="outerHTML"`** - Remove items with smooth DOM removal
- **`hx-on::after-request`** - Reset form after successful submission
- **Health check endpoint** - `/health` returns JSON status with DB ping

<br>

## Deploy to Railway

Click the button above or:

1. Fork this repo
2. Create a new project on [Railway](https://railway.com)
3. Add a **PostgreSQL** database
4. Add a service from your forked repo
5. Set the variable: `DATABASE_URL` = `${{Postgres.DATABASE_URL}}`
6. Railway auto-detects the Dockerfile and deploys

The app auto-migrates the database on startup. No manual SQL needed.

<br>

## Local Development

```bash
# Prerequisites: Go 1.25+, PostgreSQL running locally

# Install templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Clone and run
git clone https://github.com/atoolz/railway-htmx-go-templ-chi-pg.git
cd railway-htmx-go-templ-chi-pg

# Set your database URL
export DATABASE_URL="postgres://user:pass@localhost:5432/mydb?sslmode=disable"

# Generate templ files and run
templ generate
go run ./cmd/server
```

Open [http://localhost:8080](http://localhost:8080).

<br>

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DATABASE_URL` | Yes | - | PostgreSQL connection string |
| `PORT` | No | `8080` | Server port (Railway sets this automatically) |

<br>

## Part of the HTMX Railway Collection

This is one of 15 HTMX starter templates covering different backend stacks, all following the same pattern and ready for Railway deployment:

| Stack | Status |
|-------|--------|
| Bun + Elysia | Coming soon |
| .NET + Razor | Coming soon |
| Elixir + Phoenix | Coming soon |
| **Go + Chi** | **This repo** |
| Go + Echo | [Live](https://github.com/atoolz/railway-htmx-go-templ-echo-pg) |
| Go + Fiber | [Live](https://github.com/atoolz/railway-htmx-go-templ-fiber-pg) |
| Java + Spring Boot (MySQL) | [Live](https://github.com/atoolz/railway-htmx-java-spring-thymeleaf-mysql) |
| Java + Spring Boot (PostgreSQL) | [Live](https://github.com/atoolz/railway-htmx-java-spring-thymeleaf-pg) |
| Node + Express | [Live](https://github.com/atoolz/railway-htmx-node-express-ejs-pg) |
| Node + Hono | [Live](https://github.com/atoolz/railway-htmx-node-hono-jsx-pg) |
| PHP + Laravel | [Live](https://github.com/atoolz/railway-htmx-php-laravel-mysql) |
| Python + Django | [Live](https://github.com/atoolz/railway-htmx-python-django-pg) |
| Python + FastAPI | [Live](https://github.com/atoolz/railway-htmx-python-fastapi-jinja2-pg) |
| Ruby + Rails 8 | [Live](https://github.com/atoolz/railway-htmx-ruby-rails8-pg) |
| Rust + Axum + Askama | [Live](https://github.com/atoolz/railway-htmx-rust-axum-askama-pg) |

<br>

## License

[MIT](LICENSE)

---

<p align="center">
  <sub>Built by <a href="https://github.com/atoolz">AToolZ</a> for the HTMX community</sub>
</p>
