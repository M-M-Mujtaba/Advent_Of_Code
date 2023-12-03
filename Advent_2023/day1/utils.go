package day1

func resetToZ(bs []byte) {
	for i := range bs {
		if bs[i] != 'z' {
			bs[i] = 'z'
		} else {
			break
		}
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + 32
	}
	return b
}
