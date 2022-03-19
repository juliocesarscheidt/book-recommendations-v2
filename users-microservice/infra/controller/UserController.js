const BaseController = require('./BaseController');
const UserService = require('../service/UserService');
const UserRateService = require('../service/UserRateService');

const userService = new UserService();
const userRateService = new UserRateService();

class UserController extends BaseController {
  constructor() {
    super()
  }

  async UserSignUp(call, callback) {
    console.info('call.request', call.request);

    const { name, surname, email, phone, password } = call.request.userSignUp;

    try {
      const response = await userService.UserSignUp({ name, surname, email, phone, password });
      return callback(null, response);
    } catch (err) {
      return UserController.HandleError(err, callback);
    }
  }

  async UserSignIn(call, callback) {
    console.info('call.request', call.request);

    const { email, password } = call.request.userSignIn;

    try {
      const response = await userService.UserSignIn({ email, password });
      return callback(null, response);
    } catch (err) {
      return UserController.HandleError(err, callback);
    }
  }

  async CreateUser(call, callback) {
    console.info('call.request', call.request);

    const { name, surname, email, phone, password } = call.request.userRequest;

    try {
      const response = await userService.CreateUser({ name, surname, email, phone, password });
      return callback(null, response);
    } catch (err) {
      return UserController.HandleError(err, callback);
    }
  }

  async GetUser(call, callback) {
    console.info('call.request', call.request);

    const { uuid } = call.request;
    if (!uuid) return UserController.BadRequest(callback);

    try {
      const response = await userService.GetUser({ _id: uuid });
      return callback(null, response);
    } catch (err) {
      return UserController.HandleError(err, callback);
    }
  }

  async UpdateUser(call, callback) {
    console.info('call.request', call.request);

    const { uuid } = call.request;
    if (!uuid) return UserController.BadRequest(callback);
    const payload = call.request.userRequest;
    Object.keys(payload).forEach(key => {
      if (!payload[key]) delete payload[key];
    });

    try {
      const response = await userService.UpdateUser({ _id: uuid }, payload);
      return callback(null, response);
    } catch (err) {
      return UserController.HandleError(err, callback);
    }
  }

  async DeleteUser(call, callback) {
    console.info('call.request', call.request);

    const { uuid } = call.request;
    if (!uuid) return UserController.BadRequest(callback);

    try {
      const response = await userService.DeleteUser({ _id: uuid });

      // delete user rates as well
      await userRateService.DeleteUserRate({ _id: uuid });

      return callback(null, response);
    } catch (err) {
      return UserController.HandleError(err, callback);
    }
  }

  async ListUser(call, callback) {
    console.info('call.request', call.request);

    const { page, size } = call.request;

    try {
      const response = await userService.ListUser(page, size);
      return callback(null, response);
    } catch (err) {
      return UserController.HandleError(err, callback);
    }
  }
}

module.exports = UserController;
