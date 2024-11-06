package main

import (
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	league := League{wins: make(map[TeamName]int)}

	league.MatchResult("A", 1, "B", 0)
	league.MatchResult("A", 1, "C", 0)
	league.MatchResult("D", 1, "C", 0)

	RankPrinter(league, os.Stdout)
}

type TeamName string

type Team struct {
	name        TeamName
	playerNames []string
}

type League struct {
	wins map[TeamName]int
}

func (l League) MatchResult(firstTeam TeamName, firstTeamScore int, secondTeam TeamName, secondTeamScore int) {
	winner, loser := winnerTeam(firstTeam, firstTeamScore, secondTeam, secondTeamScore)
	l.wins[winner]++
	l.wins[loser] = l.wins[loser]
}

func winnerTeam(firstTeam TeamName, firstTeamScore int, secondTeam TeamName, secondTeamScore int) (TeamName, TeamName) {
	if firstTeamWins := firstTeamScore > secondTeamScore; firstTeamWins {
		return firstTeam, secondTeam
	}
	return secondTeam, firstTeam
}

type Ranker interface {
	Ranking() []string
}

func (l League) Ranking() []string {
	ranking := make([]string, 0, len(l.wins))
	for k := range l.wins {
		ranking = append(ranking, string(k))
	}

	sort.Slice(ranking, func(i, j int) bool {
		return l.wins[TeamName(ranking[i])] > l.wins[TeamName(ranking[j])]
	})

	return ranking
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	ranking := ranker.Ranking()

	joined := strings.Join(ranking, "\n")

	if _, err := writer.Write([]byte(joined)); err != nil {
		panic(err)
	}
}
