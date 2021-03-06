const fs = require('fs');
const path = require('path');
const bcrypt = require('bcryptjs');
const jwt = require('jsonwebtoken');

const encryptPassword = (password) => {
  const salt = bcrypt.genSaltSync();
  return bcrypt.hashSync(password, salt);
}

const comparePasswordSync = (passwordSource, passwordTarget) => {
  return bcrypt.compareSync(passwordSource, passwordTarget);
}

const decodeJwtToken = (token) => {
  try {
    /* eslint-disable */
    const publicKey = fs.readFileSync(path.resolve(__dirname, '../keys/jwtencyptionkey.pub'));
    /* eslint-enable */
    // Using RSA verifying method (RS256,RS384,RS512) expects privateKey for signing and publicKey for verification
    const tokenDecoded = jwt.verify(token, publicKey, {
      algorithms: ['RS256'],
      aud: 'urn:book-recommendations:api-gateway',
      iss: 'urn:book-recommendations:users-microservice'
    });
    if (new Date(tokenDecoded.exp * 1000) <= new Date()) {
      console.error('Token expired');
      return null;
    }
    return tokenDecoded;

  } catch (exception) {
    console.error('Exception', exception);
    return null;
  }
}

const encodeJwtToken = (payload) => {
  /* eslint-disable */
  const privateKey = fs.readFileSync(path.resolve(__dirname, '../keys/jwtencyptionkey.pem'));
  /* eslint-enable */
  // Using RSA signing method (RS256,RS384,RS512) expects privateKey for signing and publicKey for verification
  payload.aud = 'urn:book-recommendations:api-gateway';
  payload.iss = 'urn:book-recommendations:users-microservice';
  return jwt.sign(payload, privateKey, { algorithm: 'RS256'});
}

const generateUserToken = (user) => {
  const now = Math.floor(Date.now() / 1000);
  const exp = 60*60*1; // 1 hour
  // const exp = 1*24*60*60; // 1 day
  const payload = {
    ...user,
    iat: now,
    exp: now + exp
  };

  return encodeJwtToken(payload);
};

module.exports = {
  encryptPassword,
  comparePasswordSync,
  decodeJwtToken,
  encodeJwtToken,
  generateUserToken,
}
