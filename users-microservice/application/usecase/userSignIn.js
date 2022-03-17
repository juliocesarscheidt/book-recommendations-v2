const SignInDTO = require('../dto/SignInDTO');
const NotFoundException = require('../exception/NotFoundException');
const InvalidEmailPasswordException = require('../exception/InvalidEmailPasswordException');
const { comparePasswordSync, generateUserToken } = require('../service/EncryptionCommonService');

const execute = async ({ email, password }, userRepository) => {
  const user = await userRepository.find({ email });
  if (!user) throw new NotFoundException('Not Found');

  if (!comparePasswordSync(password, user.password)) {
    throw new InvalidEmailPasswordException('Invalid email or password');
  }

  const userBody = {
    uuid: user._id,
    name: user.name,
    surname: user.surname,
    email: user.email,
    phone: user.phone,
  };
  const accessToken = generateUserToken(userBody);

  return new SignInDTO(accessToken);
}

module.exports = {
  execute,
}
