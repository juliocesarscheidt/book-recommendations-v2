const UserRateRepository = require('../../infra/repository/memory/UserRateRepositoryMemory');
const RedisClient = require('../../infra/adapter/RedisClientMock');
const { upsertUserRate, getUserRate, deleteUserRate, listUserRate } = require('../../application/usecase/index');
const { generateRandomString } = require('../../common/stringUtils');

let userRateRepository;
let redisClient;

beforeAll(() => {
  userRateRepository = new UserRateRepository();
  redisClient = new RedisClient();
});

afterAll(() => {
});

beforeEach(() => {
});

afterEach(() => {
  userRateRepository.clean();
});

const createFakeUserRate = async (userRateRepository, rate = 8.0) => {
  const userRate = {
    _id: generateRandomString(),
    book_uuid: generateRandomString(),
    rate,
  };
  const responseCreateUserRate = await upsertUserRate(userRate, userRateRepository);

  return responseCreateUserRate.user_uuid;
}

test('It should create a user rate and return its uuid', async function() {
  const user_uuid = await createFakeUserRate(userRateRepository);

  expect(user_uuid).not.toBeNull();
  expect(user_uuid).toHaveLength(24);
});

test('It should throw not found when creating an invalid user rate', async function() {
  await expect(createFakeUserRate(userRateRepository, -2.0)).rejects.toThrow('Invalid Rate');
});

test('It should create a user rate and retrieve them by its _id', async function() {
  const _id = await createFakeUserRate(userRateRepository);
  const responseGetUserRate = await getUserRate({ _id }, userRateRepository);
  const { userRate } = responseGetUserRate;

  expect(userRate).not.toBeNull();
  expect(userRate.user_uuid).toEqual(_id);
  expect(userRate.name).not.toBeNull();
  expect(userRate.surname).not.toBeNull();
  expect(userRate.email).not.toBeNull();
  expect(userRate.phone).not.toBeNull();
});

test('It should create a user rate and they should have created_at and updated_at fields', async function() {
  const _id = await createFakeUserRate(userRateRepository);
  const responseGetUserRate = await getUserRate({ _id }, userRateRepository);
  const { userRate } = responseGetUserRate;

  expect(userRate).not.toBeNull();
  expect(userRate.created_at).not.toBeNull();
  expect(userRate.updated_at).not.toBeNull();
});

test('It should return null when retrieving an inexisting user rate', async function() {
  const responseGetUserRate = await getUserRate({ _id: null }, userRateRepository);
  const { userRate } = responseGetUserRate;

  expect(userRate).toBeNull();
});

test('It should update the user rate', async function() {
  const user_uuid = await createFakeUserRate(userRateRepository);
  let responseGetUserRate = await getUserRate({ _id: user_uuid }, userRateRepository);
  const { userRate: userRateInitial } = responseGetUserRate;

  expect(userRateInitial).not.toBeNull();

  const userRateCopy = JSON.parse(JSON.stringify({
    _id: userRateInitial.user_uuid,
    book_uuid: userRateInitial.rates[0].book_uuid,
    rate: 5,
  }));
  await upsertUserRate(userRateCopy, userRateRepository);

  responseGetUserRate = await getUserRate({ _id: user_uuid }, userRateRepository);
  const { userRate: userRateFinal } = responseGetUserRate;

  expect(userRateFinal.rates[0].rate).toBe(5);

  // fields that should be different
  expect(userRateFinal.rates).not.toBe(userRateInitial.rates);
   // field that should remain the same
  expect(userRateFinal.user_uuid).toEqual(userRateInitial.user_uuid);
  expect(userRateFinal.created_at).toEqual(userRateInitial.created_at);
});

test('It should delete the user rate and return null when retrieving them', async function() {
  const _id = await createFakeUserRate(userRateRepository);
  await deleteUserRate({ _id }, userRateRepository);

  const responseGetUserRate = await getUserRate({ _id }, userRateRepository);
  const { userRate } = responseGetUserRate;

  await expect(userRate).toBeNull();
});

test('It should list all user rates and have 0 user rates', async function() {
  const responseListUserRate = await listUserRate(0, 1000, userRateRepository);
  const { userRates } = responseListUserRate;

  expect(userRates).not.toBeNull();
  expect(userRates).toHaveLength(0);
});

test('It should list all user rates and have 1 user rates', async function() {
  await createFakeUserRate(userRateRepository);
  const responseListUserRate = await listUserRate(0, 1000, userRateRepository);
  const { userRates } = responseListUserRate;

  expect(userRates).not.toBeNull();
  expect(userRates).toHaveLength(1);
});
