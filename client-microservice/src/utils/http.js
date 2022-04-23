import axios from 'axios';

const http = axios.create({
  baseURL: '/api',
  timeout: 15000, // 15 seconds
});

http.interceptors.request.use((config) => {
  // console.log(config.method);
  // console.log(config.data);
  // console.log(config.url);
  // console.log(config.headers.common);
  return config;
});

http.interceptors.response.use((config) => {
  // console.log(config.status);
  // console.log(config.data);
  // console.log(config.headers);
  // console.log(config.statusText);
  return config;
});

export default http;
