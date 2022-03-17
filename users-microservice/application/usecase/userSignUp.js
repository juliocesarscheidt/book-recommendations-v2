const SignUpDTO = require('../dto/SignUpDTO');
const { encryptPassword, generateUserToken } = require('../service/EncryptionCommonService');

const execute = async ({ name, surname, email, phone, password }, userRepository) => {
  const uuid = ((new Date()).getTime().toString(16) + Math.random().toString(16)).replace('.', '').substring(0, 24);
  if (password) {
    password = encryptPassword(password);
  }
  const user = { _id: uuid, name, surname, email, phone, password };

  await userRepository.insert(user);

  const userBody = { uuid, name, surname, email, phone };
  const accessToken = generateUserToken(userBody);

  return new SignUpDTO(accessToken);
}

module.exports = {
  execute,
}
