package invsqrt

import "unsafe"

// func Q_rsqrt(number float32)
func Q_rsqrt(number float32) float32 {
	const threehalfs = 1.5
	var (
		i     uint32
		x2, y float32
	)

	x2 = number * 0.5
	y = number

	i = *(*uint32)(unsafe.Pointer(&y)) // evil floating point bit level hacking
	i = 0x5f3758df - (i >> 1)          // what the fuck?
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (threehalfs - (x2 * y * y)) // 1st iteration
	// y = y * (threehalfs - (x2 * y * y)) // 2nd iteration, this can be removed

	return y
}

func Sqrt(number float32) float32 {
	const onehalfs float32 = 0.5
	var (
		i     uint32
		x2, y float32
	)

	x2 = number * onehalfs
	y = number

	i = (uint32)(y)           // evil floating point bit level hacking
	i = 0x1fbd70a4 + (i >> 1) // what the fuck?
	y = (float32)(i)
	y = y*onehalfs - (x2 / y) // 1st iteration
	y = y*onehalfs - (x2 / y) // 2nd iteration, this can improve accuracy
	// y = y*onehalfs - (x2 / y) // 3rd iteration, this can be removed

	return y
}
