package main

func generateKeyStream(t, key string) string {
	ks := key
	for i := 0; ; i++ {
		if len(ks) == len(t) {
			break
		}
		if i == len(t) {
			i = 0
		}
		ks += string(ks[i])
	}
	return ks
}
