const generateRandomHexString = (size = 24) => {
  let result = '';
  const characters = 'abcdef0123456789';
  for (let i = 0; i < size; i++) {
    result += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  return result;
}

module.exports = {
  generateRandomHexString,
}
