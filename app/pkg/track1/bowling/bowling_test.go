package bowling_test

import (
	"testing"

	"git.ivxs.uk/daniel.perez/ca-golang-workshop/app/pkg/track1/bowling"
)

func TestCalculateBowlingScore(t *testing.T) {
	cases := []struct {
		name          string
		game          string
		expectedScore int
	}{
		{
			name:          "all gutter game",
			game:          "-- -- -- -- -- -- -- -- -- --",
			expectedScore: 0,
		},
		{
			name:          "no strike or spare",
			game:          "6- 3- 5- 3- 6- 10- 8- 1- 9- 0-",
			expectedScore: 51,
		},
		{
			name:          "perfect game",
			game:          "X- X- X- X- X- X- X- X- X- XX X-",
			expectedScore: 300,
		},
		{
			name:          "all spares with a final 5",
			game:          "5/ 5/ 5/ 5/ 5/ 5/ 5/ 5/ 5/ 5/ 5",
			expectedScore: 150,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			score := bowling.CalculateBowlingScore(c.game)
			if score != c.expectedScore {
				t.Errorf("Expected %d, got %d", c.expectedScore, score)
			}
		})
	}
}
