use devops;
DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  login     VARCHAR(255) NOT NULL,
  org       VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users
  (name, login, org)
VALUES
  ('Jhon Doe', 'jdoe@mail.com', 'belhard'),
  ('James Bond', 'jbond@mail.com', 'mi6'),
  ('Ivan Ivanov', 'ii@mail.com', 'org');
