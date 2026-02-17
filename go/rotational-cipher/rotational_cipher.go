package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {

	var new string
	shiftKey = shiftKey % 26

	for _, r := range plain {
		if r >= 65 && r <= 90 {
			n := rune(int(r) + shiftKey)
			if n > 90 {
				n = 65 + n - 90 - 1
			}
			new += string(n)
			continue
		}

		if r >= 97 && r <= 122 {
			n := rune(int(r) + shiftKey)
			if n > 122 {
				n = 97 + n - 122 - 1
			}
			new += string(n)
			continue
		}
		new += string(r)
	}

	return new
}
