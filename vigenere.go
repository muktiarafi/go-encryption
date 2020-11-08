package main

func VigenereEncrypt(pt, key string) string {
	var res string
	ks := generateKeyStream(pt, key)

	for i := 0; i < len(pt); i++ {
		c := (pt[i] + ks[i]) % 26
		c += 65

		res += string(c)
	}
	return res
}

// func VigenereEncrypt(pt, key string) string {
// 	var res string
// 	ks := generateKeyStream(pt, key)
// 	for i, j := 0, 0; i < len(pt); i++ {
// 		c := pt[i]

// 		switch {
// 		case c >= 'a' && c <= 'z':
// 			c = c + 'A' - 'a'
// 		case c < 'A' || c > 'Z':
// 			continue
// 		}
// 		res += string((c+ks[j]-2*'A')%26 + 'A')
// 		j = (j + 1) % len(ks)
// 	}
// 	return res
// }

func VigenereDecrypt(ct, key string) string {
	var res string
	ks := generateKeyStream(ct, key)

	for i := 0; i < len(ct); i++ {
		c := (ct[i] - ks[i] + 26) % 26
		c += 65

		res += string(c)
	}
	return res
}
