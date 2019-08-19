-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE mentor_plans (
  id INT(11) AUTO_INCREMENT NOT NULL,
  mentor_id INT(11) NOT NULL,
  type ENUM('subscription', 'once') DEFAULT 'once' NOT NULL,
  title VARCHAR(255) NOT NULL,
  detail TEXT NOT NULL NOT NULL,
  price INT NOT NULL,
  is_valid BOOLEAN NOT NULL,
  created_at datetime NOT NULL default current_timestamp,
  updated_at datetime NOT NULL default current_timestamp on update current_timestamp,
  deleted_at DATETIME,
  PRIMARY KEY(id)
  -- FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE mentor_plans;
