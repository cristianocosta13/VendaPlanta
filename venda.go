package main

import (
	"awesomeProject/modelos"
	"fmt"
)

type VendaPlanta struct {
	planta   modelos.Planta
	vendedor Vendedor
	cliente  Cliente
}

func venderPlanta(planta modelos.Planta, vendedor Vendedor, cliente Cliente) {
	var codigo int
	var desconto, valorDescontado, troco float32
	var nomeVendedor, cpfCliente string

	fmt.Print("Insira o nome do vendedor responsável pela venda: ")
	fmt.Scan(&nomeVendedor)
	fmt.Print("Insira o CPF do cliente: ")
	fmt.Scan(&cpfCliente)
	fmt.Print("Insira o código da planta a ser comprada: ")
	fmt.Scan(&codigo)

	if cliente.cpf == cpfCliente && vendedor.nome == nomeVendedor && planta.Codigo == codigo {
		planta.Estoque--
		fmt.Print("Qual o percentual de desconto?")
		fmt.Scan(&desconto)
		valorDescontado = aplicarDesconto(desconto, planta)
		fmt.Print("O desconto aplicado foi de R%", valorDescontado)
		troco = cliente.pagarCompra(valorDescontado)
		if troco < 0 {
			fmt.Print("Saldo insuficiente!")
		} else if troco > 0 {
			fmt.Print("Compra realizada com sucesso! Seu troco é de R$", troco)
		} else {
			fmt.Print("Compra realizada com sucesso e sem troco!")
		}
	} else {
		fmt.Print("As informações fornecidas não se relacionam com nenhum vendedor, cliente ou planta da loja\n")
	}
}

func aplicarDesconto(desconto float32, planta modelos.Planta) float32 {
	return planta.Preco - (planta.Preco * desconto / 100)
}
