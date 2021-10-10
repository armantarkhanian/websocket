package websocket

import "errors"

var (
	ErrMissingConfig                  = errors.New("missing config")
	ErrInvalidAuthenticationKeyLength = errors.New("authenticationKey length must be 32 or 64")
	ErrInvalidEncryptionKeyLength     = errors.New("encryptionKey length must be 16, 24 or 32")
	ErrMissingEngine                  = errors.New("missing Engine in websocket config")
	ErrInvalidClientHandler           = errors.New("invalid clientHandler")
	ErrInvalidNodeHandler             = errors.New("invalid nodeHandler")
)
