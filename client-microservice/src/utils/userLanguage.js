const setLocalStorageLanguage = (language) => {
  localStorage.setItem('language', language);
};

const getLocalStorageLanguage = () => {
  const language = localStorage.getItem('language');
  return language || 'pt-br';
};

export {
  setLocalStorageLanguage,
  getLocalStorageLanguage,
};
