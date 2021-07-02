package solver

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_solverFill(t *testing.T) {
	s := newSolver(5, 5, 5)

	s.fill(s.from)
	require.Equal(t, []Step{{5, 0}}, s.result)
	require.Equal(t, 5, s.from.CurAmount())
	require.Equal(t, 0, s.to.CurAmount())
}

func Test_solverEmpty(t *testing.T) {
	s := newSolver(5, 5, 5)

	s.fill(s.from)
	s.empty(s.from)
	require.Equal(t, []Step{{5, 0}, {0, 0}}, s.result)
	require.Equal(t, 0, s.from.CurAmount())
	require.Equal(t, 0, s.to.CurAmount())
}

func Test_solverTransfer(t *testing.T) {
	s := newSolver(5, 3, 5)

	s.fill(s.from)
	s.transfer(s.from, s.to)
	require.Equal(t, []Step{{5, 0}, {2, 3}}, s.result)
	require.Equal(t, 2, s.from.CurAmount())
	require.Equal(t, 3, s.to.CurAmount())
}

func Test_canMeasureWater(t *testing.T) {
	type in struct {
		x int
		y int
		z int
	}

	type out struct {
		solvable bool
	}

	for _, test := range []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "not enough space",
			in:   in{x: 3, y: 2, z: 19},
			out:  out{solvable: false},
		},
		{
			name: "target is 0",
			in:   in{x: 5, y: 5, z: 0},
			out:  out{solvable: true},
		},
		{
			name: "from capacity is 0",
			in:   in{x: 0, y: 2, z: 1},
			out:  out{solvable: false},
		},
		{
			name: "from capacity is 0",
			in:   in{x: 0, y: 2, z: 2},
			out:  out{solvable: true},
		},
		{
			name: "to capacity is 0",
			in:   in{x: 2, y: 0, z: 1},
			out:  out{solvable: false},
		},
		{
			name: "to capacity is 0",
			in:   in{x: 1, y: 0, z: 1},
			out:  out{solvable: true},
		},
		{
			name: "gcd is not a divider of target",
			in:   in{x: 6, y: 3, z: 4},
			out:  out{solvable: false},
		},
		{
			name: "gcd is a divider of target, from is larger",
			in:   in{x: 15, y: 3, z: 9},
			out:  out{solvable: true},
		},
		{
			name: "gcd is a divider of target, to is larger",
			in:   in{x: 3, y: 15, z: 9},
			out:  out{solvable: true},
		},
		{
			name: "gcd is a divider of target, from == false",
			in:   in{x: 4, y: 4, z: 4},
			out:  out{solvable: true},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.out.solvable, canMeasureWater(test.in.x, test.in.y, test.in.z))
		})
	}
}

func Test_Solve(t *testing.T) {
	type in struct {
		x int
		y int
		z int
	}

	type out struct {
		steps []Step
		err   error
	}

	for _, test := range []struct {
		name string
		in   in
		out  out
	}{
		{name: "solve from big to small ",
			in:  in{x: 10, y: 2, z: 4},
			out: out{steps: []Step{{2, 0}, {0, 2}, {2, 2}, {0, 4}}},
		},
		{name: "solve from small to big",
			in:  in{x: 2, y: 10, z: 4},
			out: out{steps: []Step{{2, 0}, {0, 2}, {2, 2}, {0, 4}}},
		},
		{name: "unsolvable",
			in:  in{x: 3, y: 3, z: 9},
			out: out{err: errProblemCantBeSolved},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			res, err := Solve(test.in.x, test.in.y, test.in.z)

			if test.out.err != nil {
				require.ErrorIs(t, err, test.out.err)
				return
			}

			t.Log("results", res)
			require.NoError(t, err)
			require.Equal(t, test.out.steps, res)
		})
	}
}

func Test_solve(t *testing.T) {
	type in struct {
		x int
		y int
		z int
	}

	type out struct {
		steps []Step
	}

	for _, test := range []struct {
		name string
		in   in
		out  out
	}{
		{name: "solve from big to small ",
			in:  in{x: 15, y: 3, z: 9},
			out: out{steps: []Step{{15, 0}, {12, 3}, {12, 0}, {9, 3}}},
		},
		{name: "solve from big to small ",
			in:  in{x: 3, y: 15, z: 9},
			out: out{steps: []Step{{3, 0}, {0, 3}, {3, 3}, {0, 6}, {3, 6}, {0, 9}}},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			s := newSolver(test.in.x, test.in.y, test.in.z)
			res := s.solve()

			t.Log("results", res)
			require.Equal(t, test.out.steps, res)
		})
	}
}
