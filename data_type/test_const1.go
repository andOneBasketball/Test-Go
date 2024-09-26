package main

import (
	"fmt"
	"math/big"
)

type Weekday int

func main() {
	const (
		SUNDAY Weekday = iota
		MONDAY
		TUESDAY
	)

	const (
		a = 1
		b
		c = 2
		d
		f
	)

	const (
		_ = 1 << (10*iota)
		KB
		MB
		GB
		TB
		PB
		EB
	)
	var ZB, YB big.Int
	ZB.Exp(big.NewInt(2), big.NewInt(70), nil)
	YB.Exp(big.NewInt(2), big.NewInt(80), nil)

	fmt.Println(SUNDAY, MONDAY, b, d, f)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB.String(), YB.String())
}
