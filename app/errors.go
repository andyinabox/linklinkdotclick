package app

import "errors"

var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrServerError = errors.New("server error")
var ErrMissingUserId = errors.New("recieved zero-value user id")
