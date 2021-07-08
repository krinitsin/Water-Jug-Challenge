package solver

import (
	"errors"
	"fmt"
	"waterjugserver/internal/jug"
)

type Step struct {
	X int `json:"X"`
	Y int `json:"Y"`
	Explanation string `json:"Explanation"`
}

var errProblemCantBeSolved = errors.New("the problem can't be solved")

func Solve(x int, y int, z int) ([]Step, error) {
	if !canMeasureWater(x, y, z) {
		return nil, errProblemCantBeSolved
	}

	res1 := newSolver(x, y, z).solve()
	res2 := newSolver(y, x, z).solve()

	if len(res1) < len(res2) {
		return res1, nil
	}
	return res2, nil
}

func newSolver(x, y, z int) solver {
	return solver{
		from:   jug.New(x,"X"),
		to:     jug.New(y,"Y"),
		target: z,
		result: make([]Step,0),
	}
}

func (s solver) solve() []Step {
	// fill first jug
	s.fill(s.from)

	for !s.done() {
		s.transfer(s.from, s.to)

		if s.done() {
			break
		}

		// If first jug becomes empty, fill it
		if s.from.CurAmount() == 0 {
			s.fill(s.from)
		}

		// If second jug becomes full, empty it
		if s.to.CurAmount() == s.to.Size() {
			s.empty(s.to)
		}
	}

	return s.result
}

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}

	if z == 0 {
		return true
	}

	if x == 0 {
		return y == z
	}

	if y == 0 {
		return x == z
	}

	a := gcd(x, y)

	return z%a == 0

}

func gcd(a int, b int) int {
	if a > b {
		return gcd(a-b, b)
	} else if a < b {
		return gcd(b-a, a)
	} else {
		return a
	}
}

type solver struct {
	from   jug.Jug
	to     jug.Jug
	target int
	result []Step
}

func (s *solver) transfer(from, to jug.Jug) {
	jug.Transfer(from, to)
	s.appendResults(fmt.Sprintf("transfer water from: %s, to: %s",from.Name(),to.Name()))
}

func (s *solver) fill(j jug.Jug) {
	j.Fill()
	s.appendResults(fmt.Sprintf("fill bucket: %s",j.Name()))
}

func (s *solver) empty(j jug.Jug) {
	j.Empty()
	s.appendResults(fmt.Sprintf("empty bucket: %s",j.Name()))
}

func (s *solver) done() bool {
	return s.from.CurAmount() == s.target || s.to.CurAmount() == s.target
}

func (s *solver) appendResults(explanation string) {
	s.result = append(s.result, Step{X: s.from.CurAmount(), Y: s.to.CurAmount(),Explanation: explanation})
}
