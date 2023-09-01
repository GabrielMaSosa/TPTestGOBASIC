package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	catC = "Categoria C"
	catB = "Categoria B"
	catA = "Categoria A"
)

func TestCalcSalario(t *testing.T) {
	t.Run("CASO para Categoria C", func(t *testing.T) {
		//Arrange

		time := 500
		valespC := (float64(time) / 60.0) * 1000.0
		//Act
		resC, _ := CalcSalario(catC, time)
		//	msgerr := errors.New("NO se puede ingresar minutos negativos")
		//Assert

		assert.Equal(t, valespC, resC, "Error fallo el test")

	})
	t.Run("CASO  error para Categoria C", func(t *testing.T) {
		//Arrange
		time := -500
		//	valespC := (float64(time) / 60.0) * 1000.0
		//Act
		_, err1 := CalcSalario(catC, time)
		msgerr := errors.New("NO se puede ingresar minutos negativos")
		//Assert

		assert.Equal(t, msgerr, err1, "Error fallo el test")

	})

	t.Run("CASO para Categoria B", func(t *testing.T) {
		//Arrange

		time := 500
		valespB := (float64(time) / 60.0) * 1500.0 * 1.2
		//Act
		resB, _ := CalcSalario(catB, time)

		//Assert

		assert.Equal(t, valespB, resB, "Error fallo el test")

	})

	t.Run("CASO para errores de Categoria B", func(t *testing.T) {
		//Arrange

		time := -500
		//valespB := (float64(time) / 60.0) * 1500.0 * 1.2
		//Act
		_, err2 := CalcSalario(catB, time)
		msgerr := errors.New("NO se puede ingresar minutos negativos")
		//Assert

		assert.Equal(t, msgerr, err2, "Error fallo el test")

	})

	t.Run("CASO para Categoria A", func(t *testing.T) {
		//Arrange

		time := 500
		valespA := (float64(time) / 60.0) * 3000.0 * 1.5
		//Act
		resA, _ := CalcSalario(catA, time)
		//msgerr := errors.New("NO se puede ingresar minutos negativos")
		//Assert

		assert.Equal(t, valespA, resA, "Error fallo el test")

	})

	t.Run("CASO para Categoria  con errror A", func(t *testing.T) {
		//Arrange

		time := -500
		//		valespA := (float64(time) / 60.0) * 3000.0 * 1.5
		//Act
		_, err3 := CalcSalario(catA, time)
		msgerr := errors.New("NO se puede ingresar minutos negativos")
		//Assert

		assert.Equal(t, msgerr, err3, "Error fallo el test")

	})

}
