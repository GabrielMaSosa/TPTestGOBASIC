package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGato(t *testing.T) {

	//Arrange
	val := 30
	valesp := 30 * 5.0

	//act
	res, _ := CantGato(val)
	//Assert

	assert.Equal(t, valesp, res, "Fallo test Gato")

}

func TestPerro(t *testing.T) {

	//Arrange
	val := 30
	valesp := 30 * 10.0

	//act
	res, _ := CantPerro(val)
	//Assert

	assert.Equal(t, valesp, res, "Fallo test Perro")

}
func TestTarantula(t *testing.T) {

	//Arrange
	val := 30
	valesp := 30 * 0.15

	//act
	res, _ := CantTarantula(val)
	//Assert

	assert.Equal(t, valesp, res, "Fallo test Tarantula")

}
func TestHamster(t *testing.T) {

	//Arrange
	val := 30
	valesp := 30 * 0.25

	//act
	res, _ := CantHamster(val)
	//Assert

	assert.Equal(t, valesp, res, "Fallo test Hamster")

}
