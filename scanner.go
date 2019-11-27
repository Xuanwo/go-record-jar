package recordjar

import (
	"bufio"
)

type scanner struct {
	s    *bufio.Scanner
	used bool
}

func (s *scanner) Scan() bool {
	if !s.used {
		return true
	}
	ok := s.s.Scan()
	if ok {
		s.used = false
	}
	return ok
}

func (s *scanner) Text() string {
	s.used = true
	return s.s.Text()
}
