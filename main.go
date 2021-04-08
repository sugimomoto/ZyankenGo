package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var input string
	var endFlg string

	fmt.Println("じゃんけんスタート！")
	rand.Seed(time.Now().Unix())

	for {
		fmt.Println("0,1,2のいずれかを入力してね")
		fmt.Println("0:グー　1:チョキ　2:パー")

		fmt.Scanf("%s", &input)

		val, err := checkInput(input)

		if err != nil {
			fmt.Println("入力された値が適切じゃないよ")
			continue
		}
		fmt.Println("じゃーんけーん")

		myHand := GetHand(val)
		enemyHand := GetHand(rand.Intn(2))
		fmt.Println("ぽん！")

		fmt.Println("あなた ", myHand)
		fmt.Println("あいて ", enemyHand)

		switch myHand.Judge(enemyHand) {
		case Win:
			fmt.Println("あなたの勝ち！")
		case Loss:
			fmt.Println("あなたの負け！")
		case Draw:
			fmt.Println("引き分け！")
		}

		fmt.Println("終了する場合は end と入力してね")

		fmt.Scanf("%s", &endFlg)

		if endFlg == "end" {
			break
		}
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
