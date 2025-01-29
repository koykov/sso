package sso

type header uint8

func (h *header) encode(len, flag uint8) {
	*h = header(len & 0b01111111)
	*h = *h | header(flag<<7)
}

func (h *header) decode() (uint8, uint8) {
	return uint8(*h & 0b01111111), uint8(*h >> 7)
}
