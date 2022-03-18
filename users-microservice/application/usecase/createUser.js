const CreateUserDTO = require('../dto/CreateUserDTO');
const { encryptPassword } = require('../service/EncryptionCommonService');

const execute = async ({ name, surname, email, phone, password }, userRepository, redisClient) => {
  const uuid = ((new Date()).getTime().toString(16) + Math.random().toString(16)).replace('.', '').substring(0, 24);
  if (password) {
    password = encryptPassword(password);
  }
  const user = { _id: uuid, name, surname, email, phone, password };

  await userRepository.insert(user);
  return new CreateUserDTO(uuid);
}

module.exports = {
  execute,
}
