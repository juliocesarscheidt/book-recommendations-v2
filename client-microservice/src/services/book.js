import http from '../utils/http.js';

const uploadEventCallable = (uploadEvent) => {
  console.log(`Upload progress: ${(uploadEvent.loaded / uploadEvent.total) * 100} %`);
}

const buildFormData = (title, author, genre, image) => {
  const fd = new FormData();
  fd.append('title', title);
  fd.append('author', author);
  fd.append('genre', genre);
  fd.append('image', image, image.name);
  return fd;
}

const createBook = async (title, author, genre, image) => {
  const fd = buildFormData(title, author, genre, image);
  return http.post('/book', fd, { onUploadProgress: uploadEventCallable })
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  })
}

const getBook = async (uuid) => http
.get(`/book/${uuid}`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const getBookPresignUrl = async (uuid) => http
.get(`/book/${uuid}/image/url`)
.then((response) => {
  if (!response.data) {
    return null
  }
  return response.data.data
});

const updateBookWithImage = async (uuid, { title, author, genre, image }) => {
  const fd = buildFormData(title, author, genre, image);
  return http.put(`/book/${uuid}/image`, fd, { onUploadProgress: uploadEventCallable })
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  })
}

const updateBook = async (uuid, { title, author, genre }) => http
.put(`/book/${uuid}`, { title, author, genre })
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
  getBookPresignUrl,
  updateBookWithImage,
  updateBook,
  deleteBook,
  listBook,
}
