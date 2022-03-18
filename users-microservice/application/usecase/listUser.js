const ListUserDTO = require('../dto/ListUserDTO');

const execute = async (page, size, userRepository, redisClient) => {
  const users = await userRepository.list(page, size);
  return new ListUserDTO(users);
}

module.exports = {
  execute,
}
