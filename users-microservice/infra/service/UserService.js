const { userSignUp, userSignIn, createUser, getUser, updateUser, deleteUser, listUser } = require('../../application/usecase/index');
const UserRepository = require('../repository/database/UserRepository');
const RedisClient = require('../adapter/RedisClient');

/* eslint-disable */
const REDIS_CONN_STRING = process.env.REDIS_CONN_STRING || '127.0.0.1:5672';
/* eslint-enable */

class UserService {
  userRepository;
  redisClient;

  constructor() {
    this.userRepository = new UserRepository();
    this.redisClient = new RedisClient(REDIS_CONN_STRING);
  }

  async UserSignUp({ name, surname, email, phone, password }) {
    return userSignUp({ name, surname, email, phone, password }, this.userRepository, this.redisClient);
  }

  async UserSignIn({ email, password }) {
    return userSignIn({ email, password }, this.userRepository, this.redisClient);
  }

  async CreateUser({ name, surname, email, phone, password }) {
    return createUser({ name, surname, email, phone, password }, this.userRepository, this.redisClient);
  }

  async GetUser({ _id }) {
    return getUser({ _id }, this.userRepository, this.redisClient);
  }

  async UpdateUser({ _id }, payload) {
    return updateUser({ _id }, payload, this.userRepository, this.redisClient);
  }

  async DeleteUser({ _id }) {
    return deleteUser({ _id }, this.userRepository, this.redisClient);
  }

  async ListUser(page, size) {
    return listUser(page, size, this.userRepository, this.redisClient);
  }
}

module.exports = UserService;
