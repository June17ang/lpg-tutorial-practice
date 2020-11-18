package main

import "fmt"

func main() {
	cards := map[string]string{"one": "No1", "two": "No2", "three": "No3"}

	delete(cards, "one")
	fmt.Println(cards)

	// tried to use Anonymous functions failed// todo
	// printMap := func(Cards map[string]string) string {
	// 	str := ""
	// 	for key, value := range Cards {
	// 		str = str + fmt.Sprintf("Array of index %s is %s . \n", key, value)
	// 	}
	// 	return str
	// }

	//fmt.Printf("%v", printMap)
	printMap(cards)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
