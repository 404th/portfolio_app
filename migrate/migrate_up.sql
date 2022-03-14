CREATE TABLE IF NOT EXISTS registered_users (
    uuid VARCHAR(250) PRIMARY KEY NOT NULL,
    email VARCHAR(250) UNIQUE,
    username VARCHAR(50) UNIQUE NOT NULL,
    password INT NOT NULL
);