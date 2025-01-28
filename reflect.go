package sso

type sliceh struct {
	data     uintptr
	len, cap int
}

type stringh struct {
	data uintptr
	len  int
}
