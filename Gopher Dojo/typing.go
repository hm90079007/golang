package main

import (
	"bufio"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"time"
)

func main() {
	strings := map[int]string{0: "hamabe", 1: "nanase", 2: "enako", 3: "momotuki", 4: "nitori", 5: "ikuta"}
	cnt := 0
	cntmax := 5
	start := time.Now()
	for cnt < cntmax {

		rand.Seed(time.Now().UnixNano())
		output := strings[rand.Intn(5)]
		fmt.Println(output)
		sc := bufio.NewScanner(os.Stdin)
		input := NextLine(sc)

		if output == input {
			fmt.Println(true)
			cnt++
		} else {
			fmt.Println(false)
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	secs := int((elapsed).Milliseconds())
	fmt.Println("time[ms]:", secs)

}

// NextLine is hogehoge
func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
