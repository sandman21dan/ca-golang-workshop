package passwordgenerator

import (
	"errors"
	"math/rand"
	"unicode"
)

type GeneratePasswordOptions struct {
	Digits            bool
	Letters           bool
	Uppercase         bool
	SpecialCharacters bool
}

var digits = []rune("0123456789")
var letters = []rune("abcdefghijklmnopqrstuvwxyz")
var special = []rune("$%&*#@")

const (
	digit = iota
	letter
	uppercase
	specialCharacter
)

func GeneratePassword(length int, opt GeneratePasswordOptions) (string, error) {
	if !(opt.Digits || opt.Letters || opt.SpecialCharacters) {
		return "", errors.New("invalid set of arguments, at least one must be true")
	}

	password := []rune{}

	optPool := []int{}
	if opt.Digits {
		optPool = append(optPool, digit)
	}
	if opt.Letters {
		optPool = append(optPool, letter)
	}
	if opt.SpecialCharacters {
		optPool = append(optPool, specialCharacter)
	}
	if opt.Uppercase {
		optPool = append(optPool, uppercase)
	}

	for i := 0; i < length; i++ {
		randomOpt := optPool[rand.Intn(len(optPool))]

		switch randomOpt {
		case digit:
			password = append(password, digits[rand.Intn(len(digits))])
		case letter:
			password = append(password, letters[rand.Intn(len(letters))])
		case specialCharacter:
			password = append(password, special[rand.Intn(len(special))])
		case uppercase:
			letter := letters[rand.Intn(len(letters))]
			password = append(password, unicode.ToUpper(letter))
		}
	}

	return string(password), nil
}
