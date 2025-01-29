package sso

type byteseq interface {
	~string | ~[]byte
}

type sliceh struct {
	data     uintptr
	len, cap int
}

type stringh struct {
	data uintptr
	len  int
}
