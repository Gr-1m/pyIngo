package random

import (
	"errors"
	"math/rand"
	"time"
)

var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(minNum, maxNum int) (s int) {
	if minNum == 0 {
		return seed.Intn(maxNum)
	}

	for s = 0; s <= minNum; {
		s = seed.Intn(maxNum)
	}
	return s

}

func GetRandBit(bitNum int) int {
	return seed.Intn((1 << bitNum) - 1)
}

// func RandByte(bNum int) []byte {}

func RandomChoice(population interface{}) (interface{}, error) {

	var ln int

	switch vt := population.(type) {
	case string:
		ln = len(vt)
		return string(vt[RandomInt(0, ln)]), nil
	case []interface{}:
		ln = len(vt)
		return vt[RandomInt(0, ln)], nil
	default:
		return nil, errors.New("Unsupported Type")
	}
}

func RandomChoices(population interface{}, repeat int) (cr []interface{}, err error) {

	switch vt := population.(type) {
	case string:
		for _ = range repeat {
			cr = append(cr, string(vt[RandomInt(0, len(vt))]))
		}
	case []interface{}:
		for _ = range repeat {
			cr = append(cr, vt[RandomInt(0, len(vt))])
		}
	default:
		err = errors.ErrUnsupported
		cr = nil
	}

	return
}

func GenRandomString(length int, charset string) (rs string) {
	if length == 0 {
		length = 5
	}
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyz"
	}

	var (
		buf = make([]byte, length)
	)

	for i := range buf {
		buf[i] = charset[RandomInt(0, len(charset))]
	}

	rs = string(buf)
	return
}
