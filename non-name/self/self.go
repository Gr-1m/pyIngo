package main

import (
	"unsafe"
)

type Self interface {
	Init()
}

type SelfStruct struct {
	selfptr *uintptr
}

func (s *SelfStruct) Init() {
	s.selfptr = (*uintptr)(unsafe.Pointer(&s))
}

func main() {
	var s *uintptr
	s = (*uintptr)(unsafe.Pointer(&s))
}
