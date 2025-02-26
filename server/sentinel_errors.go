package server

import "errors"

var BadRequestError = errors.New("BAD_REQUEST")
var InternalServerError = errors.New("INTERNAL_SERVER_ERROR")
var MethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")
var OperationConflict = errors.New("OPERATION_CONFLICT")
var ResourceNotFound = errors.New("RESOURCE_NOT_FOUND")
var UnhandledError = errors.New("UNHANDLED_ERROR")
