package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	//"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/samber/lo"
)

func main() {
	fmt.Println("Starting the program...")
	now := time.Now().Format("2006年01月02日 15時04分05秒")
	fmt.Println(now)
	emoji()
	misc()
	file_read()

	ret := input("Enter your name > ")
	fmt.Printf("your name is %s\n", ret)
}

func file_io() {
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func input(prompt string) string {
	print(prompt)
	buf, err := bufio.NewReader(os.Stdin).ReadBytes('\n') // 读取一行，包括换行符
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func clearString(str string) string {
    return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func regexp_01() {
    str := "Test@%String本田🌮#321gosamples.dev ـا ą ٦"
    fmt.Println(clearString(str))
		
		text := "GOSAMPLES.dev is  \t \r\n the best Golang \t\t\t\t    website in the \n world!"
    re := regexp.MustCompile(`\s+`)
    res := re.ReplaceAllString(text, " ")
    println(text,"\n--->>\n", res)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
    sb := strings.Builder{}
    sb.Grow(n)
    for i := 0; i < n; i++ {
        sb.WriteByte(charset[rand.Intn(len(charset))])
    }
    return sb.String()
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
func misc(){
	matching := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
    if strings.HasSuffix(x, "pu") {
        return "xpu", true
    }
    return "", false
})
println(matching)
// []string{"xpu", "xpu"}
}
func file_read(){

}