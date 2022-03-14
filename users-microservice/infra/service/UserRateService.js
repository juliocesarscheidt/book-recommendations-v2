const { upsertUserRate, getUserRate, deleteUserRate, listUserRate } = require('../../application/usecase/index');
const UserRateRepository = require('../repository/database/UserRateRepository');
// const UserRateRepository = require('../repository/memory/UserRateRepositoryMemory');

class UserRateService {
  userRateRepository;

  constructor() {
    this.userRateRepository = new UserRateRepository();
  }

  async UpsertUserRate({ _id, book_uuid, rate }) {
    return upsertUserRate({ _id, book_uuid, rate }, this.userRateRepository);
  }

  async GetUserRate({ _id }) {
    return getUserRate({ _id }, this.userRateRepository);
  }

  async DeleteUserRate({ _id }) {
    return deleteUserRate({ _id }, this.userRateRepository);
  }

  async ListUserRate(page, size) {
    return listUserRate(page, size, this.userRateRepository);
  }
}

module.exports = UserRateService;
