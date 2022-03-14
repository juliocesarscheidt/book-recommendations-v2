const { createUser, getUser, updateUser, deleteUser, listUser } = require('../../application/usecase/index');
const UserRepository = require('../repository/database/UserRepository');
// const UserRepository = require('../repository/memory/UserRepositoryMemory');

class UserService {
  userRepository;

  constructor() {
    this.userRepository = new UserRepository();
  }

  async CreateUser({ name, surname, email, phone }) {
    return createUser({ name, surname, email, phone }, this.userRepository);
  }

  async GetUser({ _id }) {
    return getUser({ _id }, this.userRepository);
  }

  async UpdateUser({ _id }, payload) {
    return updateUser({ _id }, payload, this.userRepository);
  }

  async DeleteUser({ _id }) {
    return deleteUser({ _id }, this.userRepository);
  }

  async ListUser(page, size) {
    return listUser(page, size, this.userRepository);
  }
}

module.exports = UserService;
