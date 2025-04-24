package socket

import "github.com/pkg/errors"

var ErrNameNotFound error = errors.New("no name associated with the provided socket ID")
