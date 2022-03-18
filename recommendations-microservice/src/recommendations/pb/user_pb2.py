# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b'\n\nuser.proto\x12.github.com.juliocesarscheidt.usersmicroservice"y\n\x04User\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07surname\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\r\n\x05phone\x18\x05 \x01(\t\x12\x12\n\ncreated_at\x18\x06 \x01(\t\x12\x12\n\nupdated_at\x18\x07 \x01(\t"\\\n\x0bUserRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07surname\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\r\n\x05phone\x18\x04 \x01(\t\x12\x10\n\x08password\x18\x05 \x01(\t"e\n\x11\x43reateUserRequest\x12P\n\x0buserRequest\x18\x01 \x01(\x0b\x32;.github.com.juliocesarscheidt.usersmicroservice.UserRequest""\n\x12\x43reateUserResponse\x12\x0c\n\x04uuid\x18\x01 \x01(\t"\x1e\n\x0eGetUserRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t"U\n\x0fGetUserResponse\x12\x42\n\x04user\x18\x01 \x01(\x0b\x32\x34.github.com.juliocesarscheidt.usersmicroservice.User"s\n\x11UpdateUserRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12P\n\x0buserRequest\x18\x02 \x01(\x0b\x32;.github.com.juliocesarscheidt.usersmicroservice.UserRequest"\x14\n\x12UpdateUserResponse"!\n\x11\x44\x65leteUserRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t"\x14\n\x12\x44\x65leteUserResponse"-\n\x0fListUserRequest\x12\x0c\n\x04page\x18\x01 \x01(\x03\x12\x0c\n\x04size\x18\x02 \x01(\x03"W\n\x10ListUserResponse\x12\x43\n\x05users\x18\x01 \x03(\x0b\x32\x34.github.com.juliocesarscheidt.usersmicroservice.User"[\n\nUserSignUp\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07surname\x18\x02 \x01(\t\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\r\n\x05phone\x18\x04 \x01(\t\x12\x10\n\x08password\x18\x05 \x01(\t"c\n\x11UserSignUpRequest\x12N\n\nuserSignUp\x18\x01 \x01(\x0b\x32:.github.com.juliocesarscheidt.usersmicroservice.UserSignUp"*\n\x12UserSignUpResponse\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\t"-\n\nUserSignIn\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t"c\n\x11UserSignInRequest\x12N\n\nuserSignIn\x18\x01 \x01(\x0b\x32:.github.com.juliocesarscheidt.usersmicroservice.UserSignIn"*\n\x12UserSignInResponse\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x01 \x01(\t"\'\n\x04Rate\x12\x11\n\tbook_uuid\x18\x02 \x01(\t\x12\x0c\n\x04rate\x18\x03 \x01(\x02"\x8a\x01\n\x08UserRate\x12\x11\n\tuser_uuid\x18\x01 \x01(\t\x12\x43\n\x05rates\x18\x02 \x03(\x0b\x32\x34.github.com.juliocesarscheidt.usersmicroservice.Rate\x12\x12\n\ncreated_at\x18\x03 \x01(\t\x12\x12\n\nupdated_at\x18\x04 \x01(\t"E\n\x0fUserRateRequest\x12\x11\n\tuser_uuid\x18\x01 \x01(\t\x12\x11\n\tbook_uuid\x18\x02 \x01(\t\x12\x0c\n\x04rate\x18\x03 \x01(\x02"q\n\x15UpsertUserRateRequest\x12X\n\x0fuserRateRequest\x18\x01 \x01(\x0b\x32?.github.com.juliocesarscheidt.usersmicroservice.UserRateRequest"+\n\x16UpsertUserRateResponse\x12\x11\n\tuser_uuid\x18\x01 \x01(\t"\'\n\x12GetUserRateRequest\x12\x11\n\tuser_uuid\x18\x01 \x01(\t"a\n\x13GetUserRateResponse\x12J\n\x08userRate\x18\x01 \x01(\x0b\x32\x38.github.com.juliocesarscheidt.usersmicroservice.UserRate"*\n\x15\x44\x65leteUserRateRequest\x12\x11\n\tuser_uuid\x18\x01 \x01(\t"\x18\n\x16\x44\x65leteUserRateResponse"1\n\x13ListUserRateRequest\x12\x0c\n\x04page\x18\x01 \x01(\x03\x12\x0c\n\x04size\x18\x02 \x01(\x03"c\n\x14ListUserRateResponse\x12K\n\tuserRates\x18\x01 \x03(\x0b\x32\x38.github.com.juliocesarscheidt.usersmicroservice.UserRate2\x98\x08\n\x0bUserService\x12\x93\x01\n\nUserSignUp\x12\x41.github.com.juliocesarscheidt.usersmicroservice.UserSignUpRequest\x1a\x42.github.com.juliocesarscheidt.usersmicroservice.UserSignUpResponse\x12\x93\x01\n\nUserSignIn\x12\x41.github.com.juliocesarscheidt.usersmicroservice.UserSignInRequest\x1a\x42.github.com.juliocesarscheidt.usersmicroservice.UserSignInResponse\x12\x93\x01\n\nCreateUser\x12\x41.github.com.juliocesarscheidt.usersmicroservice.CreateUserRequest\x1a\x42.github.com.juliocesarscheidt.usersmicroservice.CreateUserResponse\x12\x8a\x01\n\x07GetUser\x12>.github.com.juliocesarscheidt.usersmicroservice.GetUserRequest\x1a?.github.com.juliocesarscheidt.usersmicroservice.GetUserResponse\x12\x93\x01\n\nUpdateUser\x12\x41.github.com.juliocesarscheidt.usersmicroservice.UpdateUserRequest\x1a\x42.github.com.juliocesarscheidt.usersmicroservice.UpdateUserResponse\x12\x93\x01\n\nDeleteUser\x12\x41.github.com.juliocesarscheidt.usersmicroservice.DeleteUserRequest\x1a\x42.github.com.juliocesarscheidt.usersmicroservice.DeleteUserResponse\x12\x8d\x01\n\x08ListUser\x12?.github.com.juliocesarscheidt.usersmicroservice.ListUserRequest\x1a@.github.com.juliocesarscheidt.usersmicroservice.ListUserResponse2\x8a\x05\n\x0fUserRateService\x12\x9f\x01\n\x0eUpsertUserRate\x12\x45.github.com.juliocesarscheidt.usersmicroservice.UpsertUserRateRequest\x1a\x46.github.com.juliocesarscheidt.usersmicroservice.UpsertUserRateResponse\x12\x96\x01\n\x0bGetUserRate\x12\x42.github.com.juliocesarscheidt.usersmicroservice.GetUserRateRequest\x1a\x43.github.com.juliocesarscheidt.usersmicroservice.GetUserRateResponse\x12\x9f\x01\n\x0e\x44\x65leteUserRate\x12\x45.github.com.juliocesarscheidt.usersmicroservice.DeleteUserRateRequest\x1a\x46.github.com.juliocesarscheidt.usersmicroservice.DeleteUserRateResponse\x12\x99\x01\n\x0cListUserRate\x12\x43.github.com.juliocesarscheidt.usersmicroservice.ListUserRateRequest\x1a\x44.github.com.juliocesarscheidt.usersmicroservice.ListUserRateResponseB\x11Z\x0f./protofiles;pbb\x06proto3'
)


_USER = DESCRIPTOR.message_types_by_name["User"]
_USERREQUEST = DESCRIPTOR.message_types_by_name["UserRequest"]
_CREATEUSERREQUEST = DESCRIPTOR.message_types_by_name["CreateUserRequest"]
_CREATEUSERRESPONSE = DESCRIPTOR.message_types_by_name["CreateUserResponse"]
_GETUSERREQUEST = DESCRIPTOR.message_types_by_name["GetUserRequest"]
_GETUSERRESPONSE = DESCRIPTOR.message_types_by_name["GetUserResponse"]
_UPDATEUSERREQUEST = DESCRIPTOR.message_types_by_name["UpdateUserRequest"]
_UPDATEUSERRESPONSE = DESCRIPTOR.message_types_by_name["UpdateUserResponse"]
_DELETEUSERREQUEST = DESCRIPTOR.message_types_by_name["DeleteUserRequest"]
_DELETEUSERRESPONSE = DESCRIPTOR.message_types_by_name["DeleteUserResponse"]
_LISTUSERREQUEST = DESCRIPTOR.message_types_by_name["ListUserRequest"]
_LISTUSERRESPONSE = DESCRIPTOR.message_types_by_name["ListUserResponse"]
_USERSIGNUP = DESCRIPTOR.message_types_by_name["UserSignUp"]
_USERSIGNUPREQUEST = DESCRIPTOR.message_types_by_name["UserSignUpRequest"]
_USERSIGNUPRESPONSE = DESCRIPTOR.message_types_by_name["UserSignUpResponse"]
_USERSIGNIN = DESCRIPTOR.message_types_by_name["UserSignIn"]
_USERSIGNINREQUEST = DESCRIPTOR.message_types_by_name["UserSignInRequest"]
_USERSIGNINRESPONSE = DESCRIPTOR.message_types_by_name["UserSignInResponse"]
_RATE = DESCRIPTOR.message_types_by_name["Rate"]
_USERRATE = DESCRIPTOR.message_types_by_name["UserRate"]
_USERRATEREQUEST = DESCRIPTOR.message_types_by_name["UserRateRequest"]
_UPSERTUSERRATEREQUEST = DESCRIPTOR.message_types_by_name["UpsertUserRateRequest"]
_UPSERTUSERRATERESPONSE = DESCRIPTOR.message_types_by_name["UpsertUserRateResponse"]
_GETUSERRATEREQUEST = DESCRIPTOR.message_types_by_name["GetUserRateRequest"]
_GETUSERRATERESPONSE = DESCRIPTOR.message_types_by_name["GetUserRateResponse"]
_DELETEUSERRATEREQUEST = DESCRIPTOR.message_types_by_name["DeleteUserRateRequest"]
_DELETEUSERRATERESPONSE = DESCRIPTOR.message_types_by_name["DeleteUserRateResponse"]
_LISTUSERRATEREQUEST = DESCRIPTOR.message_types_by_name["ListUserRateRequest"]
_LISTUSERRATERESPONSE = DESCRIPTOR.message_types_by_name["ListUserRateResponse"]
User = _reflection.GeneratedProtocolMessageType(
    "User",
    (_message.Message,),
    {
        "DESCRIPTOR": _USER,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.User)
    },
)
_sym_db.RegisterMessage(User)

UserRequest = _reflection.GeneratedProtocolMessageType(
    "UserRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserRequest)
    },
)
_sym_db.RegisterMessage(UserRequest)

CreateUserRequest = _reflection.GeneratedProtocolMessageType(
    "CreateUserRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _CREATEUSERREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.CreateUserRequest)
    },
)
_sym_db.RegisterMessage(CreateUserRequest)

CreateUserResponse = _reflection.GeneratedProtocolMessageType(
    "CreateUserResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _CREATEUSERRESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.CreateUserResponse)
    },
)
_sym_db.RegisterMessage(CreateUserResponse)

GetUserRequest = _reflection.GeneratedProtocolMessageType(
    "GetUserRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETUSERREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.GetUserRequest)
    },
)
_sym_db.RegisterMessage(GetUserRequest)

GetUserResponse = _reflection.GeneratedProtocolMessageType(
    "GetUserResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETUSERRESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.GetUserResponse)
    },
)
_sym_db.RegisterMessage(GetUserResponse)

UpdateUserRequest = _reflection.GeneratedProtocolMessageType(
    "UpdateUserRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _UPDATEUSERREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UpdateUserRequest)
    },
)
_sym_db.RegisterMessage(UpdateUserRequest)

UpdateUserResponse = _reflection.GeneratedProtocolMessageType(
    "UpdateUserResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _UPDATEUSERRESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UpdateUserResponse)
    },
)
_sym_db.RegisterMessage(UpdateUserResponse)

DeleteUserRequest = _reflection.GeneratedProtocolMessageType(
    "DeleteUserRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _DELETEUSERREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.DeleteUserRequest)
    },
)
_sym_db.RegisterMessage(DeleteUserRequest)

DeleteUserResponse = _reflection.GeneratedProtocolMessageType(
    "DeleteUserResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _DELETEUSERRESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.DeleteUserResponse)
    },
)
_sym_db.RegisterMessage(DeleteUserResponse)

ListUserRequest = _reflection.GeneratedProtocolMessageType(
    "ListUserRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _LISTUSERREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.ListUserRequest)
    },
)
_sym_db.RegisterMessage(ListUserRequest)

ListUserResponse = _reflection.GeneratedProtocolMessageType(
    "ListUserResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _LISTUSERRESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.ListUserResponse)
    },
)
_sym_db.RegisterMessage(ListUserResponse)

UserSignUp = _reflection.GeneratedProtocolMessageType(
    "UserSignUp",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERSIGNUP,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserSignUp)
    },
)
_sym_db.RegisterMessage(UserSignUp)

UserSignUpRequest = _reflection.GeneratedProtocolMessageType(
    "UserSignUpRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERSIGNUPREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserSignUpRequest)
    },
)
_sym_db.RegisterMessage(UserSignUpRequest)

UserSignUpResponse = _reflection.GeneratedProtocolMessageType(
    "UserSignUpResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERSIGNUPRESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserSignUpResponse)
    },
)
_sym_db.RegisterMessage(UserSignUpResponse)

UserSignIn = _reflection.GeneratedProtocolMessageType(
    "UserSignIn",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERSIGNIN,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserSignIn)
    },
)
_sym_db.RegisterMessage(UserSignIn)

UserSignInRequest = _reflection.GeneratedProtocolMessageType(
    "UserSignInRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERSIGNINREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserSignInRequest)
    },
)
_sym_db.RegisterMessage(UserSignInRequest)

UserSignInResponse = _reflection.GeneratedProtocolMessageType(
    "UserSignInResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERSIGNINRESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserSignInResponse)
    },
)
_sym_db.RegisterMessage(UserSignInResponse)

Rate = _reflection.GeneratedProtocolMessageType(
    "Rate",
    (_message.Message,),
    {
        "DESCRIPTOR": _RATE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.Rate)
    },
)
_sym_db.RegisterMessage(Rate)

UserRate = _reflection.GeneratedProtocolMessageType(
    "UserRate",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERRATE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserRate)
    },
)
_sym_db.RegisterMessage(UserRate)

UserRateRequest = _reflection.GeneratedProtocolMessageType(
    "UserRateRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _USERRATEREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UserRateRequest)
    },
)
_sym_db.RegisterMessage(UserRateRequest)

UpsertUserRateRequest = _reflection.GeneratedProtocolMessageType(
    "UpsertUserRateRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _UPSERTUSERRATEREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UpsertUserRateRequest)
    },
)
_sym_db.RegisterMessage(UpsertUserRateRequest)

UpsertUserRateResponse = _reflection.GeneratedProtocolMessageType(
    "UpsertUserRateResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _UPSERTUSERRATERESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.UpsertUserRateResponse)
    },
)
_sym_db.RegisterMessage(UpsertUserRateResponse)

GetUserRateRequest = _reflection.GeneratedProtocolMessageType(
    "GetUserRateRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETUSERRATEREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.GetUserRateRequest)
    },
)
_sym_db.RegisterMessage(GetUserRateRequest)

GetUserRateResponse = _reflection.GeneratedProtocolMessageType(
    "GetUserRateResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _GETUSERRATERESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.GetUserRateResponse)
    },
)
_sym_db.RegisterMessage(GetUserRateResponse)

DeleteUserRateRequest = _reflection.GeneratedProtocolMessageType(
    "DeleteUserRateRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _DELETEUSERRATEREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.DeleteUserRateRequest)
    },
)
_sym_db.RegisterMessage(DeleteUserRateRequest)

DeleteUserRateResponse = _reflection.GeneratedProtocolMessageType(
    "DeleteUserRateResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _DELETEUSERRATERESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.DeleteUserRateResponse)
    },
)
_sym_db.RegisterMessage(DeleteUserRateResponse)

ListUserRateRequest = _reflection.GeneratedProtocolMessageType(
    "ListUserRateRequest",
    (_message.Message,),
    {
        "DESCRIPTOR": _LISTUSERRATEREQUEST,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.ListUserRateRequest)
    },
)
_sym_db.RegisterMessage(ListUserRateRequest)

ListUserRateResponse = _reflection.GeneratedProtocolMessageType(
    "ListUserRateResponse",
    (_message.Message,),
    {
        "DESCRIPTOR": _LISTUSERRATERESPONSE,
        "__module__": "user_pb2"
        # @@protoc_insertion_point(class_scope:github.com.juliocesarscheidt.usersmicroservice.ListUserRateResponse)
    },
)
_sym_db.RegisterMessage(ListUserRateResponse)

_USERSERVICE = DESCRIPTOR.services_by_name["UserService"]
_USERRATESERVICE = DESCRIPTOR.services_by_name["UserRateService"]
if _descriptor._USE_C_DESCRIPTORS == False:

    DESCRIPTOR._options = None
    DESCRIPTOR._serialized_options = b"Z\017./protofiles;pb"
    _USER._serialized_start = 62
    _USER._serialized_end = 183
    _USERREQUEST._serialized_start = 185
    _USERREQUEST._serialized_end = 277
    _CREATEUSERREQUEST._serialized_start = 279
    _CREATEUSERREQUEST._serialized_end = 380
    _CREATEUSERRESPONSE._serialized_start = 382
    _CREATEUSERRESPONSE._serialized_end = 416
    _GETUSERREQUEST._serialized_start = 418
    _GETUSERREQUEST._serialized_end = 448
    _GETUSERRESPONSE._serialized_start = 450
    _GETUSERRESPONSE._serialized_end = 535
    _UPDATEUSERREQUEST._serialized_start = 537
    _UPDATEUSERREQUEST._serialized_end = 652
    _UPDATEUSERRESPONSE._serialized_start = 654
    _UPDATEUSERRESPONSE._serialized_end = 674
    _DELETEUSERREQUEST._serialized_start = 676
    _DELETEUSERREQUEST._serialized_end = 709
    _DELETEUSERRESPONSE._serialized_start = 711
    _DELETEUSERRESPONSE._serialized_end = 731
    _LISTUSERREQUEST._serialized_start = 733
    _LISTUSERREQUEST._serialized_end = 778
    _LISTUSERRESPONSE._serialized_start = 780
    _LISTUSERRESPONSE._serialized_end = 867
    _USERSIGNUP._serialized_start = 869
    _USERSIGNUP._serialized_end = 960
    _USERSIGNUPREQUEST._serialized_start = 962
    _USERSIGNUPREQUEST._serialized_end = 1061
    _USERSIGNUPRESPONSE._serialized_start = 1063
    _USERSIGNUPRESPONSE._serialized_end = 1105
    _USERSIGNIN._serialized_start = 1107
    _USERSIGNIN._serialized_end = 1152
    _USERSIGNINREQUEST._serialized_start = 1154
    _USERSIGNINREQUEST._serialized_end = 1253
    _USERSIGNINRESPONSE._serialized_start = 1255
    _USERSIGNINRESPONSE._serialized_end = 1297
    _RATE._serialized_start = 1299
    _RATE._serialized_end = 1338
    _USERRATE._serialized_start = 1341
    _USERRATE._serialized_end = 1479
    _USERRATEREQUEST._serialized_start = 1481
    _USERRATEREQUEST._serialized_end = 1550
    _UPSERTUSERRATEREQUEST._serialized_start = 1552
    _UPSERTUSERRATEREQUEST._serialized_end = 1665
    _UPSERTUSERRATERESPONSE._serialized_start = 1667
    _UPSERTUSERRATERESPONSE._serialized_end = 1710
    _GETUSERRATEREQUEST._serialized_start = 1712
    _GETUSERRATEREQUEST._serialized_end = 1751
    _GETUSERRATERESPONSE._serialized_start = 1753
    _GETUSERRATERESPONSE._serialized_end = 1850
    _DELETEUSERRATEREQUEST._serialized_start = 1852
    _DELETEUSERRATEREQUEST._serialized_end = 1894
    _DELETEUSERRATERESPONSE._serialized_start = 1896
    _DELETEUSERRATERESPONSE._serialized_end = 1920
    _LISTUSERRATEREQUEST._serialized_start = 1922
    _LISTUSERRATEREQUEST._serialized_end = 1971
    _LISTUSERRATERESPONSE._serialized_start = 1973
    _LISTUSERRATERESPONSE._serialized_end = 2072
    _USERSERVICE._serialized_start = 2075
    _USERSERVICE._serialized_end = 3123
    _USERRATESERVICE._serialized_start = 3126
    _USERRATESERVICE._serialized_end = 3776
# @@protoc_insertion_point(module_scope)
