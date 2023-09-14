CREATE DATABASE Token;

USE Token;

CREATE TABLE tokeninfo (
  id INT PRIMARY KEY AUTO_INCREMENT,
  symbol VARCHAR(50) NOT NULL,
  price DECIMAL(15, 4) NOT NULL,
  source VARCHAR(50) NOT NULL,
  timestamp TIMESTAMP NOT NULL
);