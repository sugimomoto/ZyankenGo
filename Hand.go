package main

type Hand int

const (
	Rock Hand = iota
	Scissors
	Paper
	Undefined
)

type Result int

const (
	Win Result = iota
	Loss
	Draw
)

func (hand Hand) String() string {
	switch hand {
	case Rock:
		return "グー"
	case Scissors:
		return "チョキ"
	case Paper:
		return "パー"
	default:
		return "Undefined"
	}
}

func GetHand(input int) Hand {
	switch input {
	case 0:
		return Rock
	case 1:
		return Scissors
	case 2:
		return Paper
	default:
		return Undefined
	}
}

func (myHand Hand) Judge(enemyHand Hand) Result {

	result := (myHand - enemyHand + 3) % 3

	switch result {
	case 0:
		return Draw
	case 1:
		return Loss
	case 2:
		return Win
	default:
		return Draw
	}
}
