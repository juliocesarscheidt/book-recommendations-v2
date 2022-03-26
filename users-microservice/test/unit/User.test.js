const { faker } = require('@faker-js/faker');

const UserRepository = require('../../infra/repository/memory/UserRepositoryMemory');
const RedisClient = require('../../infra/adapter/RedisClientMock');
const { createUser, getUser, updateUser, deleteUser, listUser, userSignUp, userSignIn } = require('../../application/usecase/index');
const { decodeJwtToken } = require('../../application/service/encryptionCommonService');

let userRepository;
let redisClient;

beforeAll(() => {
  userRepository = new UserRepository();
  redisClient = new RedisClient();
});

afterAll(() => {
});

beforeEach(() => {
});

afterEach(() => {
  userRepository.clean();
});

const buildFakeUser = () => {
  const name = faker.name.firstName();
  const surname = faker.name.lastName();
  const email = faker.internet.email(name, surname).toLowerCase();
  const phone = faker.phone.phoneNumber('###########');
  const password = 'PASSWORD';

  return { name, surname, email, phone, password };
}

const createFakeUser = async (userRepository, redisClient) => {
  const userBody = buildFakeUser();
  const responseCreateUser = await createUser(userBody, userRepository, redisClient);

  return responseCreateUser.uuid;
}

test('It should create a user and return its uuid', async function() {
  const uuid = await createFakeUser(userRepository, redisClient);

  expect(uuid).not.toBeNull();
  expect(uuid).toHaveLength(24);
});

test('It should create a user through signup and return its access_token', async function() {
  const userBody = buildFakeUser();
  const SignUpResponse = await userSignUp(userBody, userRepository, redisClient);
  const { access_token } = SignUpResponse;
  expect(access_token).not.toBeNull();

  const tokenDecoded = decodeJwtToken(access_token);
  expect(tokenDecoded).not.toBeNull();
  expect(tokenDecoded.uuid).not.toBeNull();

  const redisValue = await redisClient.get(`/user/bearer/${tokenDecoded.uuid}`);
  expect(redisValue).not.toBeNull();
});

test('It should create a user, do signin and return its access_token', async function() {
  const uuid = await createFakeUser(userRepository, redisClient);
  const responseGetUser = await getUser({ _id: uuid }, userRepository, redisClient);
  const { user } = responseGetUser;

  const SignInResponse = await userSignIn({ email: user.email, password: 'PASSWORD' }, userRepository, redisClient);
  const { access_token } = SignInResponse;
  expect(access_token).not.toBeNull();

  const tokenDecoded = decodeJwtToken(access_token);
  expect(tokenDecoded).not.toBeNull();
  expect(tokenDecoded.uuid).not.toBeNull();

  const redisValue = await redisClient.get(`/user/bearer/${tokenDecoded.uuid}`);
  expect(redisValue).not.toBeNull();
});

test('It should throw Invalid Email or Password Exception when signin with an inexisting user', async function() {
  await expect(userSignIn({ email: 'admin@email.com', password: 'PASSWORD' }, userRepository, redisClient)).rejects.toThrow('Invalid Email or Password');
});

test('It should create a user and retrieve them by its _id', async function() {
  const _id = await createFakeUser(userRepository, redisClient);
  const responseGetUser = await getUser({ _id }, userRepository, redisClient);
  const { user } = responseGetUser;

  expect(user).not.toBeNull();
  expect(user.uuid).toEqual(_id);
  expect(user.name).not.toBeNull();
  expect(user.surname).not.toBeNull();
  expect(user.email).not.toBeNull();
  expect(user.phone).not.toBeNull();
});

test('It should create a user and they should have created_at and updated_at fields', async function() {
  const _id = await createFakeUser(userRepository, redisClient);
  const responseGetUser = await getUser({ _id }, userRepository, redisClient);
  const { user } = responseGetUser;

  expect(user).not.toBeNull();
  expect(user.created_at).not.toBeNull();
  expect(user.updated_at).not.toBeNull();
});

test('It should throw not found when retrieving an inexisting user', async function() {
  await expect(getUser({ _id: null }, userRepository, redisClient)).rejects.toThrow('Not Found');
});

test('It should update the user', async function() {
  const uuid = await createFakeUser(userRepository, redisClient);

  let responseGetUser = await getUser({ _id: uuid }, userRepository, redisClient);
  const { user: userInitial } = responseGetUser;

  expect(userInitial).not.toBeNull();

  const newName = faker.name.firstName();
  await updateUser({ _id: uuid }, { name: newName }, userRepository, redisClient);
  responseGetUser = await getUser({ _id: uuid }, userRepository, redisClient);
  const { user: userFinal } = responseGetUser;

  // fields that should be different
  expect(userFinal.name).not.toBe(userInitial.name);
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
  const _id = await createFakeUser(userRepository, redisClient);
  await deleteUser({ _id }, userRepository, redisClient);

  const redisValue = await redisClient.get(`/user/bearer/${_id}`);
  expect(redisValue).toBeNull();

  await expect(getUser({ _id }, userRepository, redisClient)).rejects.toThrow('Not Found');
});

test('It should throw not found when trying to delete an inexisting user', async function() {
  await expect(deleteUser({ _id: null }, userRepository, redisClient)).rejects.toThrow('Not Found');
});

test('It should list all users and have 0 users', async function() {
  const responseListUser = await listUser(0, 1000, userRepository, redisClient);
  const { users } = responseListUser;

  expect(users).not.toBeNull();
  expect(users).toHaveLength(0);
});

test('It should list all users and have 1 users', async function() {
  await createFakeUser(userRepository);
  const responseListUser = await listUser(0, 1000, userRepository, redisClient);
  const { users } = responseListUser;

  expect(users).not.toBeNull();
  expect(users).toHaveLength(1);
});
