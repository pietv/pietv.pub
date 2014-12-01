package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(new(big.Int).MulRange(1, 5))
}
