// dogs, two dogs together can beat one dog on the other side
package dog

import (
	"errors"
)

const (
	WHITE = iota
	BLACK
	NONE
)

const (
	ALIVE = 1
	DEAD  = 2
)

type Color byte
type Status byte

type Position struct {
	X, Y int
}

type Dog struct {
	DColor    Color
	DStatus   Status
	DPosition Position
}

func (dog *Dog) SetColor(color Color) error {
	if (color < WHITE) || (color > BLACK) {
		return errors.New("Color Invalid")
	}
	dog.DColor = color
	return nil
}

func (dog *Dog) SetPosition(pos Position) (Dog, error) {
	if (pos.X < -1) || (pos.X > 3) || (pos.Y < -1) || (pos.Y > 3) {
		return *dog, errors.New("pos invalid")
	}
	dog.DPosition = pos
	return *dog, nil
}

func (dog *Dog) SetStatus(status Status) error {
	if status != ALIVE && status != DEAD {
		return errors.New("status invalid")
	}
	dog.DStatus = status
	return nil
}

func (dog *Dog) BeKilled() (Dog, error) {
	if dog.DStatus == ALIVE {
		dog.DStatus = DEAD
		return *dog, nil
	} else {
		return *dog, errors.New("DEAD dog can not be killed!")
	}
}
