import http from '../utils/http.js';

const upsertRate = async (user_uuid, book_uuid, rate) => http
  .post('/user/rating', { user_uuid, book_uuid, rate })
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  });

const getRate = async (uuid) => http
  .get(`/user/rating/${uuid}`)
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
});

const deleteRate = async (uuid) => http
  .delete(`/user/rating/${uuid}`)
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  });

const listRate = async (page = 0, size = 50) => http
  .get(`/user/rating?page=${page}&size=${size}`)
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  });

export {
  upsertRate,
  getRate,
  deleteRate,
  listRate,
}
