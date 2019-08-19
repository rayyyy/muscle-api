-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE mentors (
  id INT(11) AUTO_INCREMENT NOT NULL,
  user_id INT(11) NOT NULL,
  title VARCHAR(255) NOT NULL,
  appeal_message TEXT NOT NULL,
  image1 TEXT,
  image2 TEXT,
  image3 TEXT,
  image4 TEXT,
  created_at datetime NOT NULL default current_timestamp,
  updated_at datetime NOT NULL default current_timestamp on update current_timestamp,
  deleted_at DATETIME,
  PRIMARY KEY(id)
  -- FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE mentors;
