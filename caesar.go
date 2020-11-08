package main

func CaesarCipher(pt string, key int) string {
	b := make([]byte, len(pt))

	for i := 0; i < len(pt); i++ {
		c := pt[i]
		var a byte

		switch {
		case c >= 'a' && c <= 'z':
			a = 97
		case c >= 'A' && c <= 'Z':
			a = 65
		}
		b[i] = a + ((c-a)+byte(key))%26
	}

	return string(b)
}

func CaesarEncrypt(pt string, key int) string {
	return CaesarCipher(pt, key)
}

func CaesarDecrypt(ct string, key int) string {
	return CaesarCipher(ct, -key)
}
