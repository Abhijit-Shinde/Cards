package main

import "fmt"

func main() {
	zeroValue := make(map[string]int)
	colors := map[string]string{"red": "#eb343d", "green": "#eb343d", "yellow": "#ebe834"}

	colors["red"] = "34"

	for a,b := range colors{
		fmt.Printf("%v Color hash value is %v\n",a,b)
	}
	delete(colors,"yellow")

	fmt.Println(colors)
	fmt.Print(zeroValue)
}
