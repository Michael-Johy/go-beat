package number

import "unicode"

func Find(b []byte) (int, int) {
	i := 0
	for ; i < len(b) && !unicode.IsDigit(rune(b[i])); i++ {
	}

	res := 0
	for ; i < len(b) && unicode.IsDigit(rune(b[i])); i++ {
		res = res*10 + int(b[i])
	}
	return res, i
}
