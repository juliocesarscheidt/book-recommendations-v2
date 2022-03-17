const GrpcServer = require('./GrpcServer');
const UserController = require('./infra/controller/UserController')
const UserRateController = require('./infra/controller/UserRateController')

const userController = new UserController();
const userRateController = new UserRateController();

new GrpcServer()
  .addUserServices({
    UserSignIn: userController.UserSignIn,
    UserSignUp: userController.UserSignUp,
    CreateUser: userController.CreateUser,
    GetUser: userController.GetUser,
    UpdateUser: userController.UpdateUser,
    DeleteUser: userController.DeleteUser,
    ListUser: userController.ListUser,
  })
  .addUserRateServices({
    UpsertUserRate: userRateController.UpsertUserRate,
    GetUserRate: userRateController.GetUserRate,
    DeleteUserRate: userRateController.DeleteUserRate,
    ListUserRate: userRateController.ListUserRate,
  })
  .start();
