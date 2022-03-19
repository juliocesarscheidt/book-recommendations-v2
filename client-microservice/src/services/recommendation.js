import http from '../utils/http.js';

const getRecommendations = async (uuid) => http
  .get(`/recommendation/user/${uuid}`)
  .then((response) => {
    if (!response.data) {
      return null
    }
    return response.data.data
  });

export {
  getRecommendations,
}
