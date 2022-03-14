const mongoose = require('mongoose');

const Schema = mongoose.Schema;

const UserRateSchema = new Schema({
  _id: { type: Schema.Types.ObjectId, ref: 'User' },
  rates: [
    { book_uuid: { type: String }, rate: { type: Number } },
  ],
  created_at: { type: Date },
  updated_at: { type: Date },
}, {
  timestamps: false,
  collection: 'user_rate',
});

module.exports = UserRateSchema;
