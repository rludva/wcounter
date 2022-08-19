package main

import (
	"bufio"
	"fmt"
	"github.com/rludva/wcounter/src/counter"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	//
	report := counter.Wcounter(text)
	for word, counter := range report {
		fmt.Printf("v[%v]:%v\n", word, counter)
	}
}
