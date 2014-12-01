package main

import (
	"fmt"
	"math/big"
)

func main() {
	n, _ := new(big.Int).SetString("2", 10)
	n.Mul(n, n)
	fmt.Println(n)
}
