db.createCollection('user');
db.createCollection('user_rate');

db.user.createIndex( { "email": 1 }, { unique: true } );
db.user.insertMany([
  {
    "name" : "julio",
    "surname" : "cesar",
    "email" : "admin@email.com",
    "phone" : "41995540808",
    "password" : "$2a$10$osLfeRqhrd5tea2xIXPZ4.Rid2QK6.hqfXP4f/64c4TGRNYQ5FLZ2",
    "created_at" : new Date(),
    "updated_at" : new Date(),
  }
]);
// db.user.getIndexes({})
