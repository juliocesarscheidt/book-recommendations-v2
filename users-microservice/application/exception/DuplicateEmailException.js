class DuplicateEmailException extends Error {
  constructor(err) {
    super(err)
  }
}

module.exports = DuplicateEmailException;
