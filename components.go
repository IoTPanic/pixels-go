package pixels

type Header struct {
	Sync    uint8
	Channel uint8
	LedCnt  uint16
}

type Pixel struct {
	R    uint8
	G    uint8
	B    uint8
	W    uint8
	RGBW bool
}

type Packet struct {
	Header   Header
	Pixels   []Pixel
	Checksum uint8
}
