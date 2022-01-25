// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"github.com/notnil/chess"
	"log"
	"math/rand"
	"time"
)

const MAX_INT int = 32767
const MIN_INT int = -32768

const CENTER uint64 = 0x00003C3C3C3C0000

type AgentCPU Agent

func NewAgentCPU() *AgentCPU {
	rand.Seed(time.Now().Unix())
	return &AgentCPU{}
}

func (a *AgentCPU) MakeMove(game *chess.Game) *chess.Move {
	moves := game.ValidMoves()
	//TODO consider outcome here
	maxValue := MIN_INT
	maxIndex := MIN_INT
	//values := make([]int, len(moves))
	pos := game.Position()
	var newPos *chess.Position
	for i, move := range moves {
		newPos = pos.Update(move)
		if value := -negamax(newPos, 2); value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}
	if maxIndex == MIN_INT {
		return nil //TODO remove this
	}
	log.Print("+++++++++++++++")
	log.Print(maxValue)
	log.Print(maxIndex)
	log.Print("+++++++++++++++")
	return moves[maxIndex]
}

func (a *AgentCPU) GetChannel() chan *chess.Move {
	return nil
}

func (a *AgentCPU) Stop() {
	return
}
