const DeleteUserRateDTO = require('../dto/DeleteUserRateDTO');
const NotFoundException = require('../exception/NotFoundException');

const execute = async ({ _id }, userRateRepository) => {
  const userRate = await userRateRepository.find({ _id });
  if (!userRate) throw new NotFoundException('Not Found');

  await userRateRepository.delete({ _id });
  return new DeleteUserRateDTO();
}

module.exports = {
  execute,
}
