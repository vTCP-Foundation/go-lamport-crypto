package crypto

import "errors"

var (
	ErrKeypairAlreadyPresent  = errors.New("key pair is already present")
)

const (
	lamportCExecutable = "lamportc"
)
