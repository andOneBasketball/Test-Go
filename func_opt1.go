package main

import "fmt"

const (
	MultiRewardPropTurnedOff = iota
	MultiRewardPropTurnedOn
	MultiRewardPropTurningOn
)

func fun1() (a int, b int) {
	a = 10
	b = 20
	return
}

func main() {
	a, b := fun1()
	aa := int64(10)
	aa -= 1
	fmt.Println(a, b, MultiRewardPropTurnedOff, MultiRewardPropTurnedOn, MultiRewardPropTurningOn, aa)
}
