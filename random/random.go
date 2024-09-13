package random

import (
	// crand "crypto/rand"
	"errors"
	"math/rand"
	"time"
)

var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandBit(bitNum int) int {
	return seed.Intn((1 << bitNum) - 1)
}

func RandomInt(minNum, maxNum int) (s int) {
	if minNum == 0 {
		return seed.Intn(maxNum)
	}

	for s = 0; s <= minNum; {
		s = seed.Intn(maxNum)
	}
	return s

}

func RandByte(bNum int) (b []byte) {
	b = make([]byte, bNum, bNum)
	_, err := rand.Read(b)
	if err != nil {
		return nil
	}
	return
}

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

func RandomChoices(population interface{}, repeat int) (choices []interface{}, err error) {

	switch vt := population.(type) {
	case string:
		for _ = range repeat {
			choices = append(choices, string(vt[RandomInt(0, len(vt))]))
		}
	case []interface{}:
		for _ = range repeat {
			choices = append(choices, vt[RandomInt(0, len(vt))])
		}
	default:
		err = errors.New("Unsupported Type")
		choices = nil
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
