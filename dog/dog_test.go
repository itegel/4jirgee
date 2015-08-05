package dog

import (
	"fmt"
	"testing"
)

func Test_SetPosition(t *testing.T) {
	var dogA Dog

	_, err := dogA.SetPosition(Position{1, 2})
	if (err == nil) && (dogA.DPosition.X == 1) && (dogA.DPosition.Y == 2) {
		t.Log(" set pos normal success!")
	} else {
		t.Error(" set pos normal failed!")
	}

	_, err = dogA.SetPosition(Position{-1, -2})
	if err == nil {
		t.Error("invalid pos can not be set")
	} else {
		t.Log("set invalid pos test success")
	}
}

func Test_BeKilled(t *testing.T) {
	var dogA Dog
	//normal
	dogA.DStatus = ALIVE
	_, err := dogA.BeKilled()
	if err != nil {
		t.Error("kill dog failed!")
	} else {
		t.Log("Test_BeKilled normal success")
	}

	if dogA.DStatus != DEAD {
		t.Error("kill dog failed!")
	}

	//unnormal
	var dogB Dog
	_, err = dogB.BeKilled()
	if err == nil {
		t.Error("dead dog cannot be killed!")
	} else {
		t.Log("Test_BeKilled unnormal1 success")
	}

	dogB.DStatus = DEAD
	_, err = dogB.BeKilled()
	if err == nil {
		t.Error("dead dog cannot be killed!")
	} else {
		t.Log("Test_BeKilled unnormal2 success")
	}
}

func Test_Dogs(t *testing.T) {
	var dogs []Dog
	dogs = []Dog{Dog{WHITE, ALIVE, Position{0, 0}},
		Dog{WHITE, ALIVE, Position{1, 0}},
		Dog{WHITE, ALIVE, Position{2, 0}},
		Dog{WHITE, ALIVE, Position{3, 0}},
		Dog{BLACK, ALIVE, Position{0, 3}},
		Dog{BLACK, ALIVE, Position{1, 3}},
		Dog{BLACK, ALIVE, Position{2, 3}},
		Dog{BLACK, ALIVE, Position{3, 3}},
	}
	fmt.Println(dogs)

}
