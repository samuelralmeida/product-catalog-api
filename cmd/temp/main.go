package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	count := 1
	for {
		if count >= 8 {
			break
		}

		bg := big.NewInt(60 - 1)
		n, err := rand.Int(rand.Reader, bg)
		fmt.Println(err)
		fmt.Println(n)
		count++
	}

}
