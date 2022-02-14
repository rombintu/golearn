CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    account VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL,
    role VARCHAR(10) NOT NULL DEFAULT "user"
);

INSERT INTO users (id, account, password, role) VALUES (1, "admin", "admin", "admin");