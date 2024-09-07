package hash

import (
	"hash"
)

type Hash = hash.Hash

type Hashlib interface {
	// hash.Hash

	// Copy() *Hashlib

	Digest() []byte

	Hexdigest() string

	Name() string

	Update([]byte)
}
