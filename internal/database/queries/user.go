package queries

const CreateUserQuery = `INSERT INTO users (name, email) VALUES (?, ?)`

const ListUsersQuery = `SELECT * FROM users`
