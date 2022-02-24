package bowling

import (
	"log"
	"strconv"
	"strings"
)

var symbolToScore = map[string]int{
	"X": 10,
	"/": 10,
	"-": 0,
}

func CalculateBowlingScore(game string) int {
	score := 0
	strikeIdxs := map[int]bool{}
	spareIdxs := map[int]bool{}
	lastNormalScore := 0

	// split game into cells
	for i, cell := range strings.Split(game, " ") {
		// split cells into throws
		for j, t := range cell {
			// convert rune to string
			symbol := string(t)
			throwScore, found := symbolToScore[symbol]
			var err error
			if !found {
				throwScore, err = strconv.Atoi(symbol)
				if err != nil {
					log.Fatal("Could not convert symbol to points", err)
				}
			}

			switch symbol {
			// when there's a strike remember the cell index
			case "X":
				// only add score if not on bonus cells
				if i < 10 {
					score += throwScore
					strikeIdxs[i] = true
				}
			// when there's a spare remember the cell index
			case "/":
				// only add score if not on bonus cells
				if i < 10 {
					// add the rest of pins to the score
					score += throwScore - lastNormalScore
					spareIdxs[i] = true
				}
			// Just add the score
			// keep a note of it in case of a spare
			default:
				score += throwScore
				lastNormalScore = throwScore
			}

			// If last cell was a strike, add this throw again
			if strikeIdxs[i-1] {
				score += throwScore
			}
			// If the second last cell was a strike, add this throw again
			if strikeIdxs[i-2] {
				score += throwScore
			}
			// if not a bonus cell, only check for spares if we're on a first throw
			if i < 10 && j == 0 && spareIdxs[i-1] {
				score += throwScore
			}
		}
	}

	return score
}
