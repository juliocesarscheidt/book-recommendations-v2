const { status } = require('@grpc/grpc-js');

const NotFoundException = require('../../application/exception/NotFoundException');
const InvalidRateException = require('../../application/exception/InvalidRateException');

class BaseController {
  constructor() {
  }

  static BadRequest(callback) {
    return callback({ code: 400, message: 'Bad Request', status: status.INVALID_ARGUMENT }, null);
  }

  static InternalServerError(callback) {
    return callback({ code: 500, message: 'Internal Server Error', status: status.INTERNAL }, null);
  }

  static NotFound(callback) {
    return callback({ code: 404, message: 'Not Found', status: status.NOT_FOUND }, null);
  }

  static HandleError(err, callback) {
    console.error(err);
    if (err instanceof NotFoundException) {
      return BaseController.NotFound(callback);

    } else if (err instanceof InvalidRateException) {
      return BaseController.BadRequest(callback);

    } else {
      return BaseController.InternalServerError(callback);
    }
  }
}

module.exports = BaseController;
