package queries

const CreateUserTableQuery = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP	
	)`

const CreateUserQuery = `INSERT INTO users (name, email, password_hash) VALUES (?, ?, ?)`

const ListUsersQuery = `SELECT * FROM users`

const ExistsUserByEmailQuery = `SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)`

const FindUserByEmailQuery = `SELECT * FROM users WHERE email = ?`
