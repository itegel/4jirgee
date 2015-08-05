package chess

import (
	"errors"
	"github.com/itegel/4jirgee/dog"
)

var (
	aliveDogs []dog.Dog
	deadDogs  []dog.Dog
	board     [4][4]*dog.Dog
	wDogA     *dog.Dog
	wDogB     *dog.Dog
	wDogC     *dog.Dog
	wDogD     *dog.Dog
	bDogA     *dog.Dog
	bDogB     *dog.Dog
	bDogC     *dog.Dog
	bDogD     *dog.Dog
)

func Init() error {

	wDogA = &dog.Dog{dog.WHITE, dog.ALIVE, dog.Position{0, 0}}
	wDogB = &dog.Dog{dog.WHITE, dog.ALIVE, dog.Position{1, 0}}
	wDogC = &dog.Dog{dog.WHITE, dog.ALIVE, dog.Position{2, 0}}
	wDogD = &dog.Dog{dog.WHITE, dog.ALIVE, dog.Position{3, 0}}
	bDogA = &dog.Dog{dog.BLACK, dog.ALIVE, dog.Position{0, 3}}
	bDogB = &dog.Dog{dog.BLACK, dog.ALIVE, dog.Position{1, 3}}
	bDogC = &dog.Dog{dog.BLACK, dog.ALIVE, dog.Position{2, 3}}
	bDogD = &dog.Dog{dog.BLACK, dog.ALIVE, dog.Position{3, 3}}

	board = [4][4]*dog.Dog{{wDogA, wDogB, wDogC, wDogD},
		{nil, nil, nil, nil},
		{nil, nil, nil, nil},
		{bDogA, bDogB, bDogC, bDogD}}

	return nil
}

//if only on dog left, will lose the game
// returns the color who wins, nil for no winner
func CheckWinner() (error, dog.Color) {
	wCount := 0
	bCount := 0
	var winner dog.Color
	winner = dog.NONE

	for _, tmprow := range board {
		for _, tmpdog := range tmprow {
			if tmpdog == nil {
				continue
			}

			if tmpdog.DStatus == dog.ALIVE {
				if tmpdog.DColor == dog.WHITE {
					wCount++
				} else if tmpdog.DColor == dog.BLACK {
					bCount++
				}
			} else {
				return errors.New("there cannot be dead dog in alive list!"), winner
			}
		}
	}

	if wCount <= 1 && bCount >= 2 {
		winner = dog.BLACK
	} else if bCount <= 1 {
		winner = dog.WHITE
	}

	return nil, winner
}

//traverse aliveDogs and find out dead dogs and switch it to deadDogs
func CleanDeadDogs() error {
	return nil
}
