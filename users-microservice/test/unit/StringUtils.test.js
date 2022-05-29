const { faker } = require('@faker-js/faker');

const { generateRandomString } = require('../../common/stringUtils');

beforeAll(() => {
});

afterAll(() => {
});

beforeEach(() => {
});

afterEach(() => {
});

test('It should generate a random string with 24 characters', async function() {
  const str = generateRandomString(24);

  expect(str).not.toBeNull();
  expect(str).toHaveLength(24);
});

test('It should generate a random string with 36 characters', async function() {
  const str = generateRandomString(36);

  expect(str).not.toBeNull();
  expect(str).toHaveLength(36);
});
