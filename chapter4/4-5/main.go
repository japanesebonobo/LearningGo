package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // シードの指定
	switch n := rand.Intn(10); {     // 0以上10未満の整数を戻す
	case n == 0:
		fmt.Println("少し小さすぎます:", n)
	case n > 5:
		fmt.Println("大きすぎます:", n)
	default:
		fmt.Println("いい感じの数字です:", n)
	}
}
