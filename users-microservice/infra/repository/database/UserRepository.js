var mongoose = require('mongoose');

const UserSchema = require('../../../domain/schema/UserSchema');
const MONGO_CONN_STRING = process.env.MONGO_CONN_STRING || 'mongodb://root:admin@127.0.0.1:27017';
const MONGO_DATABASE = process.env.MONGO_DATABASE || 'user_db';

class UserRepository {
  UserModel;

  constructor() {
    mongoose.connect(MONGO_CONN_STRING, { useNewUrlParser: true, dbName: MONGO_DATABASE });
    this.UserModel = mongoose.model('User', UserSchema);
  }

  async find(params) {
    const response = await this.UserModel.findOne(params);
    return response || null;
  }

  async update(params, payload) {
    await this.UserModel.updateOne(
      params,
      { $set: { ...payload, updated_at: new Date() } },
      { upsert: true },
    );
  }

  async delete(params) {
    await this.UserModel.deleteOne(params);
  }

  async insert({ _id, name, surname, email, phone, password }) {
    const user = new this.UserModel({ _id, name, surname, email, phone, password, created_at: new Date(), updated_at: new Date() });
    await user.save();
  }

  async list(page = 0, size = 1000) {
    const response = [];
    const cursor = this.UserModel.find({}).sort({ 'created_at': -1 }).skip(page * size).limit(size);
    for await (const row of cursor) {
      response.push(row);
    }
    return response;
  }

  async clean() {
  }
}

module.exports = UserRepository;
