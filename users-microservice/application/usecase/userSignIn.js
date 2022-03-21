const SignInDTO = require('../dto/SignInDTO');
const InvalidEmailPasswordException = require('../exception/InvalidEmailPasswordException');
const { comparePasswordSync, generateUserToken } = require('../service/encryptionCommonService');

const execute = async ({ email, password }, userRepository, redisClient) => {
  const user = await userRepository.find({ email });
  if (!user) throw new InvalidEmailPasswordException('Invalid Email or Password');

  if (!comparePasswordSync(password, user.password)) {
    throw new InvalidEmailPasswordException('Invalid Email or Password');
  }

  const userBody = {
    uuid: user._id,
    name: user.name,
    surname: user.surname,
    email: user.email,
    phone: user.phone,
  };
  const accessToken = generateUserToken(userBody);

  // add token to redis for further validations
  await redisClient.set(`/user/bearer/${user._id}`, accessToken, 60*60*1);

  return new SignInDTO(accessToken);
}

module.exports = {
  execute,
}
