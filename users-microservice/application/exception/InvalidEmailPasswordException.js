class InvalidEmailPasswordException extends Error {
  constructor(err) {
    super(err)
  }
}

module.exports = InvalidEmailPasswordException;
