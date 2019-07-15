-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE user_details (
  id INT(11) AUTO_INCREMENT NOT NULL,
  user_id INT(11) NOT NULL,
  short_message VARCHAR(255),
  height DECIMAL ( 5, 2 ),
  weight DECIMAL ( 5, 2 ),
  blood_type ENUM('A', 'B', 'AB', 'O', '指定なし') DEFAULT '指定なし',
  get_up_time time,
  bedtime time,
  drinking_habits ENUM('飲まない', '週1程度', '週2~3', '週4~5', '週6~7', '指定なし') DEFAULT '指定なし',
  smoking_habit ENUM('吸わない', '週1程度', '週2~3', '週4~5', '週6~7', '指定なし') DEFAULT '指定なし',
  created_at datetime NOT NULL default current_timestamp,
  updated_at datetime NOT NULL default current_timestamp on update current_timestamp,
  PRIMARY KEY(id)
  -- FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE user_details;
