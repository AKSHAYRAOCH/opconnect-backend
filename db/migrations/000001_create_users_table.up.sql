CREATE TABLE IF NOT EXISTS Users(
    Id text primary key,
    Email text unique NOT NULL,
    Password text NOT NULL,
    Role text NOT NULL
);


