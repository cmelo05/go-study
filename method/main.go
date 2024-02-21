package main

import (
	"errors"
	"fmt"
)

type Score struct {
	Id        int
	Home      string
	HomeScore int
	Away      string
	AwayScore int
}

func (sc Score) PrintScore() {
	resultString := ""
	switch {
	case sc.HomeScore == sc.AwayScore:
		resultString = "Teams drawn"
	case sc.HomeScore > sc.AwayScore:
		resultString = "Home team won"
	case sc.HomeScore < sc.AwayScore:
		resultString = "Away team won"
	default:
		resultString = "Unable to get score"
	}

	fmt.Printf("%s.\nGame %d - %s %d x %d %s\n", resultString, sc.Id, sc.Home, sc.HomeScore, sc.AwayScore, sc.Away)
}

func (sc *Score) UpdateHomeScore(score int) error {
	if score < 0 || score <= sc.HomeScore {
		return errors.New("cannot update score for a value smaller than it was before")
	}

	sc.HomeScore = score
	return nil
}

func main() {
	sc := Score{
		Id:        1,
		Home:      "Brasil",
		HomeScore: 2,
		Away:      "Alemanha",
		AwayScore: 0,
	}

	sc.PrintScore()
	if err := sc.UpdateHomeScore(1); err != nil {
		fmt.Printf("Unable to update the score!")
	} else {
		sc.PrintScore()
	}

}
