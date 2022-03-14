const UpsertUserRateDTO = require('../dto/UpsertUserRateDTO');
const InvalidRateException = require('../exception/InvalidRateException');

const execute = async ({ _id, book_uuid, rate }, userRateRepository) => {
  if (rate < 0 || rate > 10) {
    throw new InvalidRateException('Invalid Rate');
  }

  const existingUserRate = await userRateRepository.find({ _id });
  if (!existingUserRate) {
    const createdUserRate = { _id, rates: [{ book_uuid, rate }] };
    await userRateRepository.insert(createdUserRate);

  } else {
    const existingRateIndex = existingUserRate.rates.findIndex(r => r.book_uuid === book_uuid);
    const updatedUserRate = { rates: existingUserRate.rates };
    if (existingRateIndex >= 0) {
      // no need to update, return
      if (updatedUserRate.rates[existingRateIndex].rate === rate) return new UpsertUserRateDTO(_id);

      updatedUserRate.rates[existingRateIndex].rate = rate;
    } else {
      updatedUserRate.rates.push({ book_uuid, rate });
    }

    await userRateRepository.update({ _id }, updatedUserRate);
  }

  return new UpsertUserRateDTO(_id);
}

module.exports = {
  execute,
}
