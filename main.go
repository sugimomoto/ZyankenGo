package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var myHand Hand
	var enemyHand Hand
	var input string

	fmt.Println("じゃんけんスタート！")

	for {
		fmt.Println("0,1,2のいずれかを入力してね")
		fmt.Println("0:グー　1:パー　2:チョキ")

		fmt.Scanf("%s", &input)

		val, err := checkInput(input)

		if err != nil {
			fmt.Println("入力された値が適切じゃないよ")
			continue
		}

		myHand.SetHand(val)
		enemyHand.SetHand(1)

	}
}

func checkInput(input string) (output int, err error) {
	val, err := strconv.Atoi(input)

	if err != nil {
		return 0, errors.New("Input is not integer")
	}

	if !(0 <= val && 2 >= val) {
		return 0, errors.New("Input is out of range for Hand")
	}

	return val, nil
}
