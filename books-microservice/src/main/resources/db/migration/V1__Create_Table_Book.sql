CREATE TABLE "book" (
  uuid VARCHAR(24) PRIMARY KEY NOT NULL,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  genre VARCHAR(255) NOT NULL,
  image VARCHAR(255) NOT NULL,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  updated_at timestamp without time zone
);
