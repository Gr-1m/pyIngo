package util

import (
	"math/big"
)

var (
	intZero = big.NewInt(0)
	intOne  = big.NewInt(1)
)

func GCD(a, b *big.Int) (m *big.Int) {
	m = new(big.Int).GCD(nil, nil, a, b)
	return
}

func XGCD(a, b *big.Int) (m [3]*big.Int) {
	for i := range 3 {
		m[i] = new(big.Int)
	}
	m[0].GCD(m[1], m[2], a, b)
	return
}

func LCM(a, b *big.Int) (m *big.Int) {

	m = new(big.Int)
	g := new(big.Int).GCD(nil, nil, a, b)
	if g.Cmp(intOne) == 0 {
		return m.Mul(a, b)
	}

	ga := new(big.Int).Div(a, g)
	gb := new(big.Int).Div(b, g)
	return m.Mul(LCM(ga, gb), g)
	// return m.Mul(ga, gb)

}

func BytetoLong(b []byte) (m *big.Int) {
	m = new(big.Int)
	m.SetBytes(b)
	return m
}

func LongtoByte(x *big.Int) []byte {
	return x.Bytes()
}
