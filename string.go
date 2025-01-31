package sso

import "unsafe"

const (
	intSz   = 32 << (^uint(0) >> 63)
	maxLen  = 1<<(intSz-1) - 1
	payload = intSz/4 - 1
)

type String string

type byteseq interface {
	~string | ~[]byte
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

func (s *String) Append(str []byte) *String {
	return appendByteseq(s, str)
}

func (s *String) AppendString(str string) *String {
	return appendByteseq(s, str)
}

func (s *String) Reset() *String {
	s.header().hdr.encode(0, 1)
	return s
}

func (s *String) Bytes() []byte {
	ss := s.String()
	sh := (*stringh)(unsafe.Pointer(&ss))
	var bh sliceh
	bh.data, bh.len, bh.cap = sh.data, sh.len, sh.len
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func (s *String) String() string {
	sh := s.header()
	if l, flag := sh.hdr.decode(); flag == 1 {
		var h stringh
		h.data = uintptr(unsafe.Pointer(&sh.buf))
		h.len = int(l)
		return *(*string)(unsafe.Pointer(&h))
	}
	return *(*string)(s)
}

func (s *String) header() *ssoheader {
	return (*ssoheader)(unsafe.Pointer(s))
}

func assignByteseq[T byteseq](dst *String, str T) *String {
	switch l := len(str); {
	case l == 0:
		return dst
	case l <= payload:
		h := dst.header()
		copy(h.buf[:l], str)
		h.hdr.encode(uint8(l), 1)
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

func appendByteseq[T byteseq](dst *String, s T) *String {
	n := len(s)
	if n == 0 {
		return dst
	}
	h := dst.header()
	l, f := h.hdr.decode()
	if n+int(l) >= maxLen {
		panic("SSO: string length must be less than MaxInt64")
	}
	if f == 1 {
		// SSO enabled
		if n+int(l) <= payload {
			// SSO possible
			copy(h.buf[l:], s)
			h.hdr.encode(l+uint8(n), 1)
			return dst
		}
		// SSO impossible
		buf := make([]byte, n+int(l))
		copy(buf, h.buf[:l])
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
