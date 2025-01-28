package sso

import "unsafe"

const (
	intSz   = 32 << (^uint(0) >> 63)
	maxLen  = 1<<(intSz-1) - 1
	payload = intSz/4 - 1
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
		copy(s.buf[:l], str)
		s.lf_.encode(uint8(l), 1)
	case l == maxLen:
		panic("SSO: string length must be less than MaxInt64")
	default:
		buf := make([]byte, l)
		copy(buf, str)
		bh := (*bheader)(unsafe.Pointer(&buf))
		sh := (*sheader)(unsafe.Pointer(s))
		sh.data, sh.len = bh.data, bh.len
	}
	return s
}

func (s *String) Concat(str string) *String {
	return s
}

func (s *String) String() string {
	if l, flag := s.lf_.decode(); flag == 1 {
		var h sheader
		h.data = uintptr(unsafe.Pointer(&s.buf))
		h.len = int(l)
		return *(*string)(unsafe.Pointer(&h))
	}
	return *(*string)(unsafe.Pointer(s))
}
