// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"github.com/notnil/chess"
	"math/rand"
	"time"
)

const MAX_INT int = 32767
const MIN_INT int = -32768

const CENTER uint64 = 0x00003C3C3C3C0000

type AgentCPU struct {
	Agent
	depth int
}

func NewAgentCPU() *AgentCPU {
	rand.Seed(time.Now().Unix())
	return &AgentCPU{
		depth: 4,
	}
}

func (a *AgentCPU) MakeMove(game *chess.Game) *chess.Move {
	moves := game.ValidMoves()

	maxValue := MIN_INT
	maxIndex := MIN_INT

	pos := game.Position()
	var newPos *chess.Position
	for i, move := range moves {
		newPos = pos.Update(move)
		if value := -negamax(newPos, a.depth, MIN_INT, MAX_INT); value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}

	if maxIndex == MIN_INT {
		return nil //TODO remove this
	}

	return moves[maxIndex]
}

func (a *AgentCPU) GetChannel() chan *chess.Move {
	return nil
}

func (a *AgentCPU) Stop() {
	return
}
