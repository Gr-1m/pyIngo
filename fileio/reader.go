package fileio

import (
	// "bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

func FileRead(filename string, once uint) ([]byte, error) {
	// TODO: Determine if the exists
	s, err := os.Stat(filename)
	if err != nil || s.IsDir() {
		return nil, errors.New("File not Found or IsDir")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, once)
	if once == 0 {
		buf, err = FileReadOnce(file, DefaultOnceByte)
		if err != nil {
			return nil, err
		}

	} else {
		buf, err = FileReadOnce(file, once)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil

}

func FileReadN(filename string) ([][]byte, error) {
	// TODO: Determine if the exists

	s, err := os.Stat(filename)
	if err != nil || s.IsDir() {
		return nil, errors.New("File not Found or IsDir")
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	r := make([][]byte, 0)
	for {
		buf, err := FileReadOnce(file, DefaultOnceByte)
		if err != nil {
			panic(err)
		}

		r = append(r, buf)
	}

	return r, nil

}

func FileReadOnce(file *os.File, oncebyte uint) ([]byte, error) {

	// r := bufio.NewReader(file)

	buf := make([]byte, oncebyte)

	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	buf = bytes.TrimRight(buf, "\x00")

	return buf, nil
}
