package chess

import (
	//"fmt"
	"github.com/itegel/4jirgee/dog"
	"testing"
)

func Test_CheckWinner(t *testing.T) {
	err := Init()
	if err != nil {
		t.Error("init failed!")
	}

	var winner dog.Color
	err, winner = CheckWinner()
	if err == nil && winner == dog.NONE {
		t.Log("init check success!")
	}

	board = [4][4]*dog.Dog{{nil, wDogB, nil, nil},
		{nil, nil, nil, wDogC},
		{nil, nil, nil, nil},
		{bDogA, nil, nil, nil}}

	err, winner = CheckWinner()
	if err == nil && winner == dog.WHITE {
		t.Log("white winner check success!")
	}

	board = [4][4]*dog.Dog{{nil, nil, nil, nil},
		{wDogD, nil, nil, nil},
		{nil, nil, bDogC, bDogB},
		{bDogA, nil, nil, nil}}

	err, winner = CheckWinner()
	if err == nil && winner == dog.BLACK {
		t.Log("black winner check success!")
	}

}
