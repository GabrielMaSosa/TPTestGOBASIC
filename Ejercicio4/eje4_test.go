package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	min = "minimum"
	max = "maximum"
	avg = "average"
)

func TestOperation(t *testing.T) {
	/*
		t.Run("Caso Maximum", func(t *testing.T) {
			//Arrange
			maxi := Operation(max)
			//Act
			//Assert
			assert.Equal(t, maxi, maxi, "Errors MAX")
		})

		t.Run("Caso Minimum", func(t *testing.T) {
			//Arrange
			maxi := Operation(min)
			//Act
			//Assert
			assert.Equal(t, maxi, maxi, "Errors MIN")
		})

		t.Run("Caso AVerage", func(t *testing.T) {
			//Arrange
			maxi := Operation(avg)
			//Act
			//Assert
			assert.Equal(t, maxi, maxi, "Errors MIN")
		})
	*/
}

func TestOperations(t *testing.T) {
	t.Run("Maximum Happy", func(t *testing.T) {
		//Arrange
		maxi := Operation(max)
		resulestp := 4.0
		//Act
		res, _ := maxi(2.0, 3.0, 4.0)

		//Assert

		assert.Equal(t, resulestp, res, "Errors values out of range")

	})

	t.Run("Maximum Errors", func(t *testing.T) {
		//Arrange
		maxi := Operation(max)
		erresp := errors.New("Valor que no corresponde a nota")
		//Act
		_, err := maxi(-2.0, 3.0, 4.0)

		//Assert

		assert.Equal(t, erresp, err, "Errors values out of range")

	})
	//--------------------

	t.Run("Minimum Happy", func(t *testing.T) {
		//Arrange
		maxi := Operation(min)
		resulestp := 2.0
		//Act
		res, _ := maxi(2.0, 3.0, 4.0)

		//Assert

		assert.Equal(t, resulestp, res, "Errors values out of range")

	})

	t.Run("Minimum Errors", func(t *testing.T) {
		//Arrange
		maxi := Operation(min)
		erresp := errors.New("Valor que no corresponde a nota")
		//Act
		_, err := maxi(-2.0, 3.0, 4.0)

		//Assert

		assert.Equal(t, erresp, err, "Errors values out of range")

	})
	//--------------------

	t.Run("Average Happy", func(t *testing.T) {
		//Arrange
		maxi := Operation(avg)
		resulestp := 3.0
		//Act
		res, _ := maxi(2.0, 3.0, 4.0)

		//Assert

		assert.Equal(t, resulestp, res, "Errors values out of range")

	})

	t.Run("Average Errors", func(t *testing.T) {
		//Arrange
		maxi := Operation(min)
		erresp := errors.New("Valor que no corresponde a nota")
		//Act
		_, err := maxi(-2.0, 3.0, 4.0)

		//Assert

		assert.Equal(t, erresp, err, "Errors values out of range")

	})

}
