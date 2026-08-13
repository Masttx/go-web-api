package queries

const CreateAreaUserQuery = `CREATE TABLE IF NOT EXISTS area_users (area_id INTEGER REFERENCES area(id), user_id INTEGER REFERENCES user(id))`
