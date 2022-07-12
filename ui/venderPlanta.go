package ui

import (
	"awesomeProject/dados"
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

	planta, erroPlanta := BuscarPorCod(codigo)
	_, erroCliente := BuscarClientePorCpf(cpfCliente)
	_, erroVendedor := BuscarVendedorPorCpf(cpfVendedor)

	if erroPlanta!=nil && erroCliente!=nil && erroVendedor!=nil {
		//planta.Estoque--
		valorDescontado := AplicarDesconto(planta)
		troco, erroSaldo := PagarCompra(valorDescontado)
		if erroSaldo != nil {
			fmt.Print(erroSaldo)
		}else{
			fmt.Print("O troco do cliente é R$", troco)
		}
	}else {
		fmt.Print("ERRO PLANTA: ", erroPlanta)
		fmt.Print("ERRO CLIENTE: ", erroCliente)
		fmt.Print("ERRO VENDEDOR: ", erroVendedor)
	}
}
