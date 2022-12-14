package main

import "testing"

//SOMA
func TestShouldSumCorrect(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestShouldSumIncorrect(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

//SUBTRAÇÃO
func TestShouldSubCorrect(t *testing.T) {
	teste := subtracao(5, 10)
	resultado := 5

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestShouldSubIncorrect(t *testing.T) {
	teste := subtracao(5, 10)
	resultado := 10

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

//MULTIPLICAÇÃO
func TestShouldMultCorrect(t *testing.T) {
	teste := multiplicacao(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestShouldMultIncorrect(t *testing.T) {
	teste := multiplicacao(10, 10)
	resultado := 1000

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

//DIVISÃO
func TestShouldDivtCorrect(t *testing.T) {
	teste := divisao(20)
	resultado := 20

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func TestShouldDivIncorrect(t *testing.T) {
	teste := divisao(20)
	resultado := 10

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
