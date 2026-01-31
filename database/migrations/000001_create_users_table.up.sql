CREATE TABLE users (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    nisn VARCHAR(20) UNIQUE,
    full_name VARCHAR(100),
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(225) NOT NULL,
    role ENUM ("admin", "student") DEFAULT "student",
    class_group VARCHAR(20)
);