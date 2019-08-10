CREATE DATABASE IF NOT EXISTS app1;
CREATE TABLE IF NOT EXISTS app1.todo(
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100),
  detail VARCHAR(1000),
  auther VARCHAR(1000),
  created_at TIMESTAMP NOT NULL default current_timestamp,
  updated_at TIMESTAMP NOT NULL default current_timestamp on update current_timestamp
);