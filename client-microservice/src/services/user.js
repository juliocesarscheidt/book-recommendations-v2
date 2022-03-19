import http from '../utils/http.js';

const getCurrentUserInfo = async () => http
.get('/user/me')
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const createUser = async (name, surname, email, phone, password) => http
.post('/user', { name, surname, email, phone, password })
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const getUser = async (uuid) => http
.get(`/user/${uuid}`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const updateUser = async (uuid, { name, surname, email, phone }) => http
.put(`/user/${uuid}`, { name, surname, email, phone })
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const deleteUser = async (uuid) => http
.delete(`/user/${uuid}`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const listUser = async (page = 0, size = 50) => http
.get(`/user?page=${page}&size=${size}`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

export {
  getCurrentUserInfo,
  createUser,
  getUser,
  updateUser,
  deleteUser,
  listUser,
}
