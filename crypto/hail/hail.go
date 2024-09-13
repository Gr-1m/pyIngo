package hail

import (
	"math/big"
)

func Hail(n interface{}) (odd, even uint) {
	switch v := n.(type) {
	case int:
		return Hail0(uint64(v))
	case uint64:
		return Hail0(v)
	case *big.Int:
		return HailBig(v)
	default:
		return 0, 0
	}
}

func Hail0(n uint64) (odd, even uint) {
	var s = n

	for odd, even = 0, 0; s != 4; odd++ {
		for ; s&1 == 0; even++ {
			s >>= 1
		}

		s += (s << 1) + 1
	}
	odd--
	return
}

func HailBig(n *big.Int) (odd, even uint) {
	var (
		s        = n
		intOne   = big.NewInt(1)
		intThree = big.NewInt(3)
		intFour  = big.NewInt(4)
	)

	for odd, even = 0, 0; s.Cmp(intFour) != 0; odd++ {
		for ; s.Bit(0) == 0; even++ {
			s = new(big.Int).Rsh(s, 1)
		}

		s = new(big.Int).Mul(s, intThree)
		s = new(big.Int).Add(s, intOne)
	}
	odd--
	return
}

func CalcAverage(arr []int64) int64 {
	var sum int64
	for _, v := range arr {
		sum += v
	}
	return sum / int64(len(arr))
}
