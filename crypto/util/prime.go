package util

import (
	"crypto/rand"
	"math/big"
)

func IsPrime(p *big.Int) bool {
	return p.ProbablyPrime(20)
}

func GetPrime(nbit int) (p *big.Int) {
	p, _ = rand.Prime(rand.Reader, nbit)
	return
}

func NextPrime(p *big.Int) (q *big.Int) {
	two := big.NewInt(2)
	q = new(big.Int).Add(p, big.NewInt(1))
	q.SetBit(q, 0, 1)

	for !q.ProbablyPrime(20) {
		q.Add(q, two)
	}

	// to be optimized
	return
}

func NephiInverse(e, p, q *big.Int) *big.Int {
	intOne := big.NewInt(1)
	phiN := new(big.Int).Mul(new(big.Int).Sub(p, intOne), new(big.Int).Sub(q, intOne))

	return new(big.Int).ModInverse(e, phiN)
}
