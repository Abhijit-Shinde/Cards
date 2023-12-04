package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// e := englishBot{}
	// s := spanishBot{}

	printGreeting(englishBot{})
	printGreeting(spanishBot{})
}

//w/o interface
func (eb englishBot) getGreeting() string {
	return "Hello!"
} 

//w/o interface
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// //w/o interface
// func printGreetingForEnglish (eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// //w/o interface
// func printGreetingForSpanish (sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
