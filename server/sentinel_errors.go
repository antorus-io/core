package server

import "errors"

var ErrBadRequest = errors.New("BAD_REQUEST")
var ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
var ErrMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")
var ErrOperationConflict = errors.New("OPERATION_CONFLICT")
var ErrResourceNotFound = errors.New("RESOURCE_NOT_FOUND")
var ErrUnhandledError = errors.New("UNHANDLED_ERROR")
