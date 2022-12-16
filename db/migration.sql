CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION pgcrypto;
DROP TABLE IF EXISTS user_role;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;


CREATE TABLE users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    address VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE roles (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    permissions VARCHAR(255),
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE user_role (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id uuid,
    role_id uuid,
    status BOOLEAN
);

ALTER TABLE user_role ADD CONSTRAINT user_pk FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_role ADD CONSTRAINT role_pk FOREIGN KEY (role_id) REFERENCES roles (id);
