const ListUserDTO = require('../dto/ListUserDTO');

const execute = async (page, size, userRepository) => {
  const users = await userRepository.list(page, size);
  return new ListUserDTO(users);
}

module.exports = {
  execute,
}
