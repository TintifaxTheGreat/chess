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

	bbWhiteKing := uint64(binary.BigEndian.Uint64(data[:8]))
	bbWhiteQueen := uint64(binary.BigEndian.Uint64(data[8:16]))
	bbWhiteRook := uint64(binary.BigEndian.Uint64(data[16:24]))
	bbWhiteBishop := uint64(binary.BigEndian.Uint64(data[24:32]))
	bbWhiteKnight := uint64(binary.BigEndian.Uint64(data[32:40]))
	bbWhitePawn := uint64(binary.BigEndian.Uint64(data[40:48]))
	bbBlackKing := uint64(binary.BigEndian.Uint64(data[48:56]))
	bbBlackQueen := uint64(binary.BigEndian.Uint64(data[56:64]))
	bbBlackRook := uint64(binary.BigEndian.Uint64(data[64:72]))
	bbBlackBishop := uint64(binary.BigEndian.Uint64(data[72:80]))
	bbBlackKnight := uint64(binary.BigEndian.Uint64(data[80:88]))
	bbBlackPawn := uint64(binary.BigEndian.Uint64(data[88:96]))

	whitePawnsCenter := bbWhitePawn & CENTER
	blackPawnsCenter := bbBlackPawn & CENTER

	value = value + bits.OnesCount64(whitePawnsCenter)
	value = value - bits.OnesCount64(blackPawnsCenter)

	value = value + bits.OnesCount64(bbWhitePawn)*20
	value = value - bits.OnesCount64(bbBlackPawn)*20

	value = value + bits.OnesCount64(bbWhiteKnight)*60
	value = value - bits.OnesCount64(bbBlackKnight)*60

	value = value + bits.OnesCount64(bbWhiteBishop)*61
	value = value - bits.OnesCount64(bbBlackBishop)*61

	value = value + bits.OnesCount64(bbWhiteRook)*100
	value = value - bits.OnesCount64(bbBlackRook)*100

	value = value + bits.OnesCount64(bbWhiteQueen)*180
	value = value - bits.OnesCount64(bbBlackQueen)*180

	// TODO only use in early stage of game
	value = value + bits.OnesCount64(bbWhiteKnight&CENTER)
	value = value - bits.OnesCount64(bbBlackKnight&CENTER)

	value = value - bits.OnesCount64(bbWhiteQueen&CENTER)
	value = value + bits.OnesCount64(bbBlackQueen&CENTER)

	value = value - bits.OnesCount64(bbWhiteKnight&BASE_LINE)
	value = value + bits.OnesCount64(bbBlackKnight&BASE_LINE)
	value = value - bits.OnesCount64(bbWhiteBishop&BASE_LINE)
	value = value + bits.OnesCount64(bbBlackBishop&BASE_LINE)

	value = value - bits.OnesCount64(bbWhiteQueen&CENTER)
	value = value + bits.OnesCount64(bbBlackQueen&CENTER)

	value = value + bits.OnesCount64(bbWhiteKing&SAFE_KING)*5
	value = value - bits.OnesCount64(bbBlackKing&SAFE_KING)*5

	value = value + bits.OnesCount64(bbWhiteBishop&GOOD_BISHOP)
	value = value - bits.OnesCount64(bbBlackBishop&GOOD_BISHOP)

	//bbDefendingKing := bbWhiteKing
	if pos.Turn() == chess.Black {
		value *= (-1)
		//bbDefendingKing = bbBlackKing
	}

	//TODO only use this in the endgame
	/*
		if value < 0 {
			value += distance(bbWhiteKing, bbBlackKing)

			value += bits.OnesCount64(bbDefendingKing&CENT_0) * 3
			value += bits.OnesCount64(bbDefendingKing&CENT_1) * 2
			value += bits.OnesCount64(bbDefendingKing&BORD_1) * 1

		} else {
			value -= bits.OnesCount64(bbDefendingKing&CENT_0) * 3
			value -= bits.OnesCount64(bbDefendingKing&CENT_1) * 2
			value -= bits.OnesCount64(bbDefendingKing&BORD_1) * 1
		}
	*/

	return value
}

func negamax(pos *chess.Position, depth int, alpha int, beta int) int {
	outcome := pos.Status()
	if outcome != chess.NoMethod {
		switch outcome {
		case chess.Checkmate:
			alpha = -20000 - depth
		default:
			return 0
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

func distance(x uint64, y uint64) int {
	xLz := bits.LeadingZeros64(x)
	yLz := bits.LeadingZeros64(y)
	fx := xLz % 8
	fy := yLz % 8
	rx := xLz / 8
	ry := yLz / 8

	fD := fy - fx
	if fD < 0 {
		fD = -fD
	}

	rD := ry - rx
	if rD < 0 {
		rD = -rD
	}

	if rD < fD {
		return fD
	}

	return rD
}
