package passwordgenerator

type GeneratePasswordOptions struct {
	Digits            bool
	Letters           bool
	Uppercase         bool
	SpecialCharacters bool
}

func GeneratePassword(length int, opt GeneratePasswordOptions) (string, error) {
	// Your code here
	return "", nil
}
