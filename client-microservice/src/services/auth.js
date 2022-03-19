import http from '../utils/http';

const signUp = async (name, surname, email, phone, password) => http
  .post('/auth/signup', { name, surname, email, phone, password })
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  });

const signIn = async (email, password) => http
  .post('/auth/signin', { email, password })
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  });

export {
  signUp,
  signIn,
}
