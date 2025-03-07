CREATE TABLE users(
    user_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username varchar(50) NOT NULL UNIQUE,
    initials varchar(3) NOT NULL,
    nickname varchar(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


