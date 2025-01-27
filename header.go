package sso

type bheader struct {
	data     uintptr
	len, cap int
}

type sheader struct {
	data uintptr
	len  int
}
