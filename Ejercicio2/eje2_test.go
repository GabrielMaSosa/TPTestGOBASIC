package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcPromedio(t *testing.T) {
	t.Run("Caso feliz sin errores", func(t *testing.T) {
		// Arrange
		a, b, c, d := 1.0, 2.0, 3.0, 4.6
		// Act
		res, _ := CalcPromedio(a, b, c, d)
		valesperado := (a + b + c + d) / 4.0

		//Assert

		assert.Equal(t, valesperado, res, "Fallo el promedio")

	})

	t.Run("Caso con errores", func(t *testing.T) {
		//Arrange
		a, b, c, d := -1.0, 2.0, 3.0, 4.6
		//Act
		_, err := CalcPromedio(a, b, c, d)
		valoresp := errors.New("No existen notas negativas")
		assert.Equal(t, valoresp, err, "FALLO el error")

	})

	//Assert

}
