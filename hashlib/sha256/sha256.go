package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"pyIngo/hashlib/hash"
)

const (
	// The size of a SHA-256 checksum in bytes.
	Size = 32

	// The size of a SHA224 checksum in bytes.
	Size224 = 28

	// The blocksize of SHA-256 in bytes.
	BlockSize = 64

	// The name of SHA-256
	Name = "SHA256"
)

type SHA256 struct {
	Bytes []byte

	Hash hash.Hash
}

func New(b []byte) hash.Hashlib {
	return &SHA256{
		Bytes: b,

		Hash: sha256.New(),
	}
}

func (s *SHA256) BlockSize() int {
	return BlockSize
}

func (s SHA256) Copy() *SHA256 {
	h := sha256.New()
	h.Write(s.Bytes)

	var s1 = &SHA256{
		s.Bytes,
		h,
	}
	return s1
}

func (s *SHA256) Digest() []byte {
	s.Hash.Reset()
	s.Hash.Write(s.Bytes)
	return s.Hash.Sum(nil)
}

func (s *SHA256) Hexdigest() string {
	s.Hash.Reset()
	s.Hash.Write(s.Bytes)
	return hex.EncodeToString(s.Hash.Sum(nil))
}

func (s *SHA256) Name() string {
	return Name
}

func (s *SHA256) Update(b []byte) {
	if b == nil {
		return
	}
	s.Bytes = append(s.Bytes, b...)
}
