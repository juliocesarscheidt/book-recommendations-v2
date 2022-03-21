const UpdateUserDTO = require('../dto/UpdateUserDTO');
const NotFoundException = require('../exception/NotFoundException');
const DuplicateEmailException = require('../exception/DuplicateEmailException');
const { encryptPassword } = require('../service/encryptionCommonService');

const execute = async ({ _id }, payload, userRepository, redisClient) => {
  const user = await userRepository.find({ _id });
  if (!user) throw new NotFoundException('Not Found');

  if (payload.password) {
    payload.password = encryptPassword(payload.password);
  }

  try {
    await userRepository.update({ _id }, payload);
    return new UpdateUserDTO();

  } catch (err) {
    if (err.code && err.code === 11000) { // MongoServerError: E11000 duplicate key error collection
      throw new DuplicateEmailException('Duplicate Email');
    }

    throw err;
  }
}

module.exports = {
  execute,
}
