const DeleteUserRateDTO = require('../dto/DeleteUserRateDTO');

const execute = async ({ _id }, userRateRepository) => {
  await userRateRepository.delete({ _id });
  return new DeleteUserRateDTO();
}

module.exports = {
  execute,
}
