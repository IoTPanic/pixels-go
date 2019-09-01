package pixels

func marshalPacket(pkt Packet, RGBW bool) []uint8 {
	pyld := make([]uint8, 5)

	pyld[0] = 0x50
	pyld[1] = pkt.Header.Sync
	pyld[2] = pkt.Header.Channel
	var tmp uint8
	tmp = uint8(pkt.Header.LedCnt)
	pyld[3] = tmp
	tmp = uint8(pkt.Header.LedCnt >> 8)
	pyld[4] = tmp
	for _, p := range pkt.Pixels {
		if RGBW {
			pyld = append(pyld, p.R)
			pyld = append(pyld, p.G)
			pyld = append(pyld, p.B)
			pyld = append(pyld, p.W)
		} else {
			pyld = append(pyld, p.R)
			pyld = append(pyld, p.G)
			pyld = append(pyld, p.B)
		}
	}
	// TODO CALC CHECKSUM
	return pyld
}
