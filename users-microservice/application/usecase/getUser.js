const GetUserDTO = require('../dto/GetUserDTO');
const NotFoundException = require('../exception/NotFoundException');

const execute = async ({ _id }, userRepository) => {
  const user = await userRepository.find({ _id });
  if (!user) throw new NotFoundException('Not Found');

  return new GetUserDTO(user);
}

module.exports = {
  execute,
}
