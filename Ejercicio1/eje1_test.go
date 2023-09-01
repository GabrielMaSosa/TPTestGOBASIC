package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcImpuesto(t *testing.T) {
	//test1
	t.Run("TEST caso 1", func(t *testing.T) {

		//Arrange

		val1 := 49000.0
		valesperado1 := val1 * 0.0
		//Action
		resul1 := CalcImpuesto(val1)
		//Assert
		assert.Equal(t, valesperado1, resul1, "Fallo el test")

	})

	t.Run("TEST caso2", func(t *testing.T) {

		//Arrange

		val2 := 50001.0
		valesperado2 := val2 * 0.17
		// act
		resul2 := CalcImpuesto(val2)
		// Assert
		assert.Equal(t, valesperado2, resul2, "Fallo el test")

	})
	t.Run("TEST caso3", func(t *testing.T) {
		//Arrange
		val3 := 150001.0
		valesperado3 := val3 * 0.27
		//Action
		resul3 := CalcImpuesto(val3)
		//Assert
		assert.Equal(t, valesperado3, resul3, "Fallo el test")
	})

}
