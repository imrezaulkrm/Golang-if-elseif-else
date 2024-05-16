package main

import "fmt"

func main() {
	var alphabet string
	fmt.Println("Check whether a character is alphabet or not")
	fmt.Print("Write Your Alphabet: ")
	fmt.Scanf("%v", &alphabet)

	switch alphabet {
	case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
		fmt.Printf("%v is Alphabet\n", alphabet)
	default:
		fmt.Printf("%v is Not Alphabet\n", alphabet)
	}
}

/*
func main() {
	var alphabet string
	fmt.Println("Check whether a character is Uppercase or Lowercase")
	fmt.Print("Write Your Alphabet: ")
	fmt.Scanf("%v", &alphabet)

	if alphabet >= "a" && alphabet <= "z" {
		fmt.Printf("%v is Lowercase Alphabet\n", alphabet)
	} else if alphabet >= "A" && alphabet <= "Z" {
		fmt.Printf("%v is Uppercase Alphabet\n", alphabet)
	} else {
		fmt.Printf("%v is Not Alphabet\n", alphabet)
	}
}
*/
