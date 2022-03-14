class NotFoundException extends Error {
  constructor(err) {
    super(err)
  }
}

module.exports = NotFoundException;
