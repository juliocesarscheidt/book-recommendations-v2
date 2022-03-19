import http from '../utils/http.js';

const createBook = async (title, author, genre, image) => http
.post('/book', { title, author, genre, image })
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const getBook = async (uuid) => http
.get(`/book/${uuid}`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const updateBook = async (uuid, { title, author, genre, image }) => http
.put(`/book/${uuid}`, { title, author, genre, image })
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const deleteBook = async (uuid) => http
.delete(`/book/${uuid}`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const listBook = async (page = 0, size = 50) => http
.get(`/book?page=${page}&size=${size}`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

export {
  createBook,
  getBook,
  updateBook,
  deleteBook,
  listBook,
}
