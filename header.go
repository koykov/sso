package sso

const mask = 0b01111111

type header uint8

func (h *header) encode(len, flag uint8) {
	*h = header(len & mask)
	*h = *h | header(flag<<7)
}

func (h *header) decode() (uint8, uint8) {
	return uint8(*h & mask), uint8(*h >> 7)
}
