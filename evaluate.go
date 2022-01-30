// Copyright © 2021-2022 Eugen Lindorfer, Austria

package main

import (
	"encoding/binary"
	"github.com/notnil/chess"
	"math/bits"
)

const CENTER uint64 = 0x00003C3C3C3C0000
const BORD_0 uint64 = 0xff818181818181ff
const BORD_1 uint64 = 0x007e424242427e00
const CENT_1 uint64 = 0x00003c24243c0000
const CENT_0 uint64 = 0x0000001818000000
const SAFE_KING uint64 = 0xc3000000000000c3
const GOOD_BISHOP uint64 = 0x42006666004200
const BASE_LINE uint64 = 0xff000000000000ff

func evaluate(pos *chess.Position) int {
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

	value += bits.OnesCount64(bbWhitePawn&CENT_0) * 4
	value -= bits.OnesCount64(bbBlackPawn&CENT_0) * 4

	value += bits.OnesCount64(bbWhitePawn&CENT_1) * 2
	value -= bits.OnesCount64(bbBlackPawn&CENT_1) * 2

	value += bits.OnesCount64(bbWhitePawn) * 20
	value -= bits.OnesCount64(bbBlackPawn) * 20

	value += bits.OnesCount64(bbWhiteKnight) * 60
	value -= bits.OnesCount64(bbBlackKnight) * 60

	value += bits.OnesCount64(bbWhiteBishop) * 61
	value -= bits.OnesCount64(bbBlackBishop) * 61

	value += bits.OnesCount64(bbWhiteRook) * 100
	value -= bits.OnesCount64(bbBlackRook) * 100

	value += bits.OnesCount64(bbWhiteQueen) * 180
	value -= bits.OnesCount64(bbBlackQueen) * 180

	// TODO only use in early stage of game
	value += bits.OnesCount64(bbWhiteKnight & CENTER)
	value -= bits.OnesCount64(bbBlackKnight & CENTER)

	value -= bits.OnesCount64(bbWhiteQueen & CENTER)
	value += bits.OnesCount64(bbBlackQueen & CENTER)

	value -= bits.OnesCount64(bbWhiteKnight&BASE_LINE) * 2
	value += bits.OnesCount64(bbBlackKnight&BASE_LINE) * 2
	value -= bits.OnesCount64(bbWhiteBishop&BASE_LINE) * 2
	value += bits.OnesCount64(bbBlackBishop&BASE_LINE) * 2

	value += bits.OnesCount64(bbWhiteKing&SAFE_KING) * 5
	value -= bits.OnesCount64(bbBlackKing&SAFE_KING) * 5

	value += bits.OnesCount64(bbWhiteBishop & GOOD_BISHOP)
	value -= bits.OnesCount64(bbBlackBishop & GOOD_BISHOP)

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

func negamax(pos *chess.Position, depth int, alpha int, beta int, quiescence bool) int {
	outcome := pos.Status()
	if outcome != chess.NoMethod {
		switch outcome {
		case chess.Checkmate:
			alpha = -40000 - depth
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
		value := 0
		newPos = pos.Update(child)
		if child.HasTag(chess.Capture) && !quiescence {
			value = -negamax(newPos, depth, -beta, -alpha, true)
		} else {
			value = -negamax(newPos, depth-1, -beta, -alpha, quiescence)
		}

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
