package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i int

	rand.Seed(time.Now().UnixNano())
	n := 0

	for {
		n++

		i = rand.Intn(4200)
		fmt.Println(i)

		if i%42 == 0 {
			break
		}
	}

	fmt.Printf("\n**Exiting after %d iterations.\n\n", n)

loop:
	for i = 0; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)

		switch i {
		case 2, 3:
			fmt.Printf("Breaking switch, i == %d.\n", i)
			break
		case 5:
			fmt.Println("Breaking loop, i == 5.")
			break loop
		}
	}

	fmt.Println("End.")
}
