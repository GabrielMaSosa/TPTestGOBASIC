/*

Ejercicio5:





*/

package main

import "fmt"

func CantGato(data int) (float64, string) {

	return float64(data) * 5, "OK"

}

func CantPerro(data int) (float64, string) {

	return float64(data) * 10, "OK"

}
func CantHamster(data int) (float64, string) {

	return float64(data) * 0.25, "OK"

}
func CantTarantula(data int) (float64, string) {

	return float64(data) * 0.15, "OK"

}
func Dummy(data int) (float64, string) {

	return 0.0, "Animal no soportado"

}

func animal(tipo string) (func(int) (float64, string), string) {
	switch tipo {
	case "tarántulas":
		return CantTarantula, "OK"
	case "hamsters":
		return CantHamster, "OK"
	case "perros":
		return CantPerro, "OK"
	case "gatos":
		return CantGato, "OK"
	default:
		return Dummy, "No tenemos informacion al respecto de este tipo de animal"
	}

}

func main() {
	const (
		spi      = "tarántulas"
		rat      = "hamsters"
		dog      = "perros"
		cat      = "gatos"
		cantanim = 10
	)
	tipo := spi
	data, msg := animal(tipo)
	if msg == "OK" {
		cantani, msg2 := data(cantanim)
		if msg2 == "OK" {
			fmt.Println("La cantidad de alimento para ", cantanim, " ", tipo, "es", cantani)
		}

	} else {
		fmt.Println(msg)
	}

}
