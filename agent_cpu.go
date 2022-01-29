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
const BORD_0 uint64 = 0xff818181818181ff
const BORD_1 uint64 = 0x007e424242427e00
const CENT_1 uint64 = 0x00003c24243c0000
const CENT_0 uint64 = 0x0000001818000000
const SAFE_KING uint64 = 0xc3000000000000c3
const GOOD_BISHOP uint64 = 0x42006666004200
const BASE_LINE uint64 = 0xff000000000000ff

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
		value := -negamax(newPositions[move], a.depth, -beta, -alpha)
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
