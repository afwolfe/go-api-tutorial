DROP DATABASE IF EXISTS booksdb;
CREATE DATABASE booksdb;
USE booksdb;
DROP TABLE IF EXISTS books;
GRANT ALL PRIVILEGES ON books to 'appuser'@'%' WITH GRANT OPTION;

CREATE TABLE books (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  author     VARCHAR(255) NOT NULL,
  quantity   INT NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO books
  (title, author, quantity)
VALUES
  ('In Search of Lost Time', 'Marcel Proust', 2),
  ('The Great Gatsby', 'F. Scott Fitzgerald', 5),
  ('War and Peace', 'Leo Tolstoy', 6);
