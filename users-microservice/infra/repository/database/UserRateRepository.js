var mongoose = require('mongoose');

const UserRateSchema = require('../../../domain/schema/UserRateSchema');
const MONGO_CONN_STRING = process.env.MONGO_CONN_STRING || 'mongodb://root:admin@127.0.0.1:27017';
const MONGO_DATABASE = process.env.MONGO_DATABASE || 'user_db';

class UserRateRepository {
  UserRateModel;

  constructor() {
    mongoose.connect(MONGO_CONN_STRING, { useNewUrlParser: true, dbName: MONGO_DATABASE });
    this.UserRateModel = mongoose.model('UserRate', UserRateSchema);
  }

  async find(params) {
    const response = await this.UserRateModel.findOne(params);
    return response || null;
  }

  async update(params, payload) {
    await this.UserRateModel.updateOne(
      params,
      { $set: { ...payload, updated_at: new Date() } },
      { upsert: true },
    );
  }

  async delete(params) {
    await this.UserRateModel.deleteOne(params);
  }

  async insert({ _id, rates }) {
    const user = new this.UserRateModel({ _id, rates, created_at: new Date(), updated_at: new Date() });
    await user.save();
  }

  async list(page = 0, size = 1000) {
    const response = [];
    const cursor = this.UserRateModel.find({}).sort({ 'created_at': -1 }).skip(page * size).limit(size);
    for await (const row of cursor) {
      response.push(row);
    }
    return response;
  }

  async clean() {
  }
}

module.exports = UserRateRepository;
