package crypto

import "errors"

var (
	ErrValidation             = errors.New("validation error")
	ErrKeypairAlreadyPresent  = errors.New("key pair is already present")
	ErrInvalidKeysPoolDirPath = errors.New("invalid keys pool directory path")
)

const (
	lamportCExecutable = "lamportc"
)
