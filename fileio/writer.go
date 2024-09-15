package fileio

import (
	"errors"
	"os"
)

func FileWriteN(filename string, buf []byte) (uint, error) {
	// TODO: Determine if the exists
	s, err := os.Stat(filename)
	if err != nil || s.IsDir() {
		return 0, errors.New("File not Found or IsDir")
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var n uint
	for i := 0; i < len(buf); i += DefaultOnceByte {

		if len(buf) < DefaultOnceByte {
			o, _ := WriteOnce(file, buf[:DefaultOnceByte])
			n += o
			break
		}

		o, err := WriteOnce(file, buf[:DefaultOnceByte])
		if err != nil {
			panic(err)
		}
		buf = buf[DefaultOnceByte:]
		n += o
	}

	return n, nil

}

func WriteOnce(file *os.File, oncebuf []byte) (uint, error) {

	n, err := file.Write(oncebuf)
	if err != nil && n != len(oncebuf) {
		return 0, err
	}

	return uint(n), file.Sync()
}
