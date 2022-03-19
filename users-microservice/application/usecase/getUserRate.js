const GetUserRateDTO = require('../dto/GetUserRateDTO');

const execute = async ({ _id }, userRateRepository) => {
  const userRate = await userRateRepository.find({ _id });
  return new GetUserRateDTO(userRate);
}

module.exports = {
  execute,
}
