// multiple expression &&, ||, !
package main
import "fmt"
var (
	hasGas = true
	hasKeyIgnition = true
	hasBurger = true
	hasSandwich = true
)

func main (){
	if hasGas && hasKeyIgnition {
		fmt.Println("Can drive the car")
	} else {
		fmt.Println("Cannot drive the car")
	}
	if hasBurger || hasSandwich {
		fmt.Println("Can eat")
	} else {
		fmt.Println("Get something good to eat")
	}

}



