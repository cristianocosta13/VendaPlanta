package ui

import (
	"errors"
	"fmt"
)

func PagarCompra(desconto float32) (float32, error) {
	var valorPago float32
	fmt.Print("Insira o valor a ser pago: ")
	fmt.Scan(&valorPago)
	if valorPago-desconto > 0 {
		return valorPago-desconto, nil
	}else{
		return valorPago, errors.New("Saldo insuficiente!")
	}
}
