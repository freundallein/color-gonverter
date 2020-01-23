package converter

// Converter32 - convert 32bit ARGB to RGBA
func Converter32(in uint32) (out uint32) {
	alpha := in & 0xFF000000 >> 24
	rgb := in & 0x00FFFFFF << 8
	out = rgb | alpha
	return
}

// Converter64 - convert 64bit ARGBARGB to RGBARGBA (2 uint32-points in 1 uint64)
func Converter64(in uint64) (out uint64) {
	alpha := in & 0xFF000000FF000000 >> 24
	rgb := in & 0x00FFFFFF00FFFFFF << 8
	out = rgb | alpha
	return
}
