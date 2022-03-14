const GetUserRateDTO = require('../dto/GetUserRateDTO');
const NotFoundException = require('../exception/NotFoundException');

const execute = async ({ _id }, userRateRepository) => {
  const userRate = await userRateRepository.find({ _id });
  if (!userRate) throw new NotFoundException('Not Found');

  return new GetUserRateDTO(userRate);
}

module.exports = {
  execute,
}
