package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func palindromo(word string) bool {

	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")

	var length int = len(word)
	var isPalindrome bool = false

	for i := 0; i < len(word)/2; i++ {
		if word[i] == word[length-1-i] {
			isPalindrome = true
		} else {
			isPalindrome = false
			break
		}
	}

	return isPalindrome
}

func reversa(word string) string {
	arrayCadena := strings.Split(word, "")

	arrayInvertida := make([]string, 0)

	for i := 0; i < len(arrayCadena); i++ {
		arrayInvertida = append(arrayInvertida, arrayCadena[len(arrayCadena)-1-i])
	}
	arrayInvertidaString := strings.Join(arrayInvertida, "")
	return arrayInvertidaString
}

func main() {
	var palabra string
	fmt.Print("Escribe una palabra: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		palabra = scanner.Text()
	}

	result := palindromo(palabra)
	if result {
		fmt.Printf("La palabra %s es palindromo", palabra)
	} else {
		fmt.Printf("\nLa palabra %s no es palindromo", palabra)
	}

	resultado := reversa("pachitron")
	fmt.Printf("\nLa palabra %s invertida es %s", palabra, resultado)
}
