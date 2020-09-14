package consts

import "fmt"

// ErrEntityNotFound is returned (wrapped) when an entity is not found
var ErrEntityNotFound = fmt.Errorf("entity not found")
