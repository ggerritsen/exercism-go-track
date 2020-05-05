// Package tournament provides functionality around a table for a tournament
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type score struct {
	firstTeam, secondTeam, outcome string
}

func parse(s string) (*score, error) {
	ss := strings.Split(s, ";")
	if len(ss) != 3 {
		return nil, fmt.Errorf("invalid score %q", s)
	}

	score := &score{
		firstTeam:  ss[0],
		secondTeam: ss[1],
		outcome:    ss[2],
	}

	if score.firstTeam == score.secondTeam {
		return nil, fmt.Errorf("invalid score %q", s)
	}

	if score.outcome != "win" &&
		score.outcome != "draw" &&
		score.outcome != "loss" {
		return nil, fmt.Errorf("invalid score %q", s)
	}

	return score, nil
}

type teamScore struct {
	name                                       string
	matchesPlayed, wins, draws, losses, points int
}

type tournament struct {
	internal map[string]*teamScore
}

func (t *tournament) add(s score) {
	if _, ok := t.internal[s.firstTeam]; !ok {
		t.internal[s.firstTeam] = &teamScore{}
	}
	if _, ok := t.internal[s.secondTeam]; !ok {
		t.internal[s.secondTeam] = &teamScore{}
	}

	t.internal[s.firstTeam].matchesPlayed++
	t.internal[s.secondTeam].matchesPlayed++

	switch s.outcome {
	case "win":
		t.internal[s.firstTeam].wins++
		t.internal[s.firstTeam].points += 3
		t.internal[s.secondTeam].losses++
	case "draw":
		t.internal[s.firstTeam].draws++
		t.internal[s.firstTeam].points++
		t.internal[s.secondTeam].draws++
		t.internal[s.secondTeam].points++
	case "loss":
		t.internal[s.firstTeam].losses++
		t.internal[s.secondTeam].wins++
		t.internal[s.secondTeam].points += 3
	}
}

func (t *tournament) String() string {
	var teamScores []*teamScore
	for k, v := range t.internal {
		v.name = k
		teamScores = append(teamScores, v)
	}

	// sort in descending order
	sort.Slice(teamScores, func(i, j int) bool {
		if teamScores[i].points != teamScores[j].points {
			return teamScores[i].points > teamScores[j].points
		}
		return teamScores[i].name < teamScores[j].name
	})

	var b strings.Builder
	for _, s := range teamScores {
		b.WriteString(fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", s.name, s.matchesPlayed, s.wins, s.draws, s.losses, s.points))
	}

	return b.String()
}

// Tally reads all scores from r and writes the final tournament table to w.
// Returns an error for invalid input.
func Tally(r io.Reader, w io.Writer) error {
	var scores []score
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		text := strings.TrimSpace(sc.Text())
		if len(text) == 0 {
			continue
		}
		if text[0] == '#' {
			continue // ignore comment lines
		}
		s, err := parse(text)
		if err != nil {
			return err
		}
		scores = append(scores, *s)
	}

	t := &tournament{internal: map[string]*teamScore{}}
	for _, s := range scores {
		t.add(s)
	}

	_, err := w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n" + t.String()))
	return err
}
