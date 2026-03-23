# Words of Wisdom

A simple quote-sharing app built with Go and PostgreSQL. Share your favorite quotes with the world!

## Live Example

Check out the app in action: **[wordsofwisdom.duckdns.org](http://wordsofwisdom.duckdns.org)**

## Features

- 📝 Share quotes with author attribution
- 🔄 Real-time quote management
- 💾 Persistent storage with PostgreSQL
- 🚀 Built with Go for high performance
- 🛠️ Database migrations with Goose

## Setup

### 1. Build Goose Binary

First, install the Goose database migration tool:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### 2. Database Setup

Migration scripts are located in `sql/schemes/`. To set up the database, run:
```bash
goose postgres <db-URL> up
```

Replace `<db-URL>` with your PostgreSQL connection string, e.g.:
```bash
goose postgres "postgres://user:password@localhost:5432/words_of_wisdom" up
```

This will create the necessary database and tables.

### 3. Environment

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/words_of_wisdom"
```

### 4. Run the Server

```bash
go build -o wordsofwisdom
./wordsofwisdom
```

Or directly with Go:
```bash
go run main.go handlers.go
```

Server runs on `http://localhost:8080`

## API

### Get All Quotes
**GET /** - View all quotes
```bash
curl http://localhost:8080/
```

### Submit a New Quote
**POST /** - Submit a new quote
```bash
curl -X POST http://localhost:8080/ \
  -H "Content-Type: application/json" \
  -d '{
    "wisdom": "The only way to do great work is to love what you do.",
    "author": "Steve Jobs"
  }'
```

## Stack

- **Go 1.21+** - Fast, efficient backend
- **PostgreSQL** - Reliable database
- **sqlc** - Type-safe database code generation
- **Goose** - Database migration management

## Development

Pull requests and issues are welcome! Feel free to contribute.