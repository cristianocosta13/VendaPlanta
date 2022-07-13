package ui

import (
	"awesomeProject/data"
	"fmt"
)

func VenderPlanta() {
	var codigo int
	var cpfVendedor, cpfCliente string

	fmt.Print("Insira o nome do vendedor responsável pela venda: ")
	fmt.Scan(&cpfVendedor)
	fmt.Print("Insira o CPF do cliente: ")
	fmt.Scan(&cpfCliente)
	fmt.Print("Insira o código da planta a ser comprada: ")
	fmt.Scan(&codigo)

	planta, erroPlanta := data.BuscarPorCod(codigo)
	_, erroCliente := data.BuscarClientePorCpf(cpfCliente)
	_, erroVendedor := data.BuscarVendedorPorCpf(cpfVendedor)

	if erroPlanta == nil && erroCliente == nil && erroVendedor == nil {
		planta.Estoque-- //não está decrementando, problema no ponteiro
		valorDescontado := AplicarDesconto(*planta)
		troco, erroSaldo, valorPendente := PagarCompra(valorDescontado, *planta)
		if erroSaldo != nil {
			fmt.Print(erroSaldo, "\nO cliente deve pagar mais R$", valorPendente)
		} else {
			fmt.Print("O troco do cliente é R$", troco)
		}
	} else {
		fmt.Print("ERRO PLANTA: ", erroPlanta)
		fmt.Println("ERRO CLIENTE: ", erroCliente)
		fmt.Println("ERRO VENDEDOR: ", erroVendedor)
	}
}
