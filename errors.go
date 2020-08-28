package roundrobin

import "errors"

// ErrorInvalidConfig returned if given config is invalid, i.e. list is empty.
var ErrorInvalidConfig = errors.New("invalid config")
