// source: tokenlog.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.protos.TokenLogPB', null, global);
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
proto.protos.TokenLogPB = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.protos.TokenLogPB, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.protos.TokenLogPB.displayName = 'proto.protos.TokenLogPB';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.protos.TokenLogPB.prototype.toObject = function(opt_includeInstance) {
  return proto.protos.TokenLogPB.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.protos.TokenLogPB} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.TokenLogPB.toObject = function(includeInstance, msg) {
  var f, obj = {
    tokenlogid: jspb.Message.getFieldWithDefault(msg, 1, ""),
    fromusername: jspb.Message.getFieldWithDefault(msg, 2, ""),
    tousername: jspb.Message.getFieldWithDefault(msg, 3, ""),
    tokenname: jspb.Message.getFieldWithDefault(msg, 4, ""),
    tokenamount: jspb.Message.getFieldWithDefault(msg, 5, 0),
    comment: jspb.Message.getFieldWithDefault(msg, 6, ""),
    timestamp: jspb.Message.getFieldWithDefault(msg, 7, 0)
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
 * @return {!proto.protos.TokenLogPB}
 */
proto.protos.TokenLogPB.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.protos.TokenLogPB;
  return proto.protos.TokenLogPB.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.protos.TokenLogPB} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.protos.TokenLogPB}
 */
proto.protos.TokenLogPB.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setTokenlogid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFromusername(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setTousername(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setTokenname(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTokenamount(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setComment(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTimestamp(value);
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
proto.protos.TokenLogPB.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.protos.TokenLogPB.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.protos.TokenLogPB} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.protos.TokenLogPB.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTokenlogid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFromusername();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTousername();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getTokenname();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getTokenamount();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getComment();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getTimestamp();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
};


/**
 * optional string TokenLogId = 1;
 * @return {string}
 */
proto.protos.TokenLogPB.prototype.getTokenlogid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.TokenLogPB} returns this
 */
proto.protos.TokenLogPB.prototype.setTokenlogid = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string FromUserName = 2;
 * @return {string}
 */
proto.protos.TokenLogPB.prototype.getFromusername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.TokenLogPB} returns this
 */
proto.protos.TokenLogPB.prototype.setFromusername = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string ToUserName = 3;
 * @return {string}
 */
proto.protos.TokenLogPB.prototype.getTousername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.TokenLogPB} returns this
 */
proto.protos.TokenLogPB.prototype.setTousername = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string TokenName = 4;
 * @return {string}
 */
proto.protos.TokenLogPB.prototype.getTokenname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.TokenLogPB} returns this
 */
proto.protos.TokenLogPB.prototype.setTokenname = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int64 TokenAmount = 5;
 * @return {number}
 */
proto.protos.TokenLogPB.prototype.getTokenamount = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.protos.TokenLogPB} returns this
 */
proto.protos.TokenLogPB.prototype.setTokenamount = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string Comment = 6;
 * @return {string}
 */
proto.protos.TokenLogPB.prototype.getComment = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.protos.TokenLogPB} returns this
 */
proto.protos.TokenLogPB.prototype.setComment = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional int64 Timestamp = 7;
 * @return {number}
 */
proto.protos.TokenLogPB.prototype.getTimestamp = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {number} value
 * @return {!proto.protos.TokenLogPB} returns this
 */
proto.protos.TokenLogPB.prototype.setTimestamp = function(value) {
  return jspb.Message.setProto3IntField(this, 7, value);
};


goog.object.extend(exports, proto.protos);
