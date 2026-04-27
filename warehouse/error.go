package warehouse

import "errors"

var ErrItemAlreadyExist = errors.New("item already exist")
var ErrItemNotFound = errors.New("item not found")
