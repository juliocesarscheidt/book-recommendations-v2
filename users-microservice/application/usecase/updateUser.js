const UpdateUserDTO = require('../dto/UpdateUserDTO');
const NotFoundException = require('../exception/NotFoundException');
const { encryptPassword } = require('../service/EncryptionCommonService');

const execute = async ({ _id }, payload, userRepository, redisClient) => {
  const user = await userRepository.find({ _id });
  if (!user) throw new NotFoundException('Not Found');

  if (payload.password) {
    payload.password = encryptPassword(payload.password);
  }

  await userRepository.update({ _id }, payload);
  return new UpdateUserDTO();
}

module.exports = {
  execute,
}
