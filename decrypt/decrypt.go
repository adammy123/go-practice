package decrypt

// encryption logic:
// 1. convert every letter to ASCII value
// 2. add 1 to the first letter
// 3. for every letter from the second to the last letter, add the value of the previous letter
// 4. % 26 each letter until it is in (a-z) (97-122)
// 5. convert back to letters
func Decrypt(word string) string {
	decryptedWord := ""

	// 1. convert letters to ascii
	encryptedAscii := []int{}
	for _, char := range word {
		encryptedAscii = append(encryptedAscii, convertCharToAscii(char))
	}

	// 2. keep track of this
	prevAscii := encryptedAscii[0]

	// 3. convert first letter to original and append to result
	decryptedWord += convertAsciiToString(prevAscii - 1)

	//4. for every next letter, subtract prev ascii and add 26 until it is (a-z)
	for i := 1; i < len(encryptedAscii); i++ {
		currAscii := encryptedAscii[i]
		decryptedWord += convertAsciiToString(currAscii - prevAscii)
		prevAscii = currAscii
	}

	return decryptedWord
}

func convertCharToAscii(char rune) int {
	return int(char)
}

func convertAsciiToString(ascii int) string {
	for {
		if ascii < 97 {
			ascii += 26
		} else {
			break
		}
	}
	return string((rune(ascii)))
}
