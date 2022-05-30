const CreateUserDTO = require('../dto/CreateUserDTO');
const DuplicateEmailException = require('../exception/DuplicateEmailException');
const { encryptPassword } = require('../../common/encryptionUtils');
const { generateRandomString } = require('../../common/stringUtils');

const execute = async ({ name, surname, email, phone, password }, userRepository, redisClient) => {
  let uuid = generateRandomString(24);
  if (password) {
    password = encryptPassword(password);
  }
  const user = { _id: uuid, name, surname, email, phone, password };

  try {
    await userRepository.insert(user);
    return new CreateUserDTO(uuid);

  } catch (err) {
    console.error(err);
    if (err.code && err.code === 11000) { // MongoServerError: E11000 duplicate key error collection
      throw new DuplicateEmailException('Duplicate Email');
    }
    throw err;
  }
}

module.exports = {
  execute,
}
