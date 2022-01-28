// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"github.com/notnil/chess"
	"math/rand"
	"sort"
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

	alpha := MIN_INT
	beta := MAX_INT
	index := MIN_INT

	pos := game.Position()
	var newPos *chess.Position
	var priorValues = make(map[*chess.Move]int)
	for _, move := range moves {
		newPos = pos.Update(move)
		priorValues[move] = -evaluate(newPos)
	}
	keys := make([]*chess.Move, 0, len(priorValues))
	for key := range priorValues {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return priorValues[keys[i]] > priorValues[keys[j]] })

	for i, move := range keys {
		newPos = pos.Update(move) //TODO use cached position from above
		value := -negamax(newPos, a.depth, -beta, -alpha)
		if value >= beta {
			index = i
			break
		}
		if value > alpha {
			alpha = value
			index = i
		}
	}

	if index == MIN_INT {
		return nil //TODO remove this
	}

	return keys[index]
}

func (a *AgentCPU) GetChannel() chan *chess.Move {
	return nil
}

func (a *AgentCPU) Stop() {
	return
}
