package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"pyIngo/hashlib/hash"
)

const (
	// The size of a SHA-1 checksum in bytes.
	Size = 20

	// The blocksize of SHA-1 in bytes.
	BlockSize = 64

	// The name of SHA1
	Name = "SHA1"
)

type SHA1 struct {
	Bytes []byte

	Hash hash.Hash
}

func New(b []byte) hash.Hashlib {
	return &SHA1{
		Bytes: b,

		Hash: sha1.New(),
	}
}

func (s *SHA1) BlockSize() int {
	return BlockSize
}

func (s SHA1) Copy() *SHA1 {
	h := sha1.New()
	h.Write(s.Bytes)

	var s1 = &SHA1{
		s.Bytes,
		h,
	}
	return s1
}

func (s *SHA1) Digest() []byte {
	s.Hash.Reset()
	s.Hash.Write(s.Bytes)
	return s.Hash.Sum(nil)
}

func (s *SHA1) Hexdigest() string {
	s.Hash.Reset()
	s.Hash.Write(s.Bytes)
	return hex.EncodeToString(s.Hash.Sum(nil))
}

func (s *SHA1) Name() string {
	return Name
}

func (s *SHA1) Update(b []byte) {
	if b == nil {
		return
	}
	s.Bytes = append(s.Bytes, b...)
}
