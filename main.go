package main

import (
	"fmt"
	"strings"
)

type rotor struct {
	wiring   string
	position int
}

var reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {

	rotorI := rotor{wiring: "EKMFLGDQVZNTOWYHXUSPAIBRCJ"}
	rotorII := rotor{wiring: "AJDKSIRUXBLHWTMCQGZNPYFVOE"}
	rotorIII := rotor{wiring: "BDFHJLCPRTXVZNYEIWGAKMUSQO"}
	rotorI.position = 0
	rotorII.position = 0
	rotorIII.position = 0

	plaintext := "HELLO"
	encryptedText := enigmaEncrypt(plaintext, rotorI, rotorII, rotorIII)
	fmt.Println("Plaintext: ", plaintext)
	fmt.Println("Encrypted Text: ", encryptedText)

	rotorI.position = 0
	rotorII.position = 0
	rotorIII.position = 0
	decryptedText := enigmaDecrypt(encryptedText, rotorI, rotorII, rotorIII)
	fmt.Println("Decrypted Text: ", decryptedText)
}

func enigmaEncrypt(plaintext string, rotors ...rotor) string {
	plaintext = strings.ToUpper(plaintext)
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// Rotate rotors before encryption
			rotateRotors(&rotors)

			// Pass the character through the rotors from right to left
			char = substitute(char, rotors[2])
			char = substitute(char, rotors[1])
			char = substitute(char, rotors[0])

			// Pass the character through the reflector
			char = reflector(char)

			// Pass the character through the rotors from left to right
			char = substitute(char, rotors[0])
			char = substitute(char, rotors[1])
			char = substitute(char, rotors[2])

			encrypted.WriteRune(char)
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func enigmaDecrypt(plaintext string, rotors ...rotor) string {
	plaintext = strings.ToUpper(plaintext)
	var encrypted strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			// Rotate rotors before encryption
			rotateRotors(&rotors)

			char = decrypt(char, rotors[2])
			char = decrypt(char, rotors[1])
			char = decrypt(char, rotors[0])
			char = reflector(char)
			char = decrypt(char, rotors[0])
			char = decrypt(char, rotors[1])
			char = decrypt(char, rotors[2])

			encrypted.WriteRune(char)
		} else {
			// Non-alphabetic characters are not modified
			encrypted.WriteRune(char)
		}
	}

	return encrypted.String()
}

func rotateRotors(rotors *[]rotor) {
	for i := 0; i < len(*rotors); i++ {
		(*rotors)[i].position++
		if (*rotors)[i].position >= 26 {
			(*rotors)[i].position = 0
		}
	}
}

func substitute(char rune, rotor rotor) rune {
	index := (int(char-'A') + rotor.position) % 26
	return rune(rotor.wiring[index])
}

func decrypt(char rune, rotor rotor) rune {
	index := strings.IndexRune(rotor.wiring, char)
	return rune(alphabet[index])
}

func reflector(char rune) rune {
	index := strings.IndexRune(reflectorB, char)
	return rune(alphabet[index])
}
