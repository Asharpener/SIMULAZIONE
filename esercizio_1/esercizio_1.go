package main

import (
	"fmt"
	"os"
	"unicode"
)

// ciao mondo : ciao = parola pari, mondo = parola dispari
//di ciao: c = lettera pari, i lettera dispari etc.
//in ciao c dovrà diventare maiuscola, in mondo m dovrà diventare minuscola
func main() {
	rigaCmd := os.Args[1:]
	for i := 0; i < len(rigaCmd); i++ {
		fmt.Print(TrasformaParola(rigaCmd[i], i))
		if i == len(rigaCmd)-1 {
			break
		} else {
			fmt.Print(" ")
		}
	}
}

func TrasformaParola(parola string, posizione int) (parolaTrasormata string) {

	if posizione%2 == 0 {
		//parola pari
		for i := 0; i < len(parola); i++ {
			if i%2 == 0 {
				//trasforma lettera in Upper
				parolaTrasormata = parolaTrasormata + string(unicode.ToUpper(rune(parola[i])))
			} else {
				//trasforma lettera in Lower
				parolaTrasormata = parolaTrasormata + string(unicode.ToLower(rune(parola[i])))
			}
		}
	} else {
		//parola dispari
		for i := 0; i < len(parola); i++ {
			if i%2 == 0 {
				//trasforma lettera in lower
				parolaTrasormata = parolaTrasormata + string(unicode.ToLower(rune(parola[i])))
			} else {
				// trasforma lettera in Upper
				parolaTrasormata = parolaTrasormata + string(unicode.ToUpper(rune(parola[i])))
			}
		}
	}
	return parolaTrasormata
}
