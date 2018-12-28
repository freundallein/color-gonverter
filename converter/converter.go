package converter

func Converter32(in uint32) (out uint32) {
	// Convert 32bit ARGB to RGBA
	alpha := in & 0xFF000000 >> 24
	rgb := in & 0x00FFFFFF << 8
	out = rgb | alpha
	return
}

func Converter64(in uint64) (out uint64) {
	// Convert 64bit ARGBARGB to RGBARGBA (2 uint32-points in 1 uint64)
	alpha := in & 0xFF000000FF000000 >> 24
	rgb := in & 0x00FFFFFF00FFFFFF << 8
	out = rgb | alpha
	return
}
