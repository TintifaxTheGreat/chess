// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"encoding/binary"
	"github.com/notnil/chess"
	"math/bits"
)

func evaluate(pos *chess.Position) int {
	// evaluation takes places from the perspective of the player who's turn it is
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

func negamax(pos *chess.Position, depth int, alpha int, beta int) int {
	//TODO move ordering
	outcome := pos.Status()
	if outcome != chess.NoMethod {
		switch outcome {
		case chess.Checkmate:
			alpha = 20000 + depth
		default:
			return 0
		}
		if pos.Turn() == chess.Black {
			alpha *= (-1)
		}
		return alpha
	}

	if depth < 1 {
		return evaluate(pos)
	}
	var newPos *chess.Position

	children := pos.ValidMoves()

	for _, child := range children {
		newPos = pos.Update(child)
		value := -negamax(newPos, depth-1, -beta, -alpha)
		if value >= beta {
			return beta
		}
		if value > alpha {
			alpha = value
		}
	}
	return alpha
}
