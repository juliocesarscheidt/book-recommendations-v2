const { createClient } = require('redis');

class RedisClient {
  client;

  constructor(connString) {
    this.getClient(connString);
  }

  async getClient(connString) {
    this.client = createClient({ url: `redis://${connString}` });
    this.client.on('error', (err) => console.log('Failed to connect to Redis endpoint', err));
    await this.client.connect();
  }

  async get(key) {
    const response = await this.client.get(key);
    return response || null;
  }

  async set(key, value, expiryTime) {
    await this.client.set(key, JSON.stringify(value), { 'EX': expiryTime });
  }

  async del(key) {
    await this.client.del(key);
  }
}

module.exports = RedisClient;
