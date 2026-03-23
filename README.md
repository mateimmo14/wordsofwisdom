# Words of Wisdom

A simple quote-sharing app built with Go and PostgreSQL.

## Setup

1. **Database**
   ```sql
   CREATE DATABASE words_of_wisdom;
   CREATE TABLE wisdoms (
       id SERIAL PRIMARY KEY,
       data TEXT NOT NULL,
       author VARCHAR(255) DEFAULT 'Anonymous',
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

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