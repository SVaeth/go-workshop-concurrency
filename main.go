package main

import "fmt"

func main() {
	x := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	y := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	out := ""

	for _, vX := range x {
		for _, vY := range y {
			out += fmt.Sprintf("%s%s ", vX, vY)
		}
		out += "\n"
	}

	fmt.Println(out)
}
