package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func main() {
	fmt.Println("Starting the program...")
	now := time.Now().Format("2006年01月02日 15時04分05秒")
	fmt.Println(now)
	emoji()
}

func emoji() {
	msg := "Ho🌮あA"
	println(msg, "(", utf8.RuneCountInString(msg), " chars ")
	for i, cc := range msg {
		fmt.Printf("code boundary:%d .. %c\n", i, cc)
	}
	println(len([]rune(msg)), len(msg))
	rca := []rune(msg)
	for i, cc := range rca {
		fmt.Printf("(%d) .. %c\n", i, cc)
	}
	ss := []rune(rca[2:4])
	fmt.Printf("%c %c %c\n", ss[0], ss, rca[2:4]) // slice of runes
}