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
