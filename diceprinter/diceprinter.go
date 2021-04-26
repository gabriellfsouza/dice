package diceprinter

import (
	"fmt"

	"github.com/gabriellfsouza/dice"
	"github.com/gabriellfsouza/dice/diceprinter"
)

func main() {

	fmt.Println(dice.Roll(6))
	diceprinter.PrintRoll(6)
}

func PrintRoll(n int64) {
	fmt.Println(dice.Roll(6))
}
