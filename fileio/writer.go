package fileio

import (
	"os"
)

func FileWriteN(filename string, buf []byte) (uint, error) {
	// TODO: Determine if the exists

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var n uint
	for i := 0; i < len(buf); i += DefaultOnceByte {

		o, err := FileWriteOnce(file, buf[i:i+DefaultOnceByte])
		if err != nil {
			panic(err)
		}
		n += o

	}

	return n, nil

}

func FileWriteOnce(file *os.File, oncebuf []byte) (uint, error) {

	n, err := file.Write(oncebuf)
	if err != nil && n != len(oncebuf) {
		return 0, err
	}
	file.Sync()

	return uint(n), nil
}
