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
    const publicKey = fs.readFileSync(path.resolve(__dirname, '../../keys/jwtencyptionkey.pub'));
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
  const privateKey = fs.readFileSync(path.resolve(__dirname, '../../keys/jwtencyptionkey.pem'));
  // Using RSA signing method (RS256,RS384,RS512) expects privateKey for signing and publicKey for validation
  return jwt.sign(payload, privateKey, { algorithm: 'RS256'});
}

const generateUserToken = (user) => {
  const now = Math.floor(Date.now() / 1000);
  const exp = 1 * 60 * 60; // 1 hour
  // const exp = 1 * 24 * 60 * 60; // 1 day
  const payload = {
    ...user,
    iat: now,
    exp: now + exp,
    aud: 'urn:book-recommendations:api-gateway',
    iss: 'urn:book-recommendations:users-microservice'
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
