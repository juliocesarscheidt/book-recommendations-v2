const mongoose = require('mongoose');

const Schema = mongoose.Schema;

const UserSchema = new Schema({
  _id: { type: Schema.Types.ObjectId },
  name: { type: String },
  surname: { type: String },
  email: { type: String, unique: true },
  phone: { type: String },
  created_at: { type: Date },
  updated_at: { type: Date },
}, {
  timestamps: false,
  collection: 'user',
});

module.exports = UserSchema;
