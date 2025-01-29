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

func New[T byteseq](data T) (s String) {
	p := &s
	p = assignByteseq(p, data)
	return
}

func (s *String) Assign(str []byte) *String {
	return assignByteseq(s, str)
}

func (s *String) AssignString(str string) *String {
	return assignByteseq(s, str)
}

func assignByteseq[T byteseq](dst *String, str T) *String {
	switch l := len(str); {
	case l == 0:
		return dst
	case l <= payload:
		copy(dst.buf[:l], str)
		dst.hdr.encode(uint8(l), 1)
	case l == maxLen:
		panic("SSO: string length must be less than MaxInt64")
	default:
		buf := make([]byte, l)
		copy(buf, str)
		bh := (*sliceh)(unsafe.Pointer(&buf))
		sh := (*stringh)(unsafe.Pointer(dst))
		sh.data, sh.len = bh.data, bh.len
	}
	return dst
}

func (s *String) Append(str []byte) *String {
	return appendByteseq(s, str)
}

func (s *String) AppendString(str string) *String {
	return appendByteseq(s, str)
}

func appendByteseq[T byteseq](dst *String, s T) *String {
	n := len(s)
	if n == 0 {
		return dst
	}
	l, f := dst.hdr.decode()
	if f == 1 {
		// SSO enabled
		if n+int(l) <= payload {
			// SSO possible
			copy(dst.buf[l:], s)
			dst.hdr.encode(l+uint8(n), 1)
			return dst
		}
		// SSO impossible
		buf := make([]byte, n+int(l))
		copy(buf, dst.buf[:l])
		copy(buf[l:], s)
		bh := (*sliceh)(unsafe.Pointer(&buf))
		sh := (*stringh)(unsafe.Pointer(dst))
		sh.data, sh.len = bh.data, bh.len
		return dst
	}
	// Regular concat
	bs := *(*string)(unsafe.Pointer(dst))
	buf := make([]byte, n+len(bs))
	copy(buf, bs)
	copy(buf[len(bs):], s)
	bh := (*sliceh)(unsafe.Pointer(&buf))
	sh := (*stringh)(unsafe.Pointer(dst))
	sh.data, sh.len = bh.data, bh.len
	return dst
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
