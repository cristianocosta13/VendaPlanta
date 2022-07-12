package ui

import (
	"awesomeProject/modelos"
	"fmt"
)

func AplicarDesconto(planta modelos.Planta) float32 {
	var desconto float32
	fmt.Print("Insira o percental de desconto: ")
	fmt.Scan(&desconto)
	return (planta.Preco * desconto / 100)
}
