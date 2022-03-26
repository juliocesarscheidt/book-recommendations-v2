const formatPhone = (phone) => {
  if (!phone || typeof phone !== 'string') {
    return '';
  }

  let returnPhone = phone.toString().replace(/[^\d]+/g, '');
  if (returnPhone.length > 0) {
    returnPhone = `(${returnPhone.substring(0)}`;
  }

  if (returnPhone.length > 3) {
    returnPhone = `${returnPhone.substring(0, 3)})${returnPhone.substring(3)}`;
  }

  if (returnPhone.length > 8) {
    if (returnPhone.length === 11) {
      returnPhone = `${returnPhone.substring(0, 7)}-${returnPhone.substring(7, 11)}`;
    } else if (returnPhone.length === 12) {
      returnPhone = `${returnPhone.substring(0, 8)}-${returnPhone.substring(8, 12)}`;
    } else {
      returnPhone = `${returnPhone.substring(0, 9)}-${returnPhone.substring(9, 13)}`;
    }
  }

  return returnPhone;
};

const capitalize = (str) => {
  if (!str || typeof str !== 'string') {
    return '';
  }
  const stringSplitted = str.split(' ')
    .map(name => name.charAt(0).toUpperCase() + name.substring(1).toLowerCase());
  return stringSplitted.join(' ');
};

const trimLetters = (str, amount) => `${str.slice(0, amount)}...`;

const currency = value => `R$ ${parseFloat(value).toFixed(2).toLocaleString()}`;

export {
  formatPhone,
  capitalize,
  trimLetters,
  currency,
}
