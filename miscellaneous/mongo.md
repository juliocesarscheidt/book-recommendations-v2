# Mongo Commands

```bash

docker-compose -f docker-compose-infra.yaml up -d mongo
docker-compose -f docker-compose-infra.yaml logs -f --tail 50 mongo

docker-compose -f docker-compose-infra.yaml exec mongo bash


docker-compose -f docker-compose-infra.yaml exec mongo bash -c \
  "echo 'db.runCommand(\"ping\").ok' | mongosh 'mongodb://root:admin@127.0.0.1:27017/admin' --quiet"

docker-compose -f docker-compose-infra.yaml exec mongo bash -c "mongosh 'mongodb://root:admin@127.0.0.1:27017/admin'"


mongo "mongodb://root:admin@127.0.0.1:28017/admin?retryWrites=true&w=majority"

mongo --host 127.0.0.1 --port 28017 -- user_db

use admin;
db.auth("root", "admin");

show dbs;

db.runCommand({connectionStatus: 1});
show roles;


use user_db;
show collections;

db.getCollectionNames();

db.user.find({});
db.user.find({"_id": ObjectId("17f9ad3783d07a4f3e3db8ce")}).pretty();


db.user.createIndex({ email: -1 })

db.user.getIndexes()

db.user.dropIndex({ email: -1 })

```
