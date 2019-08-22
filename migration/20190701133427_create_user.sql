-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id INT(11) AUTO_INCREMENT NOT NULL,
  uid VARCHAR(128) NOT NULL,
  email VARCHAR(255) NOT NULL,
  nickname VARCHAR(50) NOT NULL,
  image TEXT NOT NULL,
  birthday DATE,
  sex ENUM('unspecified', 'male', 'female') DEFAULT 'unspecified',
  status ENUM('active', 'canceled', 'freezing') DEFAULT 'active',
  created_at DATETIME NOT NULL DEFAULT current_timestamp,
  updated_at DATETIME NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted_at DATETIME,
  PRIMARY KEY(id),
  UNIQUE (uid)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;