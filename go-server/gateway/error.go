package gateway

import "github.com/pkg/errors"

var errIOIsNil error = errors.New("IO is nil")
var errServiceIsNil error = errors.New("Socket service is nil")
var errSocketCannotBeCreated error = errors.New("Socket can not be created")
