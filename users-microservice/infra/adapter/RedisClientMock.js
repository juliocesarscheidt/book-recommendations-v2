class RedisClientMock {
  valuesMap;

  constructor(connString) {
    this.valuesMap = new Map();
    this.getClient(connString);
  }

  async getClient(connString) {
  }

  async get(key) {
    const response = this.valuesMap.get(key);
    return response || null;
  }

  async set(key, value, expiryTime) {
    this.valuesMap.set(key, JSON.stringify(value));
  }

  async del(key) {
    this.valuesMap.delete(key);
  }
}

module.exports = RedisClientMock;
