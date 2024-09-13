package pwntools

// package pwn

import ()

type Session struct {
	LHOST, RHOST string
	LPORT, RPORT uint16
}

func Remote() (s *Session, err error) {
	return
}
