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

func RandomChoices(population interface{}, repeat int) ([]interface{}, error) {
	var (
		ln  int
		buf = make([]interface{}, repeat)
		// Returning local variables here is unsafe and requires optimization
	)
	switch vt := population.(type) {
	case string:
		ln = len(vt)
		for i := range buf {
			buf[i] = string(vt[RandomInt(0, ln)])
		}
	case []interface{}:
		ln = len(vt)
		for i := range buf {
			buf[i] = vt[RandomInt(0, ln)]
		}
	default:
		return nil, errors.New("Unsupported Type")
	}

	return buf, nil
}

func GenRandomString(length int, charset string) string {
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

	return string(buf)

}
