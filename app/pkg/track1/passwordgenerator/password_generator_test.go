package passwordgenerator_test

import (
	"regexp"
	"testing"

	"git.ivxs.uk/daniel.perez/ca-golang-workshop/app/pkg/track1/passwordgenerator"
)

func TestGeneratePassword(t *testing.T) {
	suite := []struct {
		name        string
		len         int
		opt         passwordgenerator.GeneratePasswordOptions
		expected    string
		notExpected string
		expectedLen int
	}{
		{
			name: "should have only 6 digits",
			len:  6,
			opt: passwordgenerator.GeneratePasswordOptions{
				Digits: true,
			},
			expected:    `\d{6}`,
			notExpected: `[a-zA-Z$%&*#@]+`,
			expectedLen: 6,
		},
		{
			name: "should have only 8 letters",
			len:  8,
			opt: passwordgenerator.GeneratePasswordOptions{
				Digits: true,
			},
			expected:    `[a-z]{6}`,
			notExpected: `[\dA-Z$%&*#@]+`,
			expectedLen: 8,
		},
	}

	for _, c := range suite {
		t.Run(c.name, func(t *testing.T) {
			pwd := passwordgenerator.GeneratePassword(c.len, c.opt)
			if len(pwd) != c.expectedLen {
				t.Errorf("Expected length %d, got %d", c.expectedLen, len(pwd))
			}

			matched, err := regexp.Match(c.expected, []byte(pwd))
			if err != nil {
				t.Errorf("Error matching regexp: %s", err)
			}
			if !matched {
				t.Errorf("Expected %s, to match %s expression", pwd, c.expected)
			}

			if len(c.notExpected) > 0 {
				matched, err := regexp.Match(c.notExpected, []byte(pwd))
				if err != nil {
					t.Errorf("Error matching regexp: %s", err)
				}
				if matched {
					t.Errorf("Expected %s, NOT to match %s expression", pwd, c.expected)
				}
			}
		})
	}
}
