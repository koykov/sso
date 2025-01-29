package sso

import "unsafe"

const (
	intSz   = 32 << (^uint(0) >> 63)
	maxLen  = 1<<(intSz-1) - 1
	payload = intSz/4 - 1
)

type String struct {
	buf [payload]byte
	hdr header
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
		s.hdr.encode(uint8(l), 1)
	case l == maxLen:
		panic("SSO: string length must be less than MaxInt64")
	default:
		buf := make([]byte, l)
		copy(buf, str)
		bh := (*sliceh)(unsafe.Pointer(&buf))
		sh := (*stringh)(unsafe.Pointer(s))
		sh.data, sh.len = bh.data, bh.len
	}
	return s
}

func (s *String) Append(str string) *String {
	n := len(str)
	l, f := s.hdr.decode()
	if f == 1 {
		// SSO enabled
		if n+int(l) <= payload {
			// SSO possible
			copy(s.buf[l:], str)
			s.hdr.encode(l+uint8(n), 1)
			return s
		}
		// SSO impossible
		buf := make([]byte, n+int(l))
		copy(buf, s.buf[:l])
		copy(buf[l:], str)
		bh := (*sliceh)(unsafe.Pointer(&buf))
		sh := (*stringh)(unsafe.Pointer(s))
		sh.data, sh.len = bh.data, bh.len
		return s
	}
	// Regular concat
	bs := *(*string)(unsafe.Pointer(s))
	buf := make([]byte, n+len(bs))
	copy(buf, bs)
	copy(buf[len(bs):], str)
	bh := (*sliceh)(unsafe.Pointer(&buf))
	sh := (*stringh)(unsafe.Pointer(s))
	sh.data, sh.len = bh.data, bh.len
	return s
}

func (s *String) Reset() *String {
	s.hdr.encode(0, 1)
	return s
}

func (s *String) String() string {
	if l, flag := s.hdr.decode(); flag == 1 {
		var h stringh
		h.data = uintptr(unsafe.Pointer(&s.buf))
		h.len = int(l)
		return *(*string)(unsafe.Pointer(&h))
	}
	return *(*string)(unsafe.Pointer(s))
}
