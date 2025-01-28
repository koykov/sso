package sso

type lf uint8

func (lf_ *lf) encode(len, flag uint8) {
	*lf_ = *lf_ | lf(len&0b01111111)
	*lf_ = *lf_ | lf(flag<<7)
}

func (lf_ *lf) decode() (uint8, uint8) {
	return uint8(*lf_ & 0b01111111), uint8(*lf_ >> 7)
}
