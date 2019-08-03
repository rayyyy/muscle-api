-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE user_details (
  id INT(11) AUTO_INCREMENT NOT NULL,
  user_id INT(11) NOT NULL,
  short_message VARCHAR(255),
  height DECIMAL ( 5, 2 ),
  weight DECIMAL ( 5, 2 ),
  created_at datetime NOT NULL default current_timestamp,
  updated_at datetime NOT NULL default current_timestamp on update current_timestamp,
  deleted_at DATETIME,
  PRIMARY KEY(id)
  -- FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE user_details;
