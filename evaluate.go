// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"encoding/binary"
	"github.com/notnil/chess"
	"math/bits"
)

func evaluate(pos *chess.Position) int {
	// evaluation takes places from the perspective of the player who's turn it is
	// TODO consider terminal positions

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

	value = value + bits.OnesCount64(whitePawnsCenter)
	value = value - bits.OnesCount64(blackPawnsCenter)

	value = value + bits.OnesCount64(bbWhitePawn)*10
	value = value - bits.OnesCount64(bbBlackPawn)*10

	value = value + bits.OnesCount64(bbWhiteKnight)*30
	value = value - bits.OnesCount64(bbBlackKnight)*30

	value = value + bits.OnesCount64(bbWhiteBishop)*30
	value = value - bits.OnesCount64(bbBlackBishop)*30

	value = value + bits.OnesCount64(bbWhiteRook)*50
	value = value - bits.OnesCount64(bbBlackRook)*50

	value = value + bits.OnesCount64(bbWhiteQueen)*90
	value = value - bits.OnesCount64(bbBlackQueen)*90

	if pos.Turn() == chess.Black {
		value *= (-1)
	}

	return value
}

func negamax(pos *chess.Position, depth int) int {
	// TODO consider terminal positions
	if depth < 1 {
		foo := evaluate(pos) //TODO remove this
		return foo
	}
	maxValue := MIN_INT

	var newPos *chess.Position

	children := pos.ValidMoves()
	if depth == 2 {
		//log.Print(children)
	}
	for _, child := range children {
		newPos = pos.Update(child)
		value := -negamax(newPos, depth-1)
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
