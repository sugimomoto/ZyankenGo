package main

import "errors"

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
	Undefined
)

func (hand Hand) String() string {
	switch hand {
	case Rock:
		return "グー"
	case Paper:
		return "パー"
	case Scissors:
		return "チョキ"
	default:
		return "Undefined"
	}
}

func (hand Hand) SetHand(input int) (h Hand, err error) {
	switch input {
	case 0:
		return Rock, nil
	case 1:
		return Paper, nil
	case 2:
		return Scissors, nil
	default:
		return Undefined, errors.New("Undefined")
	}

}
