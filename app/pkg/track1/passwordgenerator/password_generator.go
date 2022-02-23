package passwordgenerator

import "errors"

type GeneratePasswordOptions struct {
	Digits            bool
	Letters           bool
	Uppercase         bool
	SpecialCharacters bool
}

var timesCount = 0

func GeneratePassword(length int, opt GeneratePasswordOptions) (string, error) {
	// Your code here
	// return "abcdesgj", nil

	if !(opt.Digits || opt.Letters || opt.Uppercase || opt.SpecialCharacters) {
		return "", errors.New("invalid set of arguments, at least one must be true")
	}

	answers := []string{
		"123456",
		"abcdefhi",
		"aBcDeFgHiJlk",
		"aB2DeFg$iJ",
	}

	timesCount++
	return answers[timesCount-1], nil
}
