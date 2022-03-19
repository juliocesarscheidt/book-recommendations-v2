const BaseController = require('./BaseController');
const UserRateService = require('../service/UserRateService');

const userRateService = new UserRateService();

class UserRateController extends BaseController {
  constructor() {
    super()
  }

  async UpsertUserRate(call, callback) {
    console.info('call.request', call.request);

    const { user_uuid, book_uuid, rate } = call.request.userRateRequest;

    try {
      const response = await userRateService.UpsertUserRate({ _id: user_uuid, book_uuid, rate: parseFloat(rate.toFixed(2)) });
      return callback(null, response);
    } catch (err) {
      return UserRateController.HandleError(err, callback);
    }
  }

  async GetUserRate(call, callback) {
    console.info('call.request', call.request);

    const { user_uuid } = call.request;
    if (!user_uuid) return UserRateController.BadRequest(callback);

    try {
      const response = await userRateService.GetUserRate({ _id: user_uuid });
      return callback(null, response);
    } catch (err) {
      return UserRateController.HandleError(err, callback);
    }
  }

  async DeleteUserRate(call, callback) {
    console.info('call.request', call.request);

    const { user_uuid } = call.request;
    if (!user_uuid) return UserRateController.BadRequest(callback);

    try {
      const response = await userRateService.DeleteUserRate({ _id: user_uuid });
      return callback(null, response);
    } catch (err) {
      return UserRateController.HandleError(err, callback);
    }
  }

  async ListUserRate(call, callback) {
    console.info('call.request', call.request);

    const { page, size } = call.request;

    try {
      const response = await userRateService.ListUserRate(page, size);
      return callback(null, response);
    } catch (err) {
      return UserRateController.HandleError(err, callback);
    }
  }
}

module.exports = UserRateController;
