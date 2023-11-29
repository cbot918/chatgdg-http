package main

import (
	"os"
	"syscall"
)

func (s *syn) WaitRead(f *os.File) (bool, error) {
	pfd := s.pr.Fd()
	ffd := f.Fd()
	nfd := 1
	if pfd < ffd {
		nfd += int(ffd)
	} else {
		nfd += int(pfd)
	}
	s.m.Lock()
	for {
		var r fdset
		r.Set(ffd)
		r.Set(pfd)
		n, err := syscall.Select(nfd, r.Sys(), nil, nil, nil)
		if err != nil {
			return false, err
		}
		if n > 0 {
			if r.IsSet(pfd) {
				// Command waits for access f.
				s.m.Unlock()
				return false, nil
			}
			return true, nil
		}
	}
}
