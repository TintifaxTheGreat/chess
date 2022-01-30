// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"github.com/notnil/chess"
	"log"
	"math/rand"
	"sort"
	"time"
)

const MAX_INT int = 1000000
const MIN_INT int = -1000000

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
	if len(moves) == 1 {
		return moves[0]
	}
	alpha := MIN_INT
	beta := MAX_INT
	index := MIN_INT

	priorValues := make(map[*chess.Move]int)
	newPositions := make(map[*chess.Move]*chess.Position)
	for _, move := range moves {
		newPositions[move] = game.Position().Update(move)
		priorValues[move] = -evaluate(newPositions[move])
	}

	keys := make([]*chess.Move, 0, len(priorValues))
	for key := range priorValues {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return priorValues[keys[i]] > priorValues[keys[j]] })

	for i, move := range keys {
		value := -negamax(newPositions[move], a.depth, -beta, -alpha, false)
		log.Print("am back")
		if value >= beta {
			index = i
			log.Print("in break")
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
