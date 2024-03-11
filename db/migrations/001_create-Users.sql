CREATE TABLE IF NOT EXISTS Users(
    Id text primary key,
    Email text unique NOT NULL,
    Password text NOT NULL,
    Role text NOT NULL,
    check(Role in ('Student','Mentor','Admin','Recruiter'))
);


---- create above / drop below ----

DROP TABLE IF EXISTS Users;