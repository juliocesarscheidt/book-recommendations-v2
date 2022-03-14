# Users Microservice

```bash

yarn run lint

yarn run test
yarn run test:coverage

export MONGO_CONN_STRING=mongodb://root:admin@172.16.0.3:28017
# export AMQP_CONN_STRING=amqp://rabbitmq:admin@127.0.0.1:5872/

node index.js

```

## Data Structure

```javascript

mongo
database: user_db

user
[
  {
    "_id": "aafsfafs51as", # uuid
    "name": "",
    "surname": "",
    "email": "",
    "phone": "",
    "created_at": "",
    "updated_at": ""
  }
]

user_rate
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
