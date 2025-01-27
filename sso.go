package sso

import (
	"math"
	"unsafe"
)

const (
	payload = 15
	maxLen  = math.MaxInt64
)

type String struct {
	buf [payload]byte
	lf_ lf
}

func New(data string) (s String) {
	s.Assign(data)
	return
}

func (s *String) Assign(str string) *String {
	switch l := len(str); {
	case l == 0:
		return s
	case l <= payload:
		// TODO implement copy of payload data
	case l == maxLen:
		panic("SSO: string length must be less than MaxInt64")
	default:
		buf := make([]byte, l)
		copy(buf, str)
		bh := (*bheader)(unsafe.Pointer(&buf))
		sh := (*sheader)(unsafe.Pointer(&s))
		sh.data, sh.len = bh.data, bh.len
	}
	return s
}

func (s *String) Concat(str string) *String {
	return s
}

func (s *String) String() string {
	return string(s.buf[:s.lf_])
}
