package vbc

// Encode32 encodes uint32 slice to variable byte codes.
func Encode32(vs []uint32) []byte {
	sum := 0
	for _, v := range vs {
		sum += estimateByteNum(int(v))
	}
	var offset uint32
	encoded := make([]byte, sum)
	for _, v := range vs {
		encodeVBCodeNumber32(encoded, &offset, v)
	}
	return encoded
}

// Encode64 encodes uint32 slice to variable byte codes.
func Encode64(vs []uint64) []byte {
	sum := 0
	for _, v := range vs {
		sum += estimateByteNum(int(v))
	}
	var offset uint64
	encoded := make([]byte, sum)
	for _, v := range vs {
		encodeVBCodeNumber64(encoded, &offset, v)
	}
	return encoded
}

// Decode32 decodes variable byte codes to uint32 slice.
func Decode32(bs []byte) []uint32 {
	var v uint32
	decoded := []uint32{}
	for _, b := range bs {
		v <<= 7
		v |= uint32(b & 0x7F)
		if 0 != (b & 0x80) {
			decoded = append(decoded, v)
			v = 0
		}
	}
	return decoded
}

// Decode64 decodes variable byte codes to uint64 slice.
func Decode64(bs []byte) []uint64 {
	var v uint64
	decoded := []uint64{}
	for _, b := range bs {
		v <<= 7
		v |= uint64(b & 0x7F)
		if 0 != (b & 0x80) {
			decoded = append(decoded, v)
			v = 0
		}
	}
	return decoded
}

func estimateByteNum(v int) int {
	var num int
	if (1 << 7) > v {
		num = 1
	} else if (1 << 14) > v {
		num = 2
	} else if (1 << 21) > v {
		num = 3
	} else if (1 << 28) > v {
		num = 4
	} else if (1 << 35) > v {
		num = 5
	} else if (1 << 42) > v {
		num = 6
	} else if (1 << 49) > v {
		num = 7
	} else if (1 << 56) > v {
		num = 8
	} else {
		num = 9
	}
	return num
}

func encode32(buf []byte, offset *uint32, v, num uint32) uint32 {
	if num <= 1 {
		return 0
	}
	var i uint32
	for i = 0; i < num-1; i++ {
		buf[*offset+i] = byte((v >> (7 * (num - i - 1))) & 0x7F)
	}
	return num - 1
}

func encode64(buf []byte, offset *uint64, v, num uint64) uint64 {
	if num <= 1 {
		return 0
	}
	var i uint64
	for i = 0; i < num-1; i++ {
		buf[*offset+i] = byte((v >> (7 * (num - i - 1))) & 0x7F)
	}
	return num - 1
}

func encodeVBCodeNumber32(buf []byte, offset *uint32, v uint32) {
	num := uint32(estimateByteNum(int(v)))
	lastOffset := encode32(buf, offset, v, num)
	buf[*offset+lastOffset] = byte((v & 0x7F) | 0x80)
	*offset += num
}

func encodeVBCodeNumber64(buf []byte, offset *uint64, v uint64) {
	num := uint64(estimateByteNum(int(v)))
	lastOffset := encode64(buf, offset, v, num)
	buf[*offset+lastOffset] = byte((v & 0x7F) | 0x80)
	*offset += num
}
