const DeleteUserDTO = require('../dto/DeleteUserDTO');
const NotFoundException = require('../exception/NotFoundException');

const execute = async ({ _id }, userRepository, redisClient) => {
  const user = await userRepository.find({ _id });
  if (!user) throw new NotFoundException('Not Found');

  await userRepository.delete({ _id });

  // delete token from redis
  await redisClient.del(`/user/bearer/${_id}`);

  return new DeleteUserDTO();
}

module.exports = {
  execute,
}
