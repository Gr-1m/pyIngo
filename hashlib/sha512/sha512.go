package sha512

import (
	"crypto/sha512"
	"encoding/hex"
	"pyIngo/hashlib/hash"
)

const (
	// The size of a SHA-512 checksum in bytes.
	Size = 64

	// The blocksize of SHA-512 in bytes.
	BlockSize = 128

	// The name of SHA-512
	Name = "SHA512"
)

type SHA512 struct {
	Bytes []byte

	Hash hash.Hash
}

func New(b []byte) hash.Hashlib {
	return &SHA512{
		Bytes: b,

		Hash: sha512.New(),
	}

}

func (s *SHA512) BlockSize() int {
	return BlockSize
}

func (s SHA512) Copy() *SHA512 {
	h := sha512.New()
	h.Write(s.Bytes)

	var s1 = &SHA512{
		s.Bytes,
		h,
	}
	return s1

}

func (s *SHA512) Digest() []byte {
	s.Hash.Reset()
	s.Hash.Write(s.Bytes)
	return s.Hash.Sum(nil)
}

func (s *SHA512) Hexdigest() string {
	s.Hash.Reset()
	s.Hash.Write(s.Bytes)
	return hex.EncodeToString(s.Hash.Sum(nil))
}

func (s *SHA512) Name() string {
	return Name
}

func (s *SHA512) Update(b []byte) {
	if b == nil {
		return
	}
	s.Bytes = append(s.Bytes, b...)
}
