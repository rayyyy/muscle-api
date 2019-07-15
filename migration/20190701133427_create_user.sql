-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
  id INT(11) AUTO_INCREMENT NOT NULL,
  uid VARCHAR(128) NOT NULL,
  email VARCHAR(255) NOT NULL,
  nickname VARCHAR(50) NOT NULL,
  image TEXT NOT NULL ,
  birthday DATE NOT NULL,
  sex ENUM('男性', '女性', '指定なし') DEFAULT '指定なし' NOT NULL,
  status ENUM('本客', '解約', '凍結') DEFAULT '本客' NOT NULL,
  created_at datetime NOT NULL default current_timestamp,
  updated_at datetime NOT NULL default current_timestamp on update current_timestamp,
  PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;