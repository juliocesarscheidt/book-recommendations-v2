# Data Structure

## MongoDB

```
database: user_db

table: user
[
  {
    "_id": "aafsfafs51as", # uuid
    "name": "",
    "surname": "",
    "email": "",
    "phone": "",
    "password": "",
    "created_at": "",
    "updated_at": ""
  }
]

table: user_rate
[
  {
    "_id": "aafsfafs51as", # user_uuid -> references user _id
    "rates": [
      {"book_uuid": "af45afasf", "rate": 4},
      {"book_uuid": "3421412", "rate": 6}
    ],
    "created_at": "",
    "updated_at": ""
  }
]
```

## Postgres

```
database: book_db

table: book
  uuid VARCHAR(24) PRIMARY KEY NOT NULL
  title VARCHAR(255) NOT NULL
  author VARCHAR(255) NOT NULL
  genre VARCHAR(255) NOT NULL
  image VARCHAR(255) NOT NULL
  created_at timestamp without time zone default (now() at time zone 'utc')
  updated_at timestamp without time zone
```
