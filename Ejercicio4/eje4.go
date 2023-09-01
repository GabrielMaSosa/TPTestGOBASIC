/*

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as
estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.
Ejemplo:

*/

package main

import (
	"errors"
	"fmt"
)

func Minim(notas ...float64) (float64, error) {
	vali := notas[0]
	for _, v := range notas {
		if v <= 0 {
			return 0.0, errors.New("Valor que no corresponde a nota")
		}

		if v <= vali {
			vali = v
		}
	}

	return vali, nil
}
func Maxim(notas ...float64) (float64, error) {
	vali := notas[0]
	for _, v := range notas {
		if v <= 0 {
			return 0.0, errors.New("Valor que no corresponde a nota")
		}

		if v >= vali {

			vali = v
		}
	}
	return vali, nil
}
func Average(notas ...float64) (float64, error) {
	vali := 0.0
	for _, v := range notas {
		if v <= 0 {
			return 0.0, errors.New("Valor que no corresponde a nota")
		}

		vali += v
	}
	return (vali / float64(len(notas))), nil
}
func Dummy(notas ...float64) (float64, error) {
	return 0.0, errors.New("ERROR parametro no nombrado")
}

func Operation(tipo string) func(notas ...float64) (float64, error) {

	switch tipo {
	case "minimum":
		return Minim
	case "maximum":
		return Maxim
	case "average":
		return Average
	default:
		return Dummy
	}

}

func main() {
	const (
		min = "minimum"
		max = "maximum"
		avg = "average"
	)
	mini := Operation(min)
	maxi := Operation(max)
	aveg := Operation(avg)

	res1, err1 := mini(-1, 2, 3)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	res2, err2 := maxi(-1, 2, 3)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(res2)
	}
	res3, err3 := aveg(-1, 2, 3)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(res3)
	}

}
