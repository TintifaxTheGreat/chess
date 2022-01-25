package main

import (
	"github.com/notnil/chess"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeMove(t *testing.T) {
	agent := NewAgentCPU()
	game := chess.NewGame()

	agent.depth = 4
	load, err := chess.FEN("r6B/1p5p/3p2pK/1N5P/1p3P2/3p4/1R2p3/kB1b3Q w - - 0 1")
	if err != nil {
		return
	}
	load(game)
	move := agent.MakeMove(game)
	assert.Equal(t, "h1c6", move.String())

	agent.depth = 4
	load, err = chess.FEN("1rb4r/pkPp3p/1b1P3n/1Q6/N3Pp2/8/P1P3PP/7K w - - 1 1")
	if err != nil {
		return
	}
	load(game)
	move = agent.MakeMove(game)
	assert.Equal(t, "b5d5", move.String())

	agent.depth = 6
	load, err = chess.FEN("8/4p3/1B6/2N5/2k5/1R4K1/8/7B w - - 0 1")
	if err != nil {
		return
	}
	load(game)
	move = agent.MakeMove(game)
	assert.Equal(t, "g3g4", move.String())

	agent.depth = 4
	load, err = chess.FEN("8/2Q5/8/6q1/2K5/8/8/7k b - - 0 1")
	if err != nil {
		return
	}
	load(game)
	move = agent.MakeMove(game)
	assert.Equal(t, "g5c1", move.String())
}
