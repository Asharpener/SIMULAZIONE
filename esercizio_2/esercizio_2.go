package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]

	mappa := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := i + 2; j < len(s); j++ {
			if s[i] == s[j] {
				chiave := s[i : j+1]
				mappa[chiave]++
			}
		}
	}

	slice := make([]string, 0)

	//creo la slice
	//la riempio
	//la ordino
	//dooopo stampo i suoi valori

	for chiave, _ := range mappa {
		slice = append(slice, chiave)
	}
	for i := 0; i < len(slice); i++ {
		parolaMax := i
		for j := i + 1; j < len(slice); j++ {
			if len(slice[parolaMax]) < len(slice[j]) {
				parolaMax = j
			}
		} //in questo ciclo ho confrontato la parola i con tutte le successive ad essa
		//e in parolaMax ho salvato la posziione dellla parola più lunga
		dummy := slice[parolaMax]
		slice[parolaMax] = slice[i]
		slice[i] = dummy
	}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i],"-> Occorrenze: " mappa[slice[i]])
	}
}
