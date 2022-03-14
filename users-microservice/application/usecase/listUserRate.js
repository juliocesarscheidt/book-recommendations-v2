const ListUserRateDTO = require('../dto/ListUserRateDTO');

const execute = async (page, size, userRateRepository) => {
  const userRates = await userRateRepository.list(page, size);
  return new ListUserRateDTO(userRates);
}

module.exports = {
  execute,
}
