package utilBasic

import "fmt"

func main() {
	//For Condicional
	fmt.Println("Iterador con For")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	//For While
	fmt.Println("Iterador con While")
	counter := 0
	for counter < 10 {
		fmt.Printf("%d ", counter)
		counter++
	}
	//For Forever
	counterForever := 0
	fmt.Printf("\nCounter Forever")
	for {
		fmt.Printf("%d ", counterForever)
		counterForever++
		if counterForever == 30 {
			break
		}
	}
}
