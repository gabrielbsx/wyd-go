package identity

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")

var ErrMissingName = errors.New("missing name")
var ErrInvalidPassword = errors.New("invalid password")
var ErrInvalidEmail = errors.New("invalid email")

var ErrMissingEmail = errors.New("missing email")
var ErrMissingPasswordHash = errors.New("missing password hash")
var ErrMissingRequiredField = errors.New("missing required field")

var ErrInHashPassword = errors.New("error in hash password")
