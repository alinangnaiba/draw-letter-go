package main

import (
	"fmt"
	"go-practice/src/letters"
	"strings"
)

type choiceLetter string

const (
	letterO = "O"
	letterX = "X"
	letterY = "Y"
	letterZ = "Z"
)

func main() {
	choice, num := getUserInput()
	drawChoice(choice, num)
}

func getUserInput() (string, int) {
	choice := ""
	var num int
	fmt.Print("Enter character [O, X, Y, Z]: ")
	fmt.Scanf("%s", &choice)
	choice = strings.ToUpper(choice)
	for !validateInputString(choice) {
		fmt.Println("Invalid input. Try again.")
		fmt.Print("Enter character [O, X, Y, Z]: ")
		fmt.Scanf("%s", &choice)
		choice = strings.ToUpper(choice)
	}

	fmt.Print("\nEnter a non negative odd integer: ")
	fmt.Scanf("%d", &num)
	for !validateInputNumber(num) {
		fmt.Println("Invalid input. Try again.")
		fmt.Print("Enter a non negative odd integer: ")
		fmt.Scanf("%d", &num)
	}

	return choice, num
}

func drawChoice(choice string, size int) {
	switch choice {
	case letterO:
		letter := letters.O{Size: size}
		letter.Draw()
	case letterX:
		letter := letters.X{Size: size}
		letter.Draw()
	case letterY:
		letter := letters.Y{Size: size}
		letter.Draw()
	case letterZ:
		letter := letters.Z{Size: size}
		letter.Draw()
	default:
		fmt.Println("Invalid choice.")
	}
}

func validateInputNumber(number int) bool {
	if number%2 == 0 {
		return false
	}
	return true
}

func validateInputString(str string) bool {
	arr := [4]string{letterO, letterX, letterY, letterZ}
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
		}
		continue
	}
	return false
}
