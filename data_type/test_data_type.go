package main
import "fmt"
const boilingF = 212.0
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
	d := 1e100
	fmt.Printf("d = %e\n", d)
	x := 0x1234
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	s := "hello"
	fmt.Printf("%s %[1]s\n", s)
	newline := '\n'
	fmt.Printf("%d %[1]q %[1]c\n", newline)
}