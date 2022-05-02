const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const wrapServerWithReflection = require('grpc-node-server-reflection');

/* eslint-disable */
const PROTO_PATH = __dirname + '/protofiles/user.proto';
/* eslint-enable */

class GrpcServer {
  server;

  constructor() {
    this.server = wrapServerWithReflection.default(new grpc.Server());
    return this;
  }

  getProto(protoPath) {
    // using dynamic loading of proto files
    const packageDefinition = protoLoader.loadSync(protoPath, {
      keepCase: true,
      longs: String,
      enums: String,
      defaults: true,
      oneofs: true
    });
    const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
    return protoDescriptor.github.com.juliocesarscheidt.usersmicroservice;
  }

  addUserServices(services) {
    this.server.addService(this.getProto(PROTO_PATH).UserService.service, services);
    return this;
  }

  addUserRateServices(services) {
    this.server.addService(this.getProto(PROTO_PATH).UserRateService.service, services);
    return this;
  }

  start() {
    this.server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
      console.info('[INFO] Server listening on 0.0.0.0:50051');
      this.server.start();
    });
  }
}

module.exports = GrpcServer;
