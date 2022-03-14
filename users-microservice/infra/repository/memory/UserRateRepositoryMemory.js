class UserRateRepositoryMemory {
  userRates;

  constructor() {
    this.userRates = [];
  }

  async find(params) {
    const response = this.userRates.find(u => u?._id?.toString() === params?._id?.toString());
    return response || null;
  }

  async update(params, payload) {
    const index = this.userRates.findIndex(u => u?._id?.toString() === params?._id?.toString());
    const userRate = this.userRates[index];
    this.userRates.splice(index, 1, { ...userRate, ...payload, updated_at: new Date() });
  }

  async delete(params) {
    const index = this.userRates.findIndex(u => u?._id?.toString() === params?._id?.toString());
    this.userRates.splice(index, 1);
  }

  async insert(payload) {
    this.userRates.push({ ...payload, created_at: new Date(), updated_at: new Date() });
  }

  async list(page = 0, size = 1000) {
    const offset = page * size;
    const userRates = this.userRates.slice(offset, ((page + 1) * size));
    return userRates || [];
  }

  async clean() {
    this.userRates = [];
  }
}

module.exports = UserRateRepositoryMemory;
