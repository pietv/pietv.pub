package main

import (
	"math/big"
	"os"
	"text/template"
)

func factorial(n int) *big.Int {
	one := big.NewInt(1)

	if n == 0 || n == 1 {
		return one
	}

	f := one
	for i := 1; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}

const FactorialsTmpl = `var fstrings = []string{
{{range $index, $elem := .}}{{printf "\t/* %3d */ \"%s\",\n" $index $elem}}{{end}}}
`

func main() {
	f := make([]*big.Int, 256)
	for i := 0; i < len(f); i++ {
		f[i] = factorial(i)
	}
	template.Must(template.New("Factorial").Parse(FactorialsTmpl)).Execute(os.Stdout, f)
}
