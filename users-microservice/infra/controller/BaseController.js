const { status } = require('@grpc/grpc-js');

const NotFoundException = require('../../application/exception/NotFoundException');
const InvalidRateException = require('../../application/exception/InvalidRateException');
const InvalidEmailPasswordException = require('../../application/exception/InvalidEmailPasswordException');
const DuplicateEmailException = require('../../application/exception/DuplicateEmailException');

class BaseController {
  constructor() {
  }

  static BadRequest(callback, message = 'Bad Request') {
    return callback({ code: 400, message, status: status.INVALID_ARGUMENT }, null);
  }

  static InternalServerError(callback, message = 'Internal Server Error') {
    return callback({ code: 500, message, status: status.INTERNAL }, null);
  }

  static NotFound(callback, message = 'Not Found') {
    return callback({ code: 404, message, status: status.NOT_FOUND }, null);
  }

  static HandleError(err, callback) {
    console.error(err);

    if (err instanceof InvalidRateException) {
      return BaseController.BadRequest(callback, err.message);

    } else if (err instanceof InvalidEmailPasswordException) {
      return BaseController.BadRequest(callback, err.message);

    } else if (err instanceof DuplicateEmailException) {
      return BaseController.BadRequest(callback, err.message);

    } else if (err instanceof NotFoundException) {
      return BaseController.NotFound(callback, err.message);

    } else {
      return BaseController.InternalServerError(callback);
    }
  }
}

module.exports = BaseController;
