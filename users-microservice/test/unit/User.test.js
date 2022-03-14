const { faker } = require('@faker-js/faker');

const UserRepository = require('../../infra/repository/memory/UserRepositoryMemory');
const { createUser, getUser, updateUser, deleteUser, listUser } = require('../../application/usecase/index');

let userRepository;

beforeAll(() => {
  userRepository = new UserRepository();
});

afterAll(() => {
});

beforeEach(() => {
});

afterEach(() => {
  userRepository.clean();
});

const createFakeUser = async (userRepository) => {
  const name = faker.name.firstName();
  const surname = faker.name.lastName();
  const email = faker.internet.email(name, surname).toLowerCase();
  const phone = faker.phone.phoneNumber('+55###########');
  const responseCreateUser = await createUser({ name, surname, email, phone }, userRepository);

  return responseCreateUser.uuid;
}

test('It should create a user and return its uuid', async function() {
  const uuid = await createFakeUser(userRepository);

  expect(uuid).not.toBeNull();
  expect(uuid).toHaveLength(24);
});

test('It should create a user and retrieve them by its _id', async function() {
  const _id = await createFakeUser(userRepository);
  const responseGetUser = await getUser({ _id }, userRepository);
  const { user } = responseGetUser;

  expect(user).not.toBeNull();
  expect(user.uuid).toEqual(_id);
  expect(user.name).not.toBeNull();
  expect(user.surname).not.toBeNull();
  expect(user.email).not.toBeNull();
  expect(user.phone).not.toBeNull();
});

test('It should create a user and they should have created_at and updated_at fields', async function() {
  const _id = await createFakeUser(userRepository);
  const responseGetUser = await getUser({ _id }, userRepository);
  const { user } = responseGetUser;

  expect(user).not.toBeNull();
  expect(user.created_at).not.toBeNull();
  expect(user.updated_at).not.toBeNull();
});

test('It should throw not found when retrieving an inexisting user', async function() {
  await expect(getUser({ _id: null }, userRepository)).rejects.toThrow('Not Found');
});

test('It should update the user', async function() {
  const uuid = await createFakeUser(userRepository);

  let responseGetUser = await getUser({ _id: uuid }, userRepository);
  const { user: userInitial } = responseGetUser;

  expect(userInitial).not.toBeNull();

  const newName = faker.name.firstName();
  await updateUser({ _id: uuid }, { name: newName }, userRepository);
  responseGetUser = await getUser({ _id: uuid }, userRepository);
  const { user: userFinal } = responseGetUser;

  // fields that should be different
  expect(userFinal.name).not.toBe(userInitial.name);
  expect(userFinal.updated_at).not.toBe(userInitial.updated_at);
   // field that should remain the same
  expect(userFinal.uuid).toEqual(userInitial.uuid);
  expect(userFinal.created_at).toEqual(userInitial.created_at);
  // fields that should not be null
  expect(userFinal.name).not.toBeNull();
  expect(userFinal.surname).not.toBeNull();
  expect(userFinal.email).not.toBeNull();
  expect(userFinal.phone).not.toBeNull();
});

test('It should throw not found when trying to update an inexisting user', async function() {
  await expect(updateUser({ _id: null }, { name: faker.name.firstName() }, userRepository)).rejects.toThrow('Not Found');
});

test('It should delete the user and the throw not found when we retrieve them', async function() {
  const _id = await createFakeUser(userRepository);
  await deleteUser({ _id }, userRepository);

  await expect(getUser({ _id }, userRepository)).rejects.toThrow('Not Found');
});

test('It should throw not found when trying to delete an inexisting user', async function() {
  await expect(deleteUser({ _id: null }, userRepository)).rejects.toThrow('Not Found');
});

test('It should list all users and have 0 users', async function() {
  const responseListUser = await listUser(0, 1000, userRepository);
  const { users } = responseListUser;

  expect(users).not.toBeNull();
  expect(users).toHaveLength(0);
});

test('It should list all users and have 1 users', async function() {
  await createFakeUser(userRepository);
  const responseListUser = await listUser(0, 1000, userRepository);
  const { users } = responseListUser;

  expect(users).not.toBeNull();
  expect(users).toHaveLength(1);
});
