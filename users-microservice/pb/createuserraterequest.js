/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest');


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.displayName = 'proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    userraterequest: (f = msg.getUserraterequest()) && proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest;
  return proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest;
      reader.readMessage(value,proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest.deserializeBinaryFromReader);
      msg.setUserraterequest(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserraterequest();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest.serializeBinaryToWriter
    );
  }
};


/**
 * optional UserRateRequest userRateRequest = 1;
 * @return {?proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.prototype.getUserraterequest = function() {
  return /** @type{?proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest} */ (
    jspb.Message.getWrapperField(this, proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest, 1));
};


/** @param {?proto.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest|undefined} value */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.prototype.setUserraterequest = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.prototype.clearUserraterequest = function() {
  this.setUserraterequest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.CreateUserRateRequest.prototype.hasUserraterequest = function() {
  return jspb.Message.getField(this, 1) != null;
};


