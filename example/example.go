//go:generate curry -file-suffix=_curried.go -function-modifier=C

package main

import "fmt"

// The generator will generate a curried version of the function:
//
// func appendC(seperator string) func(string) string {
// 	return func(str string) string {
// 		return append(seperator, str)
// 	}
// }
func append(seperator, str string) string {
	return str + seperator
}

type Superhero struct {
	Name string
}

// The generator will generate a curried version of the function:
//
// func (s Superhero) AbilityC(against Superhero) func(bool) string {
// 	return func(destructive bool) string {
// 		return s.Ability(against, destructive)
// 	}
// }
func (s Superhero) Ability(against Superhero, destructive bool) string {
	w := "wins"
	if !destructive {
		w = "loses"
	}
	return fmt.Sprintf("%s uses ability agains %s and %s the fight.",
		s.Name, against.Name, w,
	)
}

func main() {
	sentences := []string{
		"Programmer",
		"A person who fixed a problem",
		"That you don't know you have",
		"In a way you don't understand",
	}
	appendNewline := appendC(".\n")

	var result string
	for _, sentence := range sentences {
		result += appendNewline(sentence)
	}
	fmt.Println(result)

	batman := Superhero{Name: "Batman"}
	superman := Superhero{Name: "Superman"}
	batmanVsSuperman := batman.AbilityC(superman)
	fmt.Println(batmanVsSuperman(true))
	fmt.Println(batmanVsSuperman(false))
}
