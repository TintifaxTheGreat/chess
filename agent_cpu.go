// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"encoding/binary"
	"github.com/notnil/chess"
	"math/bits"
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
		if value := a.evaluate(newPos); value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}
	return moves[maxIndex]
}

func (a *AgentCPU) GetChannel() chan *chess.Move {
	return nil
}

func (a *AgentCPU) Stop() {
	return
}

func (a *AgentCPU) evaluate(pos *chess.Position) int {
	// evaluation takes places from the perspective of the player whose turn it is

	/*outcome := game.Outcome()
	if outcome != chess.NoOutcome {
		switch outcome {
		case chess.WhiteWon:
			return MAX_INT
		case chess.BlackWon:
			return MIN_INT
		default:
			return 0
		}
	}

	*/
	var value int = 0
	data, _ := pos.Board().MarshalBinary() //TODO make this more efficient

	//bbWhiteKing := uint64(binary.BigEndian.Uint64(data[:8]))
	bbWhiteQueen := uint64(binary.BigEndian.Uint64(data[8:16]))
	bbWhiteRook := uint64(binary.BigEndian.Uint64(data[16:24]))
	bbWhiteBishop := uint64(binary.BigEndian.Uint64(data[24:32]))
	bbWhiteKnight := uint64(binary.BigEndian.Uint64(data[32:40]))
	bbWhitePawn := uint64(binary.BigEndian.Uint64(data[40:48]))
	//bbBlackKing := uint64(binary.BigEndian.Uint64(data[48:56]))
	bbBlackQueen := uint64(binary.BigEndian.Uint64(data[56:64]))
	bbBlackRook := uint64(binary.BigEndian.Uint64(data[64:72]))
	bbBlackBishop := uint64(binary.BigEndian.Uint64(data[72:80]))
	bbBlackKnight := uint64(binary.BigEndian.Uint64(data[80:88]))
	bbBlackPawn := uint64(binary.BigEndian.Uint64(data[88:96]))

	whitePawnsCenter := bbWhitePawn & CENTER
	blackPawnsCenter := bbBlackPawn & CENTER

	value += bits.OnesCount64(whitePawnsCenter)
	value -= bits.OnesCount64(blackPawnsCenter)

	value += bits.OnesCount64(bbWhitePawn) * 10
	value -= bits.OnesCount64(bbBlackPawn) * 10

	value += bits.OnesCount64(bbWhiteKnight) * 30
	value -= bits.OnesCount64(bbBlackKnight) * 30

	value += bits.OnesCount64(bbWhiteBishop) * 30
	value -= bits.OnesCount64(bbBlackBishop) * 30

	value += bits.OnesCount64(bbWhiteRook) * 50
	value -= bits.OnesCount64(bbBlackRook) * 50

	value += bits.OnesCount64(bbWhiteQueen) * 90
	value -= bits.OnesCount64(bbBlackQueen) * 90

	if pos.Turn() == chess.White {
		value *= (-1)
	}

	return value
}
