package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [filename]")
		return
	}
	farm := NewAntFarm()
	content, err := farm.ParseInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(content)
	fmt.Println()
	farm.EdmondsKarp()
	moves := farm.SimulateAnts()
	for _, move := range moves {
		fmt.Println(move)
	}
}
