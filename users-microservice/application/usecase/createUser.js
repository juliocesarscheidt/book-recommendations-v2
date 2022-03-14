const CreateUserDTO = require('../dto/CreateUserDTO');

const execute = async ({ name, surname, email, phone }, userRepository) => {
  const uuid = ((new Date()).getTime().toString(16) + Math.random().toString(16)).replace('.', '').substring(0, 24);
  const user = { _id: uuid, name, surname, email, phone };

  await userRepository.insert(user);
  return new CreateUserDTO(uuid);
}

module.exports = {
  execute,
}
