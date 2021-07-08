package jug

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJug_Fill(t *testing.T) {
	j := New(5,"j")
	j.Fill()
	require.Equal(t, 5, j.CurAmount())
}

func TestJug_Empty(t *testing.T) {
	j := New(5,"j")
	j.Fill()
	j.Empty()
	require.Equal(t, 0, j.CurAmount())
}

func TestTransfer(t *testing.T) {
	t.Run("enough space in second bucket", func(t *testing.T) {
		from := New(5,"from")
		from.Fill()
		to := New(5,"to")
		Transfer(from, to)

		require.Equal(t, 0, from.CurAmount())
		require.Equal(t, 5, to.CurAmount())
	})

	t.Run("enough space in second bucket", func(t *testing.T) {
		from := New(6,"from")
		from.Fill()
		to := New(5,"to")
		Transfer(from, to)

		require.Equal(t, 1, from.CurAmount())
		require.Equal(t, 5, to.CurAmount())
	})
}
