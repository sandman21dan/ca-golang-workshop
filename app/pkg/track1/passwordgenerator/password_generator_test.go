package passwordgenerator_test

import (
	"regexp"
	"testing"

	"git.ivxs.uk/daniel.perez/ca-golang-workshop/app/pkg/track1/passwordgenerator"
)

func TestGeneratePassword(t *testing.T) {
	suite := []struct {
		name         string
		len          int
		opt          passwordgenerator.GeneratePasswordOptions
		expected     []string
		notExpected  []string
		expectToFail bool
	}{
		{
			name: "should have only 6 digits",
			len:  6,
			opt: passwordgenerator.GeneratePasswordOptions{
				Digits: true,
			},
			expected:    []string{`\d{6}`},
			notExpected: []string{`[a-zA-Z$%&*#@]+`},
		},
		{
			name: "should have only 8 letters",
			len:  8,
			opt: passwordgenerator.GeneratePasswordOptions{
				Letters: true,
			},
			expected:    []string{`[a-z]{8}`},
			notExpected: []string{`[\dA-Z$%&*#@]+`},
		},
		{
			name: "should have only 12 letters at least one lowercase and one uppercase",
			len:  12,
			opt: passwordgenerator.GeneratePasswordOptions{
				Letters:   true,
				Uppercase: true,
			},
			expected:    []string{`[a-zA-Z]{12}`, `[a-z]+`, `[A-Z]+`},
			notExpected: []string{`[\d$%&*#@]+`},
		},
		{
			name: "should have only 12 letters at least one lowercase and one uppercase",
			len:  10,
			opt: passwordgenerator.GeneratePasswordOptions{
				Digits:            true,
				Letters:           true,
				Uppercase:         true,
				SpecialCharacters: true,
			},
			expected:    []string{`[a-zA-Z\d$%&*#@]{10}`, `[a-z]+`, `[A-Z]+`, `\d+`, `[$%&*#@]+`},
			notExpected: []string{},
		},
		{
			name: "errors when no arguments are true",
			opt: passwordgenerator.GeneratePasswordOptions{
				Digits:            false,
				Letters:           false,
				Uppercase:         false,
				SpecialCharacters: false,
			},
			expectToFail: true,
		},
	}

	for _, c := range suite {
		t.Run(c.name, func(t *testing.T) {
			pwd, err := passwordgenerator.GeneratePassword(c.len, c.opt)

			if err != nil {
				if c.expectToFail {
					return
				}

				t.Errorf("Unexpected error, got %s", err)
			} else if c.expectToFail {
				t.Error("Expected to fail but didn't")
			}

			if len(pwd) != c.len {
				t.Errorf("Expected length %d, got %d", c.len, len(pwd))
			}

			for _, p := range c.expected {
				matched, err := regexp.Match(p, []byte(pwd))
				if err != nil {
					t.Errorf("Error matching regexp: %s", err)
				}
				if !matched {
					t.Errorf("Expected %s, to match %s expression", pwd, p)
				}
			}

			if len(c.notExpected) > 0 {
				for _, p := range c.notExpected {
					matched, err := regexp.Match(p, []byte(pwd))
					if err != nil {
						t.Errorf("Error matching regexp: %s", err)
					}
					if matched {
						t.Errorf("Expected %s, NOT to match %s expression", pwd, p)
					}
				}
			}
		})
	}
}
