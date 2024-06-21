package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println(key, value)
	}
}
func changeVal(cityMap map[string]string) {
	cityMap["Bei"] = "aaa"
}

func main() {
	cityMap := make(map[string]string)
	cityMap["Bei"] = "bei"
	cityMap["Shang"] = "shang"
	printMap(cityMap)
	delete(cityMap, "Shang")

	printMap(cityMap)

	changeVal(cityMap)

	printMap(cityMap)
}
