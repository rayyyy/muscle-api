-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id INT(11) AUTO_INCREMENT NOT NULL,
  name VARCHAR(30) NOT NULL ,
  age INT(3) NOT NULL,
  created_at datetime NOT NULL default current_timestamp,
  updated_at datetime NOT NULL default current_timestamp on update current_timestamp,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;