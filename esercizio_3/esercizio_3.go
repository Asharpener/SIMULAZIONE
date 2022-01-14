package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	ascissa   float64
	ordinata  float64
}

func main() {
	tragitto := NuovoTragitto()
	fmt.Println("Lunghezza percorso: ", Lunghezza(tragitto))
	fmt.Println("Punto oltre metà", String(PuntoOltreMetà(tragitto)))
}
func NuovoTragitto() (tragitto []Punto) {
	tragitto = make([]Punto, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ";")
		etichetta := split[0]
		ascissa, _ := strconv.ParseFloat(split[1], 64)
		ordinata, _ := strconv.ParseFloat(split[2], 54)
		valori := Punto{etichetta, ascissa, ordinata}
		tragitto = append(tragitto, valori)
	}
	return tragitto
}

func Distanza(p1, p2 Punto) float64 {
	distanza := math.Sqrt((p1.ascissa-p2.ascissa)*(p1.ascissa-p2.ascissa) + (p1.ordinata-p2.ordinata)*(p1.ordinata-p2.ordinata))
	return distanza
}

func String(p Punto) string {
	stringa := fmt.Sprintf("%s = (%f, %f)", p.etichetta, p.ascissa, p.ordinata)
	return stringa
}

func Lunghezza(tragitto []Punto) float64 {
	var risultato float64
	for i := 0; i < len(tragitto)-1; i++ {
		vardistanza := Distanza(tragitto[i], tragitto[i+1])
		risultato += vardistanza
	}
	return risultato
}

func PuntoOltreMetà(tragitto []Punto) (puntometà Punto) {
	metaLunghezza := Lunghezza(tragitto) / 2
	var risultato float64
	for i := 0; i < len(tragitto)-1; i++ {
		vardistanza := Distanza(tragitto[i], tragitto[i+1])
		risultato += vardistanza
		if risultato > metaLunghezza {
			return tragitto[i+1]
		}
	}
	fmt.Println("Errore")
	return tragitto[1]
}
