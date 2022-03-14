/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.github.com.juliocesarscheidt.usersmicroservice.User');


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
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.displayName = 'proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse';
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
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    user: (f = msg.getUser()) && proto.github.com.juliocesarscheidt.usersmicroservice.User.toObject(includeInstance, f)
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
 * @return {!proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse;
  return proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.github.com.juliocesarscheidt.usersmicroservice.User;
      reader.readMessage(value,proto.github.com.juliocesarscheidt.usersmicroservice.User.deserializeBinaryFromReader);
      msg.setUser(value);
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
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.github.com.juliocesarscheidt.usersmicroservice.User.serializeBinaryToWriter
    );
  }
};


/**
 * optional User user = 1;
 * @return {?proto.github.com.juliocesarscheidt.usersmicroservice.User}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.prototype.getUser = function() {
  return /** @type{?proto.github.com.juliocesarscheidt.usersmicroservice.User} */ (
    jspb.Message.getWrapperField(this, proto.github.com.juliocesarscheidt.usersmicroservice.User, 1));
};


/** @param {?proto.github.com.juliocesarscheidt.usersmicroservice.User|undefined} value */
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.prototype.setUser = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.prototype.clearUser = function() {
  this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse.prototype.hasUser = function() {
  return jspb.Message.getField(this, 1) != null;
};


