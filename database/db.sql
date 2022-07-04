CREATE DATABASE IF NOT EXISTS goclidb;

USE goclidb;

CREATE TABLE
    IF NOT EXISTS tasks (
        id INT NOT NULL AUTO_INCREMENT,
        title VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        PRIMARY KEY (id)
    );