var path = require('path');
var grpc = require('grpc');

const PROTO_PATH = path.resolve(__dirname, '../proto/hello.proto');
const ns = grpc.load(PROTO_PATH, 'proto', { convertFieldsToCamelCase: true });

const defaultClient = new ns.Hello(
    process.env.SERVICE_ENDPOINT,
    grpc.credentials.createInsecure()
);

Object.getOwnPropertyNames(ns.Hello.service).forEach(method => {
  ns.Hello.service[method].path = ns.Hello.service[method].path.split('/').slice(1).join('.');
});

const modifiedClientConstructor = grpc.makeGenericClientConstructor(ns.Hello.service);
const modifiedClient = new modifiedClientConstructor(
    process.env.SERVICE_ENDPOINT,
    grpc.credentials.createInsecure()
);

defaultClient.greet({}, (err, res) => {
  printResults('DEFAULT CLIENT', res, err);

  modifiedClient.greet({}, (err, res) => {
    printResults('MODIFIED CLIENT', res, err);
    process.exit(0);
  });
});

function printResults(name, res, err) {
  console.log();
  console.log(`RESPONSE FROM ${name}:`);
  console.dir(res);
  console.log();
  console.log('ERROR FROM ${name}:');
  console.dir(err);
}

