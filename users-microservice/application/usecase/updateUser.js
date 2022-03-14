const UpdateUserDTO = require('../dto/UpdateUserDTO');
const NotFoundException = require('../exception/NotFoundException');

const execute = async ({ _id }, payload, userRepository) => {
  const user = await userRepository.find({ _id });
  if (!user) throw new NotFoundException('Not Found');

  await userRepository.update({ _id }, payload);
  return new UpdateUserDTO();
}

module.exports = {
  execute,
}
