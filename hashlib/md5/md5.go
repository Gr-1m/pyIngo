package md5

import (
	"crypto/md5"
	"encoding/hex"
	"pyIngo/hashlib/hash"
)

const (
	// The size of an MD5 checksum in bytes.
	Size = 16

	// The blocksize of MD5 in bytes.
	BlockSize = 64

	// The name of MD5
	Name = "MD5"
)

type Md5 struct {
	Bytes []byte

	Hash hash.Hash
}

func New(b []byte) hash.Hashlib {
	return &Md5{
		Bytes: b,
		Hash:  md5.New(),
	}
}

func (m *Md5) BlockSize() int {
	return BlockSize
}

func (m Md5) Copy() *Md5 {
	h := md5.New()
	h.Write(m.Bytes)

	var m1 = &Md5{
		m.Bytes,
		h,
	}
	return m1
}

func (m *Md5) Digest() []byte {
	m.Hash.Reset()
	m.Hash.Write(m.Bytes)
	return m.Hash.Sum(nil)
}

func (m *Md5) Hexdigest() string {
	m.Hash.Reset()
	m.Hash.Write(m.Bytes)
	return hex.EncodeToString(m.Hash.Sum(nil))
}

func (m Md5) Name() string {
	return Name
}

func (m *Md5) Update(b []byte) {
	if b == nil {
		return
	}
	m.Bytes = append(m.Bytes, b...)
}
