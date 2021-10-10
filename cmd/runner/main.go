package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/datosh/parker"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		var ns parker.NumberSquare
		for i := 0; i < 9; i++ {
			ns.Set(i, rand.Intn(99))
		}
		orig := ns
		ns.SquareValues()
		same := ns.SameValues()
		if same >= 4 {
			fmt.Printf("Square: %v, has %d same values\n", orig, same)
		}
	}
}
