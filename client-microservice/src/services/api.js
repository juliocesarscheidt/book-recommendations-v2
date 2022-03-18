import axios from 'axios'

const signUp = async (email, password) => {
  return axios
    .post('/v1/auth/signup', { email, password })
    .then((response) => {
      console.log('API signup response', response);
      if (!response.data) {
        return null
      }
      const access_token = response.data.data
      return access_token;
    });
}

const signIn = async (email, password) => {
  return axios
    .post('/v1/auth/signin', { email, password })
    .then((response) => {
      console.log('API signin response', response);
      if (!response.data) {
        return null
      }
      const access_token = response.data.data
      return access_token;
    });
}

export {
  signUp,
  signIn,
}

