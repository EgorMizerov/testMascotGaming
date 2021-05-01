CREATE TABLE users
(
    id uuid,
    username varchar(255) UNIQUE,
    password_hash varchar(255),
    refresh_token varchar(255),
    balance decimal DEFAULT 0,
    PRIMARY KEY (id),
    CHECK ( balance >= 0 )
);

CREATE INDEX users_refresh_token_idx ON users (refresh_token);

CREATE TABLE banks
(
    id uuid,
    user_id uuid,
    currency varchar(5),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- TODO: Нужна ли роль учатника банка?
CREATE TABLE banks_users
(
    id uuid,
    user_id uuid,
    bank_id uuid,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    FOREIGN KEY (bank_id)
        REFERENCES banks(id)
        ON DELETE CASCADE
);