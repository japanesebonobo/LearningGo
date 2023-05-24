package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain", "タコの足とイカの足", "antholopologist", "タコの足は8本でイカの足は10本"}

	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf(" 「%s」の文字列は%dで、短い単語だ。\n", word, size)
		case 5:
			fmt.Printf(" 「%s」の文字列は%dで、これはちょうど良い長さだ。\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("「%s」の文字数は%dで、とても良い。", word, size)
			if n := len(word); size < n {
				fmt.Printf("%dバイトもある！\n", n)
			} else {
				fmt.Println()
			}
		}

		words := []string{"hi", "salutations", "hello"}
		for _, word := range words {
			switch wordLen := len(word); {
			case wordLen < 5:
				fmt.Println(word, "は短い単語です")
			case wordLen > 10:
				fmt.Println(word, "は長すぎる単語です")
			default:
				fmt.Println(word, "はちょうど良い長さの単語です")
			}
		}

	loop:
		for i := 0; i < 10; i++ {
			switch {
			case i%2 == 0:
				fmt.Println(i, "：偶数")
			case i%3 == 0:
				fmt.Println(i, "：3で割り切れるが2では割り切れない")
			case i%7 == 0:
				fmt.Println(i, "：ループ終了")
				break loop
			default:
				fmt.Println(i, "：退屈な数")
			}
		}
	}
}
