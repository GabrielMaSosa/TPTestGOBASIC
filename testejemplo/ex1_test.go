package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {
	//inicializamos variables
	num1 := 5
	num2 := 5
	resultadoesperado := 10
	//	ejecutamos el test
	resultado := Suma(num1, num2)

	assert.Equal(t, resultadoesperado, resultado, "La suma no es correcta")

}
