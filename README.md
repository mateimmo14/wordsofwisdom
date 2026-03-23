# Words of Wisdom

A simple quote-sharing app built with Go and PostgreSQL.

## Setup

1. **Database**
   
   Migration scripts are located in `sql/schemes/`. To set up the database, run:
   ```bash
   goose up
   ```
   This will create the necessary database and tables.

2. **Environment**
   ```bash
   export DATABASE_URL="postgres://user:password@localhost:5432/words_of_wisdom"
   ```

3. **Run**
   ```bash
   go run main.go handlers.go
   ```
   Server runs on `http://localhost:8080`

## API

**GET /** - View quotes  
**POST /** - Submit quote
```json
{
  "wisdom": "Your quote here",
  "author": "Your name"
}
```

## Stack

- Go 1.21+
- PostgreSQL
- sqlc for database code generation.
- Goose for database migrations