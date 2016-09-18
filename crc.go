package elkm1gold

func crc(msg []byte) (crc byte) {
	for _, b := range msg {
		crc = crc + b
	}
	crc = (((crc & 0xff) ^ 0xff) + 1) & 0xff
	return
}
