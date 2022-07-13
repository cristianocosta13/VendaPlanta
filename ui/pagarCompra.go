package ui

import (
	"awesomeProject/modelos"
	"errors"
	"fmt"
)

func PagarCompra(desconto float32, planta modelos.Planta) (float32, error, float32) {
	var valorPago float32
	fmt.Print("Insira o valor a ser pago: ")
	fmt.Scan(&valorPago)
	if valorPago > planta.Preco {
		return valorPago - (planta.Preco - desconto), nil, 0
	} else {
		return valorPago, errors.New("Saldo insuficiente!"), planta.Preco - valorPago
	}
}
