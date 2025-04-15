package lib

var charset = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'_', '@',
}

func FindPasswordNaive(password string) string {
	word := ""
	for k := 0; k < len(charset); k++ {
		word = string(charset[k])
		if word == password {
			return word
		}
		for i := range charset {
			word = string(charset[k]) + string(charset[i])
			if word == password {
				return word
			}
			for j := range charset {
				word = string(charset[k]) + string(charset[i]) + string(charset[j])
				if word == password {
					return word
				}
				for l := range charset {
					word = string(charset[k]) + string(charset[i]) + string(charset[j]) + string(charset[l])
					if word == password {
						return word
					}
					for m := range charset {
						word = string(charset[k]) + string(charset[i]) + string(charset[j]) + string(charset[l]) + string(charset[m])
						if word == password {
							return word
						}
					}
				}
			}
		}
	}

	return ""
}

func FindPasswordByCombination(password string) string {
	n := len(charset)
	maxLen := len(password)

	for length := 1; length <= maxLen; length++ {
		indices := make([]int, length)

		for {
			word := make([]byte, length)
			for i, idx := range indices {
				word[i] = charset[idx]
			}

			if string(word) == password {
				return string(word)
			}

			i := length - 1
			for i >= 0 && indices[i] == n-1 {
				indices[i] = 0
				i--
			}

			if i < 0 {
				break
			}

			indices[i]++
		}
	}

	return ""
}
