package sso

type lf uint8

func (lf_ *lf) setl(len uint8) {
	*lf_ = *lf_ | lf(len&0b11111110)
}

func (lf_ *lf) setf(flag uint8) {
	*lf_ = *lf_ | lf(flag<<7)
}

func (lf_ *lf) getl() uint8 {
	return uint8(*lf_ & 0b11111110)
}

func (lf_ *lf) getf() uint8 {
	return uint8(*lf_ >> 7)
}
