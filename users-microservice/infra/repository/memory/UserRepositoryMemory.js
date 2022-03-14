class UserRepositoryMemory {
  users;

  constructor() {
    this.users = [];
  }

  async find(params) {
    const response = this.users.find(u => u?._id?.toString() === params?._id?.toString());
    return response || null;
  }

  async update(params, payload) {
    const index = this.users.findIndex(u => u?._id?.toString() === params?._id?.toString());
    const user = this.users[index];
    this.users.splice(index, 1, { ...user, ...payload, updated_at: new Date() });
  }

  async delete(params) {
    const index = this.users.findIndex(u => u?._id?.toString() === params?._id?.toString());
    this.users.splice(index, 1);
  }

  async insert(payload) {
    this.users.push({ ...payload, created_at: new Date(), updated_at: new Date() });
  }

  async list(page = 0, size = 1000) {
    const offset = page * size;
    const users = this.users.slice(offset, ((page + 1) * size));
    return users || [];
  }

  async clean() {
    this.users = [];
  }
}

module.exports = UserRepositoryMemory;
